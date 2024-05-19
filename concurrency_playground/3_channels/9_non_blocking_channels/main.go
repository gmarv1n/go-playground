package main

import "fmt"

func tryToReadFromChannel(ch chan string) (string, bool) {
	select {
	case value := <-ch:
		// если есть что читать из канала - то этот кейс
		return value, true
	default:
		// если с канала нечего читать - выход из функции
		return "", false
	}
}

func tryToWriteToChannel(ch chan string, value string) bool {
	select {
	case ch <- value:
		// если в канале есть куда писать - вернется тру и запишется
		return true
	default:
		// если нет - то нет
		return false
	}
}

func tryToReadOrWriteToChannel(ch1 chan string, ch2 chan string) {
	select {
	case <-ch1:
		fmt.Println("read") // либо успешно прочитается
	case ch2 <- "test":
		fmt.Println("write") // либо успешно запишется
		//default:
		//	// без дефолта будет паника
		//	fmt.Println("done")
	}
}

func main() {
	//ch := make(chan string)
	//close(ch)

	//tryToWriteToChannel(ch, "test") // panic

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		tryToReadOrWriteToChannel(ch1, ch2)
	}() // вот так нет дедлока, видимо там канал подотрется или я хз не понял

	tryToReadOrWriteToChannel(ch1, ch2) // вот так есть

}

// Lesson #5 video, time: __:__:__
