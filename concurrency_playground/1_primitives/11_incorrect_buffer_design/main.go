package main

import "sync"

type Buffer struct {
	mutex sync.Mutex
	data  []int
}

func NewBuffer() *Buffer {
	return &Buffer{}
}

func (b *Buffer) Add(value int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.data = append(b.data, value)
}

func (b *Buffer) Data() []int {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	// PIZDETC: нельзя просто возвращать ссылку на элемент структуры, все начнут ее юзать
	//return b.data
	// нада отдать копию:

	var sl []int
	copy(sl, b.data)

	return sl
}

func (b *Buffer) ForEach(action func(int)) { // Или вот так через лямбду ебаную
	if action == nil {
		return
	}

	b.mutex.Lock()
	defer b.mutex.Unlock()

	for _, val := range b.data {
		action(val)
	}
}

// Lesson #3 video, time: 01:09:36
