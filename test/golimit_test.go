package test

import (
	"fmt"
	"testing"

	"github.com/zengzhengrong/zzgo/zgo"
)

func TestRateLimt(t *testing.T) {
	// golimit
	golimit := zgo.NewGoLimit(10)
	for i := 0; i < 100; i++ {
		golimit.Run(i, func(n interface{}) { fmt.Println(n) })
	}
}
