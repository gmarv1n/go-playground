package main

import "fmt"

type Future struct {
	result chan interface{}
}

func NewFuture(task func() interface{}) *Future {
	future := &Future{
		result: make(chan interface{}),
	}

	go func() {
		future.result <- task() // тут мы стартуем колбек
	}()

	return future
}

func (f *Future) Get() interface{} {
	return <-f.result // вызывая гет вешаем горутину из которой
	// вызываем пока не будет выполнена таска
	// как тока будет выполнена - горутина из которой вызвали
	// разлочится и получит значение
}

func main() {
	callback := func() interface{} {
		// operation
		return "success"
	}

	future := NewFuture(callback)
	result := future.Get()

	fmt.Println(result)
}

// Lesson #6 video, time: 00:00:00
