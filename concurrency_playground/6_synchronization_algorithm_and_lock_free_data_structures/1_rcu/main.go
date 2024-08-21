package main

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

type Cache struct {
	mutex sync.RWMutex
	data  map[string]string
}

func NewCache(ctx context.Context) *Cache {
	cache := &Cache{
		data: make(map[string]string),
	}

	go cache.synchronize(ctx)

	return cache
}

func (c *Cache) synchronize(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	// раз в минуту инвалидируем кеш
	for {
		// приоритизируем чек контекст дана
		select {
		case <-ctx.Done():
			return
		default:

		}

		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			var temp map[string]string
			// data = ... - get from remote storage

			// 5. грузим темп в ансейф поинтер
			ptr := unsafe.Pointer(&temp)

			// 5. грузим дату в ансейф поинтер
			dataPtr := unsafe.Pointer(&c.data)

			// Эта шляпа атомарно грузит значение ptr
			// в ссылку dataPtr
			atomic.StorePointer(&dataPtr, ptr)

			// 6. прикол в том, что мы избавились от
			// лочки

			//c.mutex.Lock()
			//c.data = temp // эффективная лочка, тут под ней тока значение присваивается новое и все
			//c.mutex.Unlock()
		}
	}
}

func (c *Cache) Get(key string) (string, bool) {
	// геттер для кеша

	// лочим на чтение
	//c.mutex.RLock()
	//defer c.mutex.RUnlock() // по деферу отпускаем

	// 1. мапа - это поинтер
	// приводим мапу к поинтеру
	ptr := unsafe.Pointer(&c.data)

	// 2. локально грузим мапу в переменную
	data := atomic.LoadPointer(&ptr)

	// 3. кастим ансейф поинтер в мапу по указателю
	m := (*map[string]string)(data)

	// 4. разыменовываем
	// т.е. достаем в переменную значение которое лежит
	// по указателю
	mm := *m

	// достаем значение
	value, found := mm[key]
	return value, found

	// RCU значит Read Copy Update
	// каждый читает значение, копирует его себе и потом если
	// надо то обновляет

	// достаем значение
	//value, found := c.data[key]
	//return value, found
}

// Lesson #8 video, time: 00:07:50
