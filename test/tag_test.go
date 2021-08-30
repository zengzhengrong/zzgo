package test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/zengzhengrong/zzgo/tag"
	"github.com/zengzhengrong/zzgo/zdup"
)

func TestConenct(t *testing.T) {
	opt, err := redis.ParseURL("redis://192.168.2.100:6379/0")
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opt)
	if err := tag.Init(client); err != nil {
		panic(err)
	}
}

func TestTagCount(t *testing.T) {
	TestConenct(t)
	bigtg, err := tag.NewTag("man", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(bigtg.Count(nil))

}

func TestLen(t *testing.T) {
	TestConenct(t)
	bigtg, err := tag.NewTag("man", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(bigtg.Len())
}

func TestGetFullName(t *testing.T) {
	TestConenct(t)
	tg, err := tag.NewTag("not_man", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(tg.GetFullName())

}

func TestAll(t *testing.T) {
	TestConenct(t)
	bigtg, err := tag.NewTag("man", nil)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 100; i++ {

	}
	rr := bigtg.All(time.Second * 20)
	fmt.Println(rr)
	fmt.Println(zdup.DuplicateInt64(rr))
}

func TestAND(t *testing.T) {
	TestConenct(t)
	tg, err := tag.NewTag("not_man", nil)
	tg2, err := tag.NewTag("woman", nil)
	if err != nil {
		panic(err)
	}
	tgAND := tag.AND(false, tg, tg2)
	fmt.Println(tgAND.All())
}

func TestOR(t *testing.T) {
	TestConenct(t)
	tg, err := tag.NewTag("not_man", nil)
	tg2, err := tag.NewTag("woman", nil)
	if err != nil {
		panic(err)
	}
	tgOR := tag.OR(false, tg, tg2)
	fmt.Println(tgOR.All())
}

func TestXOR(t *testing.T) {
	TestConenct(t)
	tg, err := tag.NewTag("not_man", nil)
	tg3, err := tag.NewTag("woman", nil)
	if err != nil {
		panic(err)
	}
	tgXOR := tag.XOR(false, tg3, tg)
	fmt.Println(tgXOR.All())
}

func TestGetRange(t *testing.T) {
	TestConenct(t)
	tg, err := tag.NewTag("not_man", nil)
	if err != nil {
		panic(err)
	}
	// 1111 1111 1110 0111 0111 0011
	// fmt.Println(tg.All())
	s := tg.GetRange(0, 4)

	fmt.Println(s)
}

func TestSetRange(t *testing.T) {
	TestConenct(t)
	tg, err := tag.NewTag("not_man", nil)
	if err != nil {
		panic(err)
	}
	// 1111 1111 1110 0111 0111 0011
	s := tg.SetRange(3, "00010000")

	fmt.Println(s)
}

func TestFastAll(t *testing.T) {
	TestConenct(t)
	tg, err := tag.NewTag("man", nil)
	if err != nil {
		panic(err)
	}
	// 1111 1111 1110 0111 0111 0011
	all := tg.All()
	fall, err := tg.FastAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(reflect.DeepEqual(all, fall))
}

func BenchmarkAll(b *testing.B) {
	TestConenct(&testing.T{})
	tg, err := tag.NewTag("man", nil)
	if err != nil {
		panic(err)
	}
	for n := 0; n < b.N; n++ {
		tg.All()
	}
}

func BenchmarkFastAll(b *testing.B) {
	TestConenct(&testing.T{})
	tg, err := tag.NewTag("man", nil)
	if err != nil {
		panic(err)
	}
	for n := 0; n < b.N; n++ {
		tg.FastAll()
	}
}

func TestBinary(t *testing.T) {
	var a byte
	var bs []byte
	fmt.Println(a)
	var b = byte('s')
	fmt.Println(b)
	fmt.Println("---------------")
	bs = appendBinaryString(bs, b)
	fmt.Println(bs)
	fmt.Println(string(bs))
}

func appendBinaryString(bs []byte, b byte) []byte {
	var zero byte = byte('0')
	var one byte = byte('1')
	var a byte
	for i := 0; i < 8; i++ {
		a = b
		fmt.Println(b)
		b <<= 1
		b >>= 1
		fmt.Println(b)
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
