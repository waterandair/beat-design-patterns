package singleton_pattern

import (
	"testing"
)

// 单例模式基准测试
// goos: linux
// goarch: amd64
// pkg: github.com/waterandair/beat-design-patterns/pkg/singleton_pattern
// BenchmarkSing-16                        300000000                5.24 ns/op
// BenchmarkSingletonMutex-16              30000000                43.4 ns/op
// BenchmarkSingletonDCMutex-16            500000000                3.47 ns/op
// BenchmarkSingletonThreadUnsafe-16       2000000000               0.50 ns/op
// PASS

// 从测试结果中可以发现，对于安全的单例模式实现，使用 sync.Once 的执行效率最高

func benchmarkGetInstance() interface{} {
	return new(struct{})
}
func BenchmarkSing(b *testing.B) {
	b.ResetTimer()
	singleton := NewSingleton(benchmarkGetInstance)
	for i := 0; i < b.N; i++ {
		singleton.GetInstance()
	}
}

func BenchmarkSingletonMutex(b *testing.B) {
	b.ResetTimer()
	singleton := NewSingletonMutex(benchmarkGetInstance)
	for i := 0; i < b.N; i++ {
		singleton.GetInstance()
	}
}

func BenchmarkSingletonDCMutex(b *testing.B) {
	b.ResetTimer()
	singleton := NewSingletonDCMutex(benchmarkGetInstance)
	for i := 0; i < b.N; i++ {
		singleton.GetInstance()
	}
}

func BenchmarkSingletonThreadUnsafe(b *testing.B) {
	b.ResetTimer()
	singleton := NewSingletonThreadUnsafe(benchmarkGetInstance)
	for i := 0; i < b.N; i++ {
		singleton.GetInstance()
	}
}
