package main

import (
	"fmt"
	_ "runtime/pprof"
	"time"
)

func main() {
	// GENERATOR PATTERN
	generatorCount := 1
	genCh := generatorFunc("Hello, number!", generatorCount)
	for i := 1; i <= generatorCount; i++ {
		fmt.Println("Generated:", <-genCh)
	}

	// FAN IN PATTERN
	fanIn1Cnt := 3
	fanIn2Cnt := 4

	fanIn := fanInFunc(
		generatorFunc("Hello from gen 1!", fanIn1Cnt),
		generatorFunc("Hello from gen 2!", fanIn2Cnt),
	)

	for i := 1; i <= fanIn1Cnt+fanIn2Cnt; i++ {
		fmt.Println("Generated:", <-fanIn)
	}

	// MESSAGE PATTERN
	genWithMessage := fanInFuncWithMessage(
		generatorFuncWithMsh(Message{msg: "First msg"}, 2),
		generatorFuncWithMsh(Message{msg: "Second msg"}, 2),
	)

	for i := 0; i < 2; i++ {
		msg1 := <-genWithMessage
		fmt.Println(msg1.msg)
		msg2 := <-genWithMessage
		fmt.Println(msg2.msg)

		msg1.waitCh <- true
		msg2.waitCh <- true
	}
}

// GENERATOR PATTERN
func generatorFunc(msg string, count int) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 1; i <= count; i++ {
			ch <- fmt.Sprintf("Msg: \"%s\", counter: %d", msg, i)

			time.Sleep(time.Second * 1)
		}
	}()

	return ch
}

// FAN IN PATTERN
func fanInFunc(chan1, chan2 <-chan string) <-chan string {
	ch := make(chan string)

	go func() {
		for {
			ch <- <-chan1
		}
	}()
	go func() {
		for {
			ch <- <-chan2
		}
	}()

	return ch
}


// MESSAGE PATTERN
type Message struct {
	msg    string
	waitCh chan bool
}

func fanInFuncWithMessage(chan1, chan2 <-chan Message) <-chan Message {
	ch := make(chan Message)

	go func() {
		for {
			ch <- <-chan1
		}
	}()
	go func() {
		for {
			ch <- <-chan2
		}
	}()

	return ch
}

func generatorFuncWithMsh(msg Message, count int) <-chan Message {
	ch := make(chan Message)
	waitCh := make(chan bool)
	go func() {
		for i := 1; i <= count; i++ {
			ch <- Message{
				msg:    fmt.Sprintf("Msg: \"%s\", counter: %d", msg.msg, i),
				waitCh: waitCh,
			}

			time.Sleep(time.Second * 1)

			<-waitCh

			fmt.Println(fmt.Sprintf("DONE msg: %s, Iteration: %d", msg.msg, i))
		}
	}()

	return ch
}
