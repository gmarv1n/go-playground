package main

import "fmt"

func binarySearch(list []string, item string) int {
	start := 0
	end := len(list) - 1

	count := 0

	for start <= end {
		count++
		mid := (start + end) / 2
		if list[mid] > item {
			end = mid - 1
		}
		if list[mid] < item {
			start = mid + 1
		}
		if list[mid] == item {

			fmt.Println("count", count)
			fmt.Println("res:", list[mid])

			fmt.Println(fmt.Sprintf("log 2 of %d = %d", countSq(count), count))
			fmt.Println(fmt.Sprintf("u need multiply 2 on itself %d times to get %d", count, countSq(count)))
			return mid
		}
	}

	fmt.Println(fmt.Sprintf("log 2 of %d = %d", countSq(count), count))
	fmt.Println(fmt.Sprintf("u need multiply 2 on itself %d times to get %d", count, countSq(count)))

	return 0
}

func countSq(count int) int {
	n := 1
	for i := 0; i < count; i++ {
		n = n * 2
	}
	return n
}

func callBinSearch() {
	list := []string{}
	for i := 0; i < 1024000; i++ {
		list = append(list, fmt.Sprintf("A%d", i+1))
	}

	binarySearch(list, "A1")
}
