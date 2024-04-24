package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return len(nums)
	}

	start := 0
	end := 1

	foundDupl := false
	for end < len(nums) {
		if nums[start] == nums[end] {
			if end+1 == len(nums) && start == 0 {
				return removeDuplicates(nums[:1])
			}

			end++
			foundDupl = true

			if end == len(nums) && start != 0 {
				return removeDuplicates(append(nums[:start], nums[end-1:]...))
			}

			continue
		}

		if nums[start] != nums[end] && foundDupl == false {
			if end+1 == len(nums) {
				break
			}
			start++
			end++
		}

		if nums[start] != nums[end] && foundDupl == true {
			return removeDuplicates(append(nums[:start], nums[end-1:]...))
		}
	}

	return len(nums)
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 2, 2}))
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	fmt.Println(removeDuplicates([]int{1, 1, 1}))

	//nums := []int{1, 1, 1}
	//nums = nums[:1]
	//
	//fmt.Println(nums)
}
