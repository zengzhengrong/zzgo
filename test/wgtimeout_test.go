package test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/zengzhengrong/zzgo/zwg"
)

func TestWgTimeOut(t *testing.T) {
	wg := sync.WaitGroup{}
	c := make(chan struct{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int, close <-chan struct{}) {
			defer wg.Done()
			pop := <-close
			fmt.Println(pop)
			fmt.Println(num)
		}(i, c)
	}

	go func() {
		for i := 0; i < 10; i++ {
			c <- struct{}{}
		}
	}()

	if zwg.WaitTimeout(&wg, time.Millisecond*1) {
		close(c)
		fmt.Println("timeout exit")
	}
	time.Sleep(time.Second * 1)
}
