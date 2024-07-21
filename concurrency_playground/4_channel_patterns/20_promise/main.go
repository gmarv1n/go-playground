package main

import (
	"errors"
	"fmt"
	"time"
)

type Promise struct {
	closeCh     chan struct{}
	closeDoneCh chan struct{}
	value       interface{}
	err         error
}

func (p *Promise) Then(successCb func(interface{}), errCb func(error)) {
	go func() {
		<-p.closeCh
		defer close(p.closeDoneCh)

		if p.err == nil {
			successCb(p.value)
		} else {
			errCb(p.err)
		}
	}()

	<-p.closeDoneCh
}

func NewPromise(task func() (interface{}, error)) *Promise {
	promise := &Promise{
		closeCh:     make(chan struct{}),
		closeDoneCh: make(chan struct{}),
	}

	go func() {
		defer close(promise.closeCh)
		promise.value, promise.err = task()
	}()

	return promise
}

func main() {
	callback := func() (interface{}, error) {
		// some long operation
		time.Sleep(time.Second)
		return "success", nil
	}

	callbackError := func() (interface{}, error) {
		// some long operation
		time.Sleep(time.Second)
		return "not success", errors.New("some err")
	}

	promise := NewPromise(callback)
	promise.Then(
		func(value interface{}) {
			fmt.Println(value)
		},
		func(err error) {
			fmt.Println(err.Error())
		},
	)

	promise2 := NewPromise(callbackError)
	promise2.Then(
		func(value interface{}) {
			fmt.Println(value)
		},
		func(err error) {
			fmt.Println(err.Error())
		},
	)
}

// Lesson #6 video, time: 01:09:20
