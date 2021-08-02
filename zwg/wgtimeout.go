package zwg

import (
	"sync"
	"time"
)

// WaitTimeout is timeout for witgroup
func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {

	// 如果timeout到了超时时间返回true
	// 如果WaitGroup自然结束返回false
	to := make(chan bool)

	go func(c chan bool, d time.Duration) {
		time.AfterFunc(d, func() {
			to <- true
		})
	}(to, timeout)

	go func(c chan bool) {
		wg.Wait()
		to <- false
	}(to)
	return <-to

}
