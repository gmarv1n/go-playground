package main

// тут функция принимает канал только для записи
// если сделать так: <-chan то будет ошибка
func f(out chan<- int) {
	out <- 100
	close(out) // закрыть канал можно если канал на запись

	// если передать аргументом канал только на чтение <-chan
	// то ни писать, ни закрывать будет нельзя
}

func main() {
	var ch = make(chan int)
	f(ch)

	//можно создать канал только на запись
	var ch2 = make(chan<- int)
	_ = ch2

	// можно создать канал только на чтение
	var ch3 = make(<-chan int)
	_ = ch3

	// но тогда нельзя будет или прочитать из него
	// или нельзя будет писать в него
	// зачем это вообще надо неизвестно
}

// Lesson #5 video, time: __:__:__
