package singleton_pattern

import (
	"sync"
)

type Singleton interface {
	GetInstance() interface{}
}

func NewSingleton(new func() interface{}) *singleton {
	return &singleton{
		new,
	}
}

var once = &sync.Once{}
var instance interface{}

type singleton struct {
	new func() interface{}
}

func (s *singleton) GetInstance() interface{} {
	once.Do(func() {
		instance = s.new()
	})

	return instance
}
