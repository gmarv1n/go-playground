package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	const runtime = 1 * time.Second

	greedyWorker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			mutex.Lock()
			time.Sleep(3 * time.Nanosecond)
			mutex.Unlock()
			count++
		}

		fmt.Printf("Greedy worker was able to exec %v work loops\n", count)
	}

	politeWorker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			mutex.Lock()
			time.Sleep(1 * time.Nanosecond)
			mutex.Unlock()

			mutex.Lock()
			time.Sleep(1 * time.Nanosecond)
			mutex.Unlock()

			mutex.Lock()
			time.Sleep(1 * time.Nanosecond)
			mutex.Unlock()

			// в данном случае тк этой горутине нужно несколько раз заюзать ресурсы
			// (лочить разлочить мютекс), в итоге она выполнится меньше раз, чем
			// та которая берет 1 лочку на 3 наносек. Такое голодание может быть
			// и с ресурсами цпу и с памятью и с дескрипторами и с коннектами к бд...

			count++
		}

		fmt.Printf("Polite worker was able to exec %v work loops\n", count)
	}

	wg.Add(2)
	go greedyWorker()
	go politeWorker()

	wg.Wait()
}

// Lesson #_ video, time: __:__:__
