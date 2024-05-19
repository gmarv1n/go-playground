package main

import (
	"fmt"
)

func main() {
	unBufCh()
	bufCh()
}

func bufCh() {
	ch := make(chan int, 1)

	// тут канал буферизированный
	ch <- 100 // 1 вызов при 1 элементе в буфере не будет блокирующим
	// исполгнение выйдет из функции bufCh() и все будет норм

	// у канала можно посмотрел len и cap
	fmt.Println("Buffered Channel:")
	fmt.Println(len(ch))
	fmt.Println(cap(ch))

	// НО если переполнить буфер - то операция уже будет блокирующая и будет
	// дедлок

	//ch <- 100
}

func unBufCh() {
	ch := make(chan int)
	defer close(ch)
	fmt.Println("Not Buffered Channel:")

	go func() {
		val := <-ch
		fmt.Println(val)
	}()

	ch <- 100
	// у канала можно посмотрел len и cap
	fmt.Println(len(ch))
	fmt.Println(cap(ch))

	// deadlock - горутина main ждет пока кто-нибудь не прочитает 100 из канала

	// если сделать так:
	//val := <-ch
	//fmt.Println(val)
	// то тоже дедлок, потому что горутина заблокировалась на ch <- 100

	// нужно до ch <- 100 стартануть горутину которая прочитает, тогда все будет ок
}

// Lesson #5 video, time: 00:16:34
