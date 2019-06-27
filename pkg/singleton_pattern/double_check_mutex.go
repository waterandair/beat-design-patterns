package singleton_pattern

import "sync"

func NewSingletonDCMutex(new func() interface{}) *singletonDCMutex {
	return &singletonDCMutex{
		new,
	}
}

var instanceDCMutex interface{}
var muDC sync.Mutex

type singletonDCMutex struct {
	new func() interface{}
}

func (s *singletonDCMutex) GetInstance() interface{} {
	if instanceDCMutex == nil {
		// 没有被实例化才会加锁
		muDC.Lock()
		defer muDC.Unlock()

		if instanceDCMutex == nil {
			instanceDCMutex = s.new()
		}
	}

	return instanceDCMutex
}
