package test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/zengzhengrong/zzgo/zmap"
)

func TestMutexMap(t *testing.T) {
	m := &zmap.Map{
		C:   make(map[string]*zmap.Entry),
		RMX: &sync.RWMutex{},
	}

	wg := &sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		if i == 0 {
			go func(m *zmap.Map) {
				defer wg.Done()
				value := m.Get("zzr", time.Duration(time.Second*10))
				fmt.Println(value)
			}(m)
		}
		if i == 1 {
			go func(m *zmap.Map) {
				defer wg.Done()
				value := m.Get("zzr", time.Duration(time.Second*10))
				fmt.Println(value)
			}(m)
		}
	}
	go func(m *zmap.Map) {
		time.Sleep(time.Second * 2)
		fmt.Println("set")
		m.Set("zzr", "zengzhengrongs")
		fmt.Println("set finsh")
	}(m)
	wg.Wait()
}
