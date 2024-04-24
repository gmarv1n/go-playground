package main

import "sync"

type Counters struct {
	mu sync.Mutex
	m  map[string]int
}

func (c *Counters) Load(key string) (int, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, found := c.m[key]
	return value, found
}

func (c *Counters) Store(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.m[key] = value
}

type CountersRW struct {
	mu sync.RWMutex // вот оно
	m  map[string]int
}

func (c *CountersRW) Load(key string) (int, bool) {
	c.mu.TryRLock() // по сути для читателя будет свой мутекс
	defer c.mu.RUnlock()

	value, found := c.m[key]
	return value, found
}

func (c *CountersRW) Store(key string, value int) {
	c.mu.Lock() // сработает только когда нет читателей
	defer c.mu.Unlock()

	c.m[key] = value
}

func main() {

}

// Lesson #_ video, time: __:__:__
