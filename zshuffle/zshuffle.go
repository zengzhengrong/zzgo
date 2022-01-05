package zshuffle

import (
	"math/rand"
	"time"
)

// Shuffle is 打乱切片元素
func Shuffle[T any](e []T) {
	if len(e) > 1 {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(e), func(x, y int) {
			e[x], e[y] = e[y], e[x]
		})
	}
}
