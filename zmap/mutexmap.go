package zmap

import (
	"sync"
	"time"
)

type MutexMap interface {
	Set(key string, val interface{})                   //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Get(key string, timeout time.Duration) interface{} //读取一个key，如果key不存在阻塞，等待key存在或者超时
}

//Map is 并发安全的map
type Map struct {
	C   map[string]*Entry
	RMX *sync.RWMutex
}
type Entry struct {
	ch      chan struct{}
	value   interface{}
	isExist bool
}

// Set 设置一个key
func (m *Map) Set(key string, val interface{}) {
	m.RMX.Lock() // 上锁
	defer m.RMX.Unlock()
	if e, ok := m.C[key]; ok {
		e.value = val // 如果键的值存在 则覆盖原有的
		e.isExist = true
		close(e.ch) // 关闭 chnnal
	} else {
		// 不存在则新建
		e = &Entry{ch: make(chan struct{}), isExist: true, value: val}
		m.C[key] = e
		close(e.ch)
		// 注意close 一个channel也可以使select语句不再阻塞
	}
}

// Get 读取一个key 不存在则阻塞 , key
func (m *Map) Get(key string, timeout time.Duration) interface{} {
	m.RMX.Lock()
	if e, ok := m.C[key]; ok && e.isExist {
		m.RMX.Unlock()
		return e.value
	} else if !ok {
		// 没有该键 则生成一个
		e = &Entry{ch: make(chan struct{}), isExist: false}
		m.C[key] = e

		m.RMX.Unlock()
		select {
		case <-e.ch:
			return e.value
		case <-time.After(timeout):

			return nil
		}
	} else {
		// 键存在 但是值不存在(nil) isExist is false

		m.RMX.Unlock()
		select {
		case <-e.ch:
			return e.value
		case <-time.After(timeout):

			return nil
		}
	}
}
