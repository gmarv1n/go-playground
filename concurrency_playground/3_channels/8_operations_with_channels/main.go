package main

import (
	"fmt"
	"time"
)

func writeToNilChannel() {
	var ch chan int
	ch <- 1
}

func readFromNilChannel() {
	var ch chan int
	<-ch
}

func closeNilChannel() {
	var ch chan int
	close(ch)
}

func rangeNilChannel() {
	var ch chan int
	for val := range ch {
		fmt.Println(val)
	}
}

func openNilChannel() {
	var ch chan int
	// сначала создаем нил чанал

	go func() {
		ch = make(chan int)
		ch <- 100 // вот тут горутина залочится, но ниже есть чтение
		close(ch)
	}()
	// в горутине отложенно создаем, пишем и закрываем

	time.Sleep(100 * time.Millisecond)
	// ждем 100мс

	for value := range ch {
		fmt.Println(value)
	}
	// читаем
	fmt.Println(<-ch)
	// читаем
}

func closeChannelAnyTimes() {
	ch := make(chan int)
	close(ch)
	close(ch)
}

func readFromChannel() {
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20

	val, ok := <-ch
	fmt.Println(val, ok)
	// 10, true

	close(ch)
	val, ok = <-ch
	fmt.Println(val, ok)
	// 20, true

	// ch do not have data
	val, ok = <-ch
	fmt.Println(val, ok)
	// 0, false
}

func readAnyChannels() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 100
	}()

	go func() {
		ch2 <- 200
	}()

	select {
	case val1 := <-ch1:
		fmt.Println(val1)
	case val2 := <-ch2:
		fmt.Println(val2)
	}
}

func writeToClosedChannel() {
	ch := make(chan int, 2)
	ch <- 10

	close(ch)
	ch <- 20
}

func writeToClosedBufferedChannel() {
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20

	go func() {
		ch <- 30
	}()

	time.Sleep(100 * time.Millisecond)
	close(ch)

	for value := range ch {
		fmt.Println(value)
	}
}

func getEventAfterClose() {
	ch := make(chan int, 2)
	go func() {
		<-ch
		fmt.Println("event 1")
	}()

	time.Sleep(100 * time.Millisecond)
	close(ch)

	<-ch
	fmt.Println("event 2")
}

func main() {
	//writeToNilChannel()
	// deadlock - горутина пишет в nil channel, который никто не прочитает
	// и при этом заблокируется

	//readFromNilChannel()
	// deadlock - то же что и в случае с тем чтобы писать в nil канал

	//closeNilChannel()
	// panic - закрывать nil канал нельзя

	//rangeNilChannel()
	// deadlock - там под капотом просто блокирующее чтение

	//openNilChannel()
	// все ок пока есть time.Sleep(100 * time.Millisecond), чтобы горутина
	// с созджанием и тд успела отработать

	//closeChannelAnyTimes()
	// panic - закрываем закрытый канал

	//readFromChannel()
	// тут важно что значение ,ok будет false
	// толкьо после того как будет вычитано все что было записано в канал

	//readAnyChannels()
	// выведится из какого то одного канала

	//writeToClosedChannel()
	// PANIC - запись в закрытый канал

	//writeToClosedBufferedChannel()
	// в основном паника, но иногда функция может завершиться раньше
	// чем заблокированная горутина попытается записать, тогда без паники

	getEventAfterClose()
	// восновном только евент 2, но можети успеть пролезть и евент 1
}

// Lesson #5 video, time: 00:48:00
