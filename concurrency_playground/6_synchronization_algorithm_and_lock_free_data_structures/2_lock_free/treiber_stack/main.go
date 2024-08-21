package main

import (
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

type item struct {
	value int
	next  unsafe.Pointer
}

type Stack struct {
	head unsafe.Pointer
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(value int) {
	// создаем новую ноду
	node := &item{value: value}

	for {
		// атомарно загружаем первую ноду из стека
		head := atomic.LoadPointer(&s.head)

		// устанавливаем новой ноде некст - загруженный хед
		node.next = head

		// запускаем CAS - если голова сейчас все еще та же
		// то свапаем ее атомиком на новую ноду
		// если уже нет, то идем опять в лупе
		if atomic.CompareAndSwapPointer(&s.head, head, unsafe.Pointer(node)) {
			return
		}
	}
}

func (s *Stack) Pop() int {
	for {
		// атомарно грузим хед
		head := atomic.LoadPointer(&s.head)

		if head == nil {
			// типа ошибка
			return -1
		}

		// атомарно грузим некст у хеда
		next := atomic.LoadPointer(&(*item)(head).next)

		// запускаем CAS - если хед все еще хед
		// то заменаем хед на некст из того что мы только что загрузили
		// в некс, если нет то в очередной луп
		if atomic.CompareAndSwapPointer(&s.head, head, next) {
			return (*item)(head).value
		}
	}
}

func main() {
	stack := NewStack()

	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 50; i++ {
		go func(value int) {
			defer wg.Done()
			stack.Push(value)
			stack.Push(value)
			stack.Push(value)
		}(i)
	}

	time.Sleep(100 * time.Millisecond)

	for i := 0; i < 50; i++ {
		go func(value int) {
			defer wg.Done()
			stack.Pop()
			stack.Pop()
			stack.Pop()
		}(i)
	}

	wg.Wait()
}

// Фишка тут в том, что нет блокирующих операций,
// ни мютексов, ни каналов а это значит
// что горутины не будут лочиться

// Lesson #8 video, time: 00:22:46
