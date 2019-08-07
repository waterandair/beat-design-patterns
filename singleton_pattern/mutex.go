package singleton_pattern

import "sync"

func NewSingletonMutex(new func() interface{}) *singletonMutex {
	return &singletonMutex{
		new,
	}
}

var instanceMutex interface{}
var mu sync.Mutex

type singletonMutex struct {
	new func() interface{}
}

func (s *singletonMutex) GetInstance() interface{} {
	// 无脑加锁
	mu.Lock()
	defer mu.Unlock()

	if instanceMutex == nil {
		instanceMutex = s.new()
	}

	return instanceMutex
}
