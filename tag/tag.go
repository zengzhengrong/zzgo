package tag

import (
	"context"
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/zengzhengrong/zzgo/zgo"
	"github.com/zengzhengrong/zzgo/ztime"
	"github.com/zengzhengrong/zzgo/zwg"
)

const (
	// 与运算  0011 0000 and 0111 0000 = 0011 0000
	and int = 1
	// 或运算 0011 0000 and 0101 0010 = 0111 0010
	or int = 2
	// 异或运算 0011 0000 and 0101 0010 = 0110 1101
	xor  int  = 3
	zero byte = byte('0')
	one  byte = byte('1')
)

var (
	RDB                   *redis.Client
	GOLIMIT_BUFFER        int = 100
	DEFAULT_OPERATION_TTL int = 600
	once                  sync.Once
)

// Result Map is offset of bitmap result
type Result map[int64]int64

func Init(client *redis.Client) error {
	ctx := context.Background()
	RDB = client
	status := RDB.Ping(ctx)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

// Tag is 标签管理 使用redis bitmap 方法来存储和管理对应id
type Tag struct {
	Name string     `json:"name"`
	Time *time.Time `json:"time"`
	ctx  context.Context
}

// GetFullName is 获取实际的redis key
func (t *Tag) GetFullName() string {

	if t.Time != nil {
		unix := ztime.GetTimeUnix(*t.Time)
		return fmt.Sprintf("%s:%d", t.Name, unix)
	}
	return t.Name
}

func (t *Tag) set(bit int, offsets ...int64) Result {
	result := make(Result, len(offsets))
	key := t.GetFullName()
	for _, offset := range offsets {
		switch bit {
		case 1:
			r := RDB.SetBit(t.ctx, key, offset, 1)
			result[offset] = r.Val()
		case 0:
			r := RDB.SetBit(t.ctx, key, offset, 0)
			result[offset] = r.Val()
		}

	}
	return result
}

// SetOnce is offset (id) 设置bit为1,返回为0，表明已对该位进行bit 1操作
func (t *Tag) SetOnce(offset int64) int64 {
	r := t.set(1, offset)[offset]
	return r
}

// UnSetOnce is offset (id) 设置bit为0,返回为1，表明已对该位进行bit 0操作
func (t *Tag) UnSetOnce(offset int64) int64 {
	r := t.set(0, offset)[offset]
	return r
}

// Set is offset (id) 设置bit为1,返回Result键值类型,如果offset值为0，表明已对该位进行bit 1操作
func (t *Tag) Set(offsets ...int64) Result {
	return t.set(1, offsets...)
}

// UnSet is offset (id) 设置bit为0,返回Result键值类型,如果offset值为1，表明已对该位进行bit 0操作
func (t *Tag) UnSet(offsets ...int64) Result {
	return t.set(0, offsets...)
}

// Get is offset (id) 获取该offset的bit, 返回Result键值类型,如果没有该offset 则默认是0
func (t *Tag) Get(offsets ...int64) Result {
	result := make(Result, len(offsets))
	key := t.GetFullName()
	for _, offset := range offsets {
		r := RDB.GetBit(t.ctx, key, offset)
		result[offset] = r.Val()
	}

	return result
}

// GetOnce is offset (id) 获取该offset的bit, 如果没有该offset 则默认是0
func (t *Tag) GetOnce(offset int64) int64 {
	key := t.GetFullName()
	r := RDB.GetBit(t.ctx, key, offset)
	return r.Val()
}

// Count is 计算当前有多少bit 为1的offset，opt 起始位和结束位
// 起始位和结束位是按照8bit的长度来计算
// 例如 start 0, end 1 表示第一位第二个位共16个bit, 同时左右都是闭区间
func (t *Tag) Count(opt *redis.BitCount) int64 {
	key := t.GetFullName()
	result := RDB.BitCount(t.ctx, key, opt)
	return result.Val()
}

// GetRange 返回开始字节位到结束字节位的二进制字符串
func (t *Tag) GetRange(start, end int64) string {
	key := t.GetFullName()
	result := RDB.GetRange(t.ctx, key, start, end)
	b, err := result.Bytes()
	if err != nil {
		return ""
	}
	binbytes := BytesToBinary(b, false)
	return string(binbytes)
}

// SetRange is  在offset 字节位设置二进制字符串 返回该key的字节长度
func (t *Tag) SetRange(offset int64, binary string) int64 {
	key := t.GetFullName()
	r := Str2DEC(binary)
	s := string(rune(r))
	result := RDB.SetRange(t.ctx, key, offset, s)
	return result.Val()
}

// Str2DEC is 二进制字符串转十进制 int
func Str2DEC(s string) (num int) {
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		num += (int(s[l-i-1]) - 48) << uint8(i)
	}
	return
}

// Pos is 查找字符串中第一个设置为1或0的bit位, pos 可填入开始的位数和结束位数
func (t *Tag) Pos(bit int64, pos ...int64) int64 {
	key := t.GetFullName()
	result := RDB.BitPos(t.ctx, key, bit, pos...)
	return result.Val()
}

// Not is new a Tag instance of not bitmap key
func (t *Tag) Not() *Tag {
	key := t.GetFullName()
	destKey := fmt.Sprintf("%s_%s", "not", key)
	RDB.BitOpNot(t.ctx, destKey, key)
	return &Tag{Name: destKey, Time: nil, ctx: context.Background()}
}

// Len is count key length 获取key 的长度 返回数字是对应的字节数数
func (t *Tag) Len() int64 {
	key := t.GetFullName()
	reuslt := RDB.StrLen(t.ctx, key)
	return reuslt.Val()
}

func BytesToBinary(bs []byte, addSpace bool) []byte {
	l := len(bs)
	bl := l*8 + l + 1
	buf := make([]byte, 0, bl)
	for _, b := range bs {
		buf = appendBinaryString(buf, b)
		if addSpace {
			buf = append(buf, byte(' '))
		}
	}
	return buf
}

// append bytes of string in binary format.
func appendBinaryString(bs []byte, b byte) []byte {
	var a byte
	for i := 0; i < 8; i++ {
		a = b
		b <<= 1
		b >>= 1
		switch a {
		case b:
			bs = append(bs, zero)
		default:
			bs = append(bs, one)
		}
		b <<= 1
	}
	return bs
}

// FastAll is 直接获取整个字符串，在程序里面做offset和值的处理,只需要请求一次redis
func (t *Tag) FastAll() ([]int64, error) {
	key := t.GetFullName()
	size := t.Count(nil)
	result := make([]int64, size)
	r := RDB.Get(t.ctx, key)
	if r.Err() != nil {
		return []int64{}, r.Err()
	}
	b, err := r.Bytes()
	if err != nil {
		return []int64{}, err
	}
	binbytes := BytesToBinary(b, false)
	wg := &sync.WaitGroup{}
	num := 64
	if size < 64 {
		num = int(size)
	}
	bytesSplit(num, binbytes, func(l []byte, lower, hight int) {
		wg.Add(1)
		go func(l []byte, lower, hight int) {
			defer wg.Done()
			for ; lower < hight; lower++ {
				switch l[lower] {
				case one:
					size := atomic.AddInt64(&size, int64(-1))
					result[size] = int64(lower)
				default:
					continue
				}
			}
		}(l, lower, hight)
	})
	wg.Wait()
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	return result, nil
}

//bytesSplit is 切片分割, 按num 个数进行分割
func bytesSplit(num int, list []byte, exec func(l []byte, low, hight int)) {

	for i := 1; i <= int(math.Floor(float64(len(list)/num)))+1; i++ {
		// 将list 除于 num 获取 分割次数 非整数时 舍去小数点 后+1
		low := num * (i - 1) // 左索引

		if low > len(list) {
			// 左索引 大于长度列表长度则直接返回
			return
		}
		high := num * i // 右索引
		if high > len(list) {
			// 如果右索引大于 list长度，则取list的长度
			high = len(list)
		}
		exec(list, low, high)
	}
}

// All 获取所有bit为1 的offsets
func (t *Tag) All(timeout ...time.Duration) []int64 {
	var i int64
	size := t.Count(nil)

	result := make([]int64, size)
	leng := t.Len()
	gl := zgo.NewGoLimit(GOLIMIT_BUFFER)
	gr := zgo.NewGoLimit(GOLIMIT_BUFFER)
	// 计算含有1的字节位的offset,返回-1表示没有包含1的字节位
	wg := &sync.WaitGroup{}
	flag := make(chan struct{})
	fail := make(chan bool)
	if len(timeout) > 0 {
		go func(timeout time.Duration) {
			<-flag
			// time out
			if zwg.WaitTimeout(wg, timeout) {
				fail <- true
			} else {
				fail <- false
			}

		}(timeout[0])
		go func(gl, gr *zgo.GoLimit) {
			if <-fail {
				gr.Cancel()
				gl.Cancel()
				fail <- true
			} else {
				fail <- false
			}

		}(gl, gr)
	}

	for i = 0; i < leng; i++ {

		wg.Add(1)
		if len(timeout) > 0 && i == 0 {
			flag <- struct{}{}
		}

		gl.Run(i, func(i interface{}) {
			defer wg.Done()
			ii := i.(int64)
			if t.Pos(1, ii, ii) != -1 {
				wg.Add(1)
				gr.Run(ii, func(n interface{}) {
					defer wg.Done()
					nn := n.(int64)
					offsetStart := (nn * 8)
					offsetEnd := (nn + 1) * 8
					for ; offsetStart < offsetEnd; offsetStart++ {
						if t.GetOnce(offsetStart) == 1 {
							size := atomic.AddInt64(&size, int64(-1))
							result[size] = offsetStart
						}

					}
				})

			}
		})
	}
	if len(timeout) == 0 {
		wg.Wait()
	} else {
		if <-fail {
			return nil
		} else {
			wg.Wait()
		}
	}

	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	return result
}

func operationKeyName(t string, tags ...*Tag) ([]string, string) {
	sep := fmt.Sprintf("_%s_", t)

	names := make([]string, len(tags))
	for index, tag := range tags {
		names[index] = tag.GetFullName()

	}

	return names, strings.Join(names, sep)
}

func operation(ttl bool, t int, tags ...*Tag) *Tag {
	var names []string
	var destKey string
	ctx := context.Background()
	switch t {
	case and:
		names, destKey = operationKeyName("AND", tags...)
		RDB.BitOpAnd(ctx, destKey, names...)
	case or:
		names, destKey = operationKeyName("OR", tags...)
		RDB.BitOpOr(ctx, destKey, names...)
	case xor:
		names, destKey = operationKeyName("XOR", tags...)
		RDB.BitOpXor(ctx, destKey, names...)
	}
	return &Tag{Name: destKey, Time: nil, ctx: ctx}
}

// AND is and operation of bitmap
func AND(ttl bool, tags ...*Tag) *Tag {
	return operation(ttl, and, tags...)
}

// OR is and operation of bitmap
func OR(ttl bool, tags ...*Tag) *Tag {
	return operation(ttl, or, tags...)
}

// XOR is and operation of bitmap
func XOR(ttl bool, tags ...*Tag) *Tag {
	return operation(ttl, xor, tags...)

}

// NewTag is generate Tag instance if time is nil , return no time keyname of redis key
func NewTag(name string, time *time.Time) (*Tag, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if RDB == nil {
		return nil, fmt.Errorf("must init redis client before")
	}
	return &Tag{Name: name, Time: time, ctx: context.Background()}, nil

}
