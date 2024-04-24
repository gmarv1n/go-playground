package main

func main() {
	//removeElement([]int{3, 2, 2, 3}, 3)
	//removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	//removeElement([]int{1}, 1)
	removeElement([]int{2, 2, 2}, 0)

}

func removeElement(nums []int, val int) int {
	if len(nums) == 1 && val == nums[0] {
		return len(nums[:0])
	}

	if len(nums) < 2 {
		return len(nums)
	}
	
	start := 0
	end := len(nums) - 1

	idx := map[int]bool{}

	for {
		if start > end {
			break
		}
		if start == end {
			if nums[start] != val {
				idx[start] = true
			}

			break
		}
		if nums[start] != val {
			idx[start] = true
		}
		if nums[end] != val {
			idx[end] = true
		}

		start++
		end--
	}

	cnt := 0
	oldNums := make([]int, len(nums))
	copy(oldNums, nums)

	for i, _ := range idx {
		nums[cnt] = oldNums[i]
		cnt++
	}
	nums = nums[:len(idx)]

	return len(nums)
}
