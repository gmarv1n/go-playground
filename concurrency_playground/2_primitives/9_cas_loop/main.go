package main

import (
	"sync/atomic"
)

func IncrementAndGet(pointer *int32) int32 {
	// Вот он CAS loop:
	for {
		currentValue := atomic.LoadInt32(pointer)
		// another goroutine may change the value here
		nextValue := currentValue + 1
		// another goroutine may change the value here
		if atomic.CompareAndSwapInt32(pointer, currentValue, nextValue) {
			// another goroutine may change the value here (???) скорее всего нет
			return nextValue
		}
	}
}

// Lesson #_ video, time: __:__:__
