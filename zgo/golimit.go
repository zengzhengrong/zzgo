package zgo

import (
	"context"
	"sync"
)

/*
在资源有限的情况下的并发处理,根据channel 缓冲来决定并发数
*/

//GoLimit is control rate signal
type GoLimit struct {
	C      chan struct{}
	Ctx    context.Context
	Cancel context.CancelFunc
	once   *sync.Once
}

func (l GoLimit) Run(i interface{}, f func(n interface{})) {
	defer func() {
		if recover() != nil {
			return
		}
	}()
	select {
	case <-l.Ctx.Done():
		l.once.Do(func() { close(l.C) })
		return
	default:
		l.C <- struct{}{}
		go func(n interface{}) {
			f(n)
			<-l.C
			// 在 goroutine 里面释放掉channel的值,以便可以继续执行
		}(i)
		// 以缓存10个为例，第一批会并发执行10个goroutine 如果有一个goroutine 执行完毕则会继续执行下一个goroutine,保持只有10个goroutine在运行
	}

}

func NewGoLimit(buffer int) *GoLimit {
	once := new(sync.Once)
	ctx, cancel := context.WithCancel(context.Background())
	return &GoLimit{make(chan struct{}, buffer), ctx, cancel, once}
}
