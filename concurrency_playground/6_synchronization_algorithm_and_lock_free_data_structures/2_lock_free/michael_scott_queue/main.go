package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

type item struct {
	value int
	next  unsafe.Pointer
}

type Queue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

func NewQueue() Queue {
	dummy := &item{}

	return Queue{
		head: unsafe.Pointer(dummy),
		tail: unsafe.Pointer(dummy),
	}
}

// Чтобы сделать пуш надо сделать:
// 1. Заменить тейл на новую ноду
// 2. У предыдущей ноды засеттить в некст новую ноду

func (q *Queue) Push(value int) {
	// создаем ноду
	node := &item{value: value}

	for {
		// грузим тейл и его некст
		tail := atomic.LoadPointer(&q.tail)
		next := atomic.LoadPointer(&(*item)(tail).next)

		// если тейл все еще такой же
		if tail == atomic.LoadPointer(&q.tail) {
			if next == nil {
				// если некста у тейла нет, пытаемся CASом
				// поменять у тейла некс на новую ноду
				// это по сути шаг 2
				if atomic.CompareAndSwapPointer(
					&(*item)(tail).next,
					next,
					unsafe.Pointer(node),
				) {
					// если некст у тейла успешно поменялся,
					// то меняем тейл на новую ноду
					atomic.CompareAndSwapPointer(
						&q.tail,
						tail,
						unsafe.Pointer(node),
					)

					// вопрос а если тут неуспешно будет выше?
					// ответ - как будто это значит что тейл уже
					// подменила другая горутина

					return
				}
			} else {
				// а сюда провалится если где то в другой
				// горутине уже успел подмениться некст у тейла
				//
				// пытаемся поменять тейл, который
				// захерачился из другой горутины на новую
				// ноду

				// а чо с некстом у ноды тейл подмененной из
				// другой горутины???
				// ответ - этой операцией мы если че как бдуто
				// доделываем работу другой горутины
				atomic.CompareAndSwapPointer(
					&q.tail,
					tail,
					unsafe.Pointer(node),
				)
			}
		}
	}
}

// чтобы сделать поп, нужно:
// 1. Переставить некст ноду хеда в хед

func (q *Queue) Pop() int {
	for {
		// грузим хед, тейл и некст хеда
		head := atomic.LoadPointer(&q.head)
		tail := atomic.LoadPointer(&q.tail)
		next := atomic.LoadPointer(&(*item)(head).next)

		// если хед еще тот же
		if head == atomic.LoadPointer(&q.head) {
			// если тейл и хед - одна и та же нода
			if head == tail {
				// если некст нил - то типа пустая очередь
				if next == nil {
					return -1
				} else {
					// иначе если ктото уже засеттил некст
					// меняем тейл на загруженный из хеда
					// некст (типа помогаем другой горутине)
					atomic.CompareAndSwapPointer(
						&q.tail,
						tail,
						next,
					)
				}
			} else {
				// если хед не равен тейлу
				// достаем значение некста хеда
				// не понял почему его возвращаем
				// должны же возвращать хед.валуе ??? :(
				value := (*item)(next).value
				// подменяем хед на некст предыдущего хеда
				if atomic.CompareAndSwapPointer(&q.head, head, next) {
					return value
				}
			}
		}
	}
}

func main() {
	queue := NewQueue()

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
}

// Lesson #8 video, time: 00:40:40
