package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type DistributedDatabase struct{}

func (d *DistributedDatabase) Query(address string, key string) string {
	time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	return fmt.Sprintf("[%s]: value", address)
}

var database DistributedDatabase

func Query(addresses []string, query string) string {
	result := make(chan string)

	wg := sync.WaitGroup{}
	wg.Add(len(addresses))

	for _, address := range addresses {
		go func(address string) {
			defer wg.Done()
			// defer close(result) // так нельзя - так рано или поздно попадется кейс
			// что горутина вставшая на лочку в записи в канал потом успеет
			// закрыть его и будет паника
			select {
			case result <- database.Query(address, query):
			default:
				return
			}
		}(address)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return <-result
}

func main() {
	addresses := []string{
		"127.0.0.1",
		"127.0.0.2",
		"127.0.0.3",
	}

	value := Query(addresses, "GET key_1")
	fmt.Println(value)
}

// Lesson #6 video, time: __:__:__
