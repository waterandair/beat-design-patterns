package singleton_pattern

import (
	"sync"
	"testing"
	"time"
)

const cycle = 10

// newInstance 返回一个创建单例的函数，并通过闭包方式记录该函数被调用的次数，因为最后需要调用该函数获取到次数，所以初始设为为 -1 而不是 0
func newInstance() func() interface{} {
	i := -1
	return func() interface{} {
		// 为了模拟多个线程同时运行创建对象的函数，这里暂停 1s, 等待所有线程执行至此
		time.Sleep(time.Second)
		i++
		return i
	}
}

// TestSingleton 测试用 sync.Once 实现的单例模式
func TestSingleton(t *testing.T) {
	fn := newInstance()

	singleton := NewSingleton(fn)

	wg := new(sync.WaitGroup)
	start := make(chan struct{})

	for i := 0; i < cycle; i++ {
		wg.Add(1)
		go func() {
			<-start

			singleton.GetInstance()
			wg.Done()
		}()
	}

	close(start)
	wg.Wait()

	num := fn()
	if num != 1 {
		t.Errorf("want 1 but get %v", num)
	}
}

// TestSingletonMutex 测试无脑用锁实现的单例模式
func TestSingletonMutex(t *testing.T) {
	fn := newInstance()

	singleton := NewSingletonMutex(fn)

	wg := new(sync.WaitGroup)
	start := make(chan struct{})

	for i := 0; i < cycle; i++ {
		wg.Add(1)
		go func() {
			<-start

			singleton.GetInstance()
			wg.Done()
		}()
	}

	close(start)
	wg.Wait()

	num := fn()
	if num != 1 {
		t.Errorf("want 1 but get %v", num)
	}
}

// TestSingletonDCMutex 测试双重判断加锁的方式实现的单例模式
func TestSingletonDCMutex(t *testing.T) {
	fn := newInstance()

	singleton := NewSingletonDCMutex(fn)

	wg := new(sync.WaitGroup)
	start := make(chan struct{})

	for i := 0; i < cycle; i++ {
		wg.Add(1)
		go func() {
			<-start

			singleton.GetInstance()
			wg.Done()
		}()
	}

	close(start)
	wg.Wait()

	num := fn()
	if num != 1 {
		t.Errorf("want 1 but get %v", num)
	}
}

// TestSingletonThreadUnsafe 测试不安全的单例模式， 结果发现创建对象的函数调用了大于 1 的次数
func TestSingletonThreadUnsafe(t *testing.T) {
	fn := newInstance()

	singleton := NewSingletonThreadUnsafe(fn)

	wg := new(sync.WaitGroup)
	start := make(chan struct{})

	for i := 0; i < cycle; i++ {
		wg.Add(1)
		go func() {
			<-start

			singleton.GetInstance()
			wg.Done()
		}()
	}

	close(start)
	wg.Wait()

	num := fn()
	if num != 1 {
		t.Errorf("want 1 but get %v", num)
	}
}
