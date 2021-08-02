package zgo

/*
在资源有限的情况下的并发处理,根据channel 缓冲来决定并发数
*/

//GoLimit is control rate signal
type GoLimit struct {
	c chan struct{}
}

func (l GoLimit) Run(i interface{}, f func(n interface{})) {

	l.c <- struct{}{}
	go func(n interface{}) {
		f(n)
		<-l.c
		// 在 goroutine 里面释放掉channel的值,以便可以继续执行
	}(i)
	// 以缓存10个为例，第一批会并发执行10个goroutine 如果有一个goroutine 执行完毕则会继续执行下一个goroutine,保持只有10个goroutine在运行

}

func NewGoLimit(buffer int) *GoLimit {
	return &GoLimit{make(chan struct{}, buffer)}
}
