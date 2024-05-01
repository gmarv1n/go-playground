package main

import (
	"runtime"
	"sync/atomic"
)

// SpinLock - По сути это Mutex
type SpinLock struct {
	state atomic.Bool
}

func NewSpinLock() SpinLock {
	return SpinLock{}
}

func (s *SpinLock) Lock() {
	for !s.state.CompareAndSwap(false, true) {
		// active waiting here until someone release the state to false
		// из минусов - тут запаркуется горутина и перейдет в waiting
		runtime.Gosched() // Gosched() не переведет горутину в waiting, если его заюзать
	}
	// finally lock
}

func (s *SpinLock) TryLock() bool {
	return s.state.CompareAndSwap(false, true)
}

func (s *SpinLock) Unlock() {
	if !s.state.CompareAndSwap(true, false) {
		panic("incorrect usage")
	}
}

// Lesson #_ video, time: __:__:__
