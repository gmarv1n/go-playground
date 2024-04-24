package main

import "sync"

var resource1 int
var resource2 int

func normalizeResources(lhs, rhs *sync.Mutex) {
	lhs.Lock()
	rhs.Lock()

	// normalization

	rhs.Unlock()
	lhs.Unlock()
}

func main() {
	var mutex1 sync.Mutex
	var mutex2 sync.Mutex

	wg := sync.WaitGroup{}
	wg.Add(1000)

	for i := 0; i < 500; i++ {
		go func() {
			defer wg.Done()
			normalizeResources(&mutex1, &mutex2)
		}()
	}

	//for i := 0; i < 500; i++ {
	//	go func() {
	//		defer wg.Done()
	//		normalizeResources(&mutex2, &mutex1)
	//	}()
	//}
	//
	// ^ нужно согласовывать порядок мьютексов:
	// во втором цикле происходит так что одна горутина
	// лочит один мьютекс, а вторая лочит второй мьютекс
	// и все и пиздец

	for i := 0; i < 500; i++ {
		go func() {
			defer wg.Done()
			normalizeResources(&mutex1, &mutex2)
		}()
	}

	wg.Wait()
}

// Lesson #_ video, time: __:__:__
