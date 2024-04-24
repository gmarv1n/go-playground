package main

import (
	"fmt"
	_ "runtime/pprof"
)

func quickSort(ints []int) []int {
	fmt.Print("Call! ")

	if len(ints) < 2 {
		return ints
	}

	pivot := len(ints) / 2
	base := ints[pivot]

	less := make([]int, 0)
	greater := make([]int, 0)
	equal := make([]int, 0)

	for _, num := range ints {
		switch {
		case num > base:
			greater = append(greater, num)
		case num < base:
			less = append(less, num)
		case num == base:
			equal = append(equal, num)
		}
	}

	less = append(quickSort(less), equal...)

	return append(less, quickSort(greater)...)
}
