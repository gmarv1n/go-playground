package main

import (
	"fmt"
	"iter"
	_ "runtime/pprof"
)

type IntSeq struct {
	items []int
	iter  iter.Seq[int]
	next  func() (int, bool)
}

func NewIntSeq(items []int) *IntSeq {
	o := &IntSeq{items: items}
	o.iter = count(len(o.items))

	return o
}

func (o *IntSeq) GenerateVal() (int, bool) {
	if o.next == nil {
		next, _ := iter.Pull(o.iter)
		o.next = next
	}

	v, ok := o.next()
	if !ok {
		return 0, false
	}

	return o.items[v], true
}

func count(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range n {
			if !yield(i) {
				break
			}
		}
	}
}

func main() {
	// cd {project_root}
	// GOEXPERIMENT=rangefunc go install ./go_features/iterators
	// GOEXPERIMENT=rangefunc go build -o ./go_features/iterators/ ./go_features/iterators
	// ./go_features/iterators/iterators
	// $home/go/go1.22.5/src/iter/iter.go

	intSeq := NewIntSeq([]int{4, 5, 6})

	fmt.Println(intSeq.GenerateVal()) // 4 true
	fmt.Println(intSeq.GenerateVal()) // 5 true
	fmt.Println(intSeq.GenerateVal()) // 6 true
	fmt.Println(intSeq.GenerateVal()) // 0 false
}
