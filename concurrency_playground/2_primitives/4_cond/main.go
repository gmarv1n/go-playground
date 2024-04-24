package main

import (
	"fmt"
	"sync"
	"time"
)

func subscribe(name string, data map[string]string, c *sync.Cond) {
	c.L.Lock()

	// Wait нужно дергать обязательно а цикле на условии какого то события
	// это нужно по такой причине: нам приходит Broadcast, но перед тем как взять
	// лок другая горутина уже успела взять лок
	// А так получается мы получим бродкаст, поставим лок, проверим в следующей итерации
	// что нужное нам состояние данных соблюдено и если кто-то уже успел
	// например удалить данные (для этого примера), то мы опять уйдем в wait и
	// снимем лок
	for len(data) == 0 {
		c.Wait() // Wait снимает блокировки и ждет Broadcast()
		// когда приходить Broadcast - снова ставится Lock
		// Wait() переводит горутину в Waiting
	}

	fmt.Printf("[%s] %s\n", name, data["key"])

	c.L.Unlock()
}

func publish(name string, data map[string]string, c *sync.Cond) {
	time.Sleep(time.Second)

	c.L.Lock()
	data["key"] = "value"
	c.L.Unlock()

	fmt.Printf("[%s] data publisher\n", name)

	// Cond можно переиспользовать и броадкастить много раз
	c.Broadcast()
}

func main() {
	data := map[string]string{}
	cond := sync.NewCond(&sync.Mutex{})

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		subscribe("subscriber_1", data, cond)
	}()

	go func() {
		defer wg.Done()
		subscribe("subscriber_2", data, cond)
	}()

	go func() {
		defer wg.Done()
		publish("publisher", data, cond)
	}()

	wg.Wait()
}

// Lesson #_ video, time: __:__:__
