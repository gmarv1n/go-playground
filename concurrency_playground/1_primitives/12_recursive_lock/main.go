package main

import "sync"

type Cache struct {
	mutex sync.Mutex
	data  map[string]string
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Set(key, value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[key] = value
}

func (c *Cache) Get(key string) string {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	//if c.Size() > 0 { // вот тут будет panic изза двойного лока
	//	return c.data[key]
	//}
	// NOTE: такая хуйня ловится тестами

	if c.sizeUnsafe() > 0 {
		return c.data[key]
	}

	return ""
}

func (c *Cache) sizeUnsafe() int {
	return len(c.data)
}

func (c *Cache) Size() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.sizeUnsafe()
}

// Lesson #3 video, time: 01:13:22
