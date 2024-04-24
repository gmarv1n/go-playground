package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func operation() error {
	return nil
}

func doSomething() {
	mutex.Lock()

	err := operation()
	if err != nil {
		mutex.Unlock()
		return
	}

	err = operation()
	if err != nil {
		mutex.Unlock()
		return
	}

	mutex.Unlock()
}

func doSomethingRight() {
	mutex.Lock()
	defer mutex.Unlock() // here with defer cause everywhere before return unlock

	err := operation()
	if err != nil {
		return
	}

	err = operation()
	if err != nil {
		return
	}
}

func withLock(mu *sync.Mutex, action func()) {
	if action == nil {
		return
	}

	mu.Lock()
	defer mu.Unlock()
	action()
}

func doSomethingWithUnlockOnOperation() {
	var err error

	withLock(&mutex, func() {
		err = operation()
		fmt.Println(err)
	})
	if err != nil {
		return
	}

	// конченая хуйня
	// 1. вызовется withLock
	// 2. мутекс заклочится
	// 3. выполнется func() { err = operation(); fmt.Println(err) })
	// 4. разлочится мутекс по деферу

	withLock(&mutex, func() {
		err = operation()
		fmt.Println(err)
	})
	if err != nil {
		return
	}
}

func main() {
	doSomethingWithUnlockOnOperation()
}

// Lesson #3 video, time: 00:57:25
