package main

import (
	"fmt"
	_ "runtime/pprof"
)

func recursiveSum(ints []int) int {
	fmt.Print("Call! ")

	if len(ints) <= 1 {
		return ints[0]
	}

	divided := len(ints) / 2

	return recursiveSum(ints[0:divided]) + recursiveSum(ints[divided:])
}

func main() {
	//fmt.Println("\n", recursiveSum([]int{4, 5, 6, 7, 2, 3, 4, 5}))

	fmt.Println("\n", quickSort([]int{8, 3, 4, 4, 7, 6, 9, 1, 5, 12, 54, 2, 65, 8, 6, 65, 43, 11, -5}))

	fmt.Println(recursiveSum([]int{4, 5, 6, 7, 2, 3, 4, 5}))

}
