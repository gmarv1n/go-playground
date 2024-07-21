package main

import (
	"fmt"
	"time"
)

type Future struct {
	result <-chan interface{}
}

func NewFuture(result <-chan interface{}) Future {
	return Future{
		result: result,
	}
}

func (f Future) Get() interface{} {
	return <-f.result
}

type Promise struct {
	result   chan interface{}
	promised bool
}

func NewPromise() Promise {
	return Promise{
		result: make(chan interface{}, 1),
	}
}

func (p *Promise) Set(value interface{}) {
	if p.promised {
		return
	}

	p.promised = true
	p.result <- value
	close(p.result)
}

func (p *Promise) GetFuture() Future {
	return NewFuture(p.result)
}

func main() {
	promise := NewPromise()
	go func() {
		time.Sleep(time.Second)
		promise.Set("Test")
	}()

	future := promise.GetFuture()
	result := future.Get()
	fmt.Println(result)
}

// Lesson #6 video, time: 00:00:00
