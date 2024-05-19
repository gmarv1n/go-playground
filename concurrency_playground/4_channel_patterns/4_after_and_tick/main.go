package main

import (
	"fmt"
	"time"
)

func main() {
	//for {
	//	select {
	//	case <-time.After(5 * time.Second):
	//		fmt.Println("timeout")
	//		return
	//	case <-time.Tick(time.Second):
	//		fmt.Println("tick")
	//	}
	//}
	// так на каждую итерацию будет создаваться новый
	// таймер и тикер и крутиться эта шляпа будет бесконечно,
	// потому что таймер никогда не открутится

	timerCh := time.After(5 * time.Second)
	tickerCh := time.Tick(time.Second)

	// эти текут, юзать только если рили надо чтобы чтото бесконечно
	// крутилось, иначе NewTicker/Timer с .Close()

	for {
		select {
		case <-timerCh:
			fmt.Println("timeout")
			return
		case <-tickerCh:
			fmt.Println("tick")
		}
	}
}

// Lesson #6 video, time: 00:00:00
