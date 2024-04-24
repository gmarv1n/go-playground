package main

import "fmt"

func selectionSortAsc(nums []int) []int {
	res := make([]int, 0, len(nums))
	reducedSlice := nums
	counter := 0
	for i := 0; i < len(nums); i++ {
		counter++
		cnt, minVal, nextSlice := minOfSlice(reducedSlice)
		res = append(res, minVal)
		reducedSlice = nextSlice
		counter += cnt
	}

	fmt.Println(fmt.Sprintf("Need %d iterations for slice with %d elements", counter, len(nums)))

	return res
}

func minOfSlice(nums []int) (int, int, []int) {
	minVal := nums[0]
	minIndex := 0

	counter := 0
	for i := 0; i < len(nums); i++ {
		counter++
		if nums[i] < minVal {
			minVal = nums[i]
			minIndex = i
		}
	}

	return counter, minVal, append(nums[:minIndex], nums[minIndex+1:]...)
}

func callSelectionSearch() {
	fmt.Println(selectionSortAsc([]int{5, 3, 6, 2, 1}))
	fmt.Println(selectionSortAsc([]int{6, 3, 7, 9, 2, 55, 68, 1, 0, 65}))
	fmt.Println(selectionSortAsc([]int{6, 3, 7, 9, 2, 55, 68, 1, 0, 65, 43, 6, 22, 77, 3, 2, 34, 9, 1, 5, 4, 3, 2, 6, 8, 45, 3, 25256, 775}))

}
