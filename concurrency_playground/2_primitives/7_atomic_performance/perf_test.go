package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

// go test -bench=. perf_test.go

func BenchmarkMutexAdd(b *testing.B) {
	var number int32
	var mutex sync.Mutex
	for i := 0; i < b.N; i++ {
		mutex.Lock()
		number++
		mutex.Unlock()
	}
}

func BenchmarkAtomicAdd(b *testing.B) {
	var number atomic.Int32
	for i := 0; i < b.N; i++ {
		number.Add(1)
	}
}

func BenchmarkAdd(b *testing.B) {
	var number int32
	for i := 0; i < b.N; i++ {
		number++
	}
}

// Lesson #4 video, time: 00:35:40
