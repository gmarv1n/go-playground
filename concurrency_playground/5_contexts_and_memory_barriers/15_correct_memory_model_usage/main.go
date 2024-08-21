package main

import (
	"log"
	"runtime"
	"sync/atomic"
)

var a string
var done atomic.Bool

func setup() {
	a = "hello, world"
	done.Store(true) // 1. вот тут атомик, под капотом там полный барьер

	// под капотом атомиков store и load барьер
	// под капотом мьютексов acquire и release барьер

	if done.Load() {
		log.Println(len(a))
	}
}

func main() {
	// go run -race main.go // нет дата рейса
	// но иногда не будет успевать вывестись log.Println(len(a))

	go setup()

	for !done.Load() { // 2. когда тут будет true, точно будет значение в a
		runtime.Gosched() // 2. вот тут прервется main, но точно будет значение в a
	}

	log.Println(a) // expected to print: hello, world
	// 3. а вот тут все хорошо будет
}

// Lesson #8 video, time: 00:00:00
