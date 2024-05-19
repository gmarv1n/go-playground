package main

import "sync"

var (
	actions int
	mutex   sync.Mutex
	buffer  chan struct{}
)

func consumer() {
	for i := 0; i < 1000; i++ {
		// 8. опять залочили
		mutex.Lock()
		// 9. инкрементнули
		actions++
		// 2. тут мы ждем на залоченом мутексе
		// 10. тут мы залочились на канале
		<-buffer
		// 3. пролчитали с канала
		mutex.Unlock()
		// 4. разлочили мьютекс
	}
}

func producer() {
	for i := 0; i < 1000; i++ {
		// 1. вот тут он записал в канал и залочился
		// 11. и видимо вот тут мы не можем писать в канал, потому что
		// на шаге 10 мы взяли лочку мьютексом
		buffer <- struct{}{}
		// 5. залочили мютекс
		mutex.Lock()
		actions++
		// 6. заинкрементили
		mutex.Unlock()
		// 7. разлочили
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	buffer = make(chan struct{}, 1)

	// короче тут 2 мьютекса, которые лочат друг друга и изза этого дедлок
	// второй мьютекс внутри канала
	go func() {
		defer wg.Done()
		consumer()
	}()

	go func() {
		defer wg.Done()
		producer()
	}()

	wg.Wait()
}

// Lesson #6 video, time: 00:00:00
