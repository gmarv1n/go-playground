package main

import (
	"sync"
)

type Stack struct {
	mutex sync.Mutex
	data  []string
}

func NewStack() *Stack {
	return &Stack{}
}

func (b *Stack) Push(value string) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.data = append(b.data, value)
}

func (b *Stack) Pop() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if len(b.data) == 0 {
		panic("pop: stack is empty")
	}

	//b.mutex.Lock()
	//defer b.mutex.Unlock() // NOT HERE cause len() causes data race

	b.data = b.data[:len(b.data)-1]
}

func (b *Stack) Top() string {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if len(b.data) == 0 {
		panic("top: stack is empty")
	}

	//b.mutex.Lock()
	//defer b.mutex.Unlock() // NOT HERE cause len() causes data race

	return b.data[len(b.data)-1]
}

var stack Stack

func producer() {
	for i := 0; i < 1000; i++ {
		stack.Push("message")
	}
}

func (b *Stack) TopAndPop() string {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if len(b.data) == 0 {
		panic("top: stack is empty")
	}

	//b.mutex.Lock()
	//defer b.mutex.Unlock() // NOT HERE cause len() causes data race

	val := b.data[len(b.data)-1]
	b.data = b.data[:len(b.data)-1]
	return val
}

func consumer() {
	for i := 0; i < 10; i++ {
		// _ = stack.Top() // Bad here: there can be race condition between Top and Pop
		// interruption
		// stack.Pop()

		_ = stack.TopAndPop()
	}
}

func main() {
	producer()

	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			consumer()
		}()
	}

	wg.Wait()
}
