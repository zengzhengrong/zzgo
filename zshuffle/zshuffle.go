package zshuffle

import (
	"math/rand"
	"time"
)

// 打乱随机字符串
func ShuffleString(s *string) {
	if len(*s) > 1 {
		b := []byte(*s)
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(b), func(x, y int) {
			b[x], b[y] = b[y], b[x]
		})
		*s = string(b)
	}
}

// 打乱随机slice
func ShuffleSliceBytes(b []byte) {
	if len(b) > 1 {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(b), func(x, y int) {
			b[x], b[y] = b[y], b[x]
		})
	}
}

// 打乱slice int
func ShuffleSliceInt(i []int) {
	if len(i) > 1 {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(i), func(x, y int) {
			i[x], i[y] = i[y], i[x]
		})
	}
}

// 打乱slice interface
func ShuffleSliceInterface(i []interface{}) {
	if len(i) > 1 {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(i), func(x, y int) {
			i[x], i[y] = i[y], i[x]
		})
	}
}
