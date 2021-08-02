package test

import (
	"fmt"
	"testing"

	"github.com/zengzhengrong/zzgo/zqueue"
)

func TestQueue(t *testing.T) {
	q := zqueue.New()
	for i := 0; i < 100; i++ {
		q.Push(i)
	}
	fmt.Println(q.Len())
	top := q.Peek()
	fmt.Println(top)
	fmt.Println(q.Len())
}
