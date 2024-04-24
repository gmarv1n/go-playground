package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	start1 := 0
	start2 := 0

	numsTmp := make([]int, len(nums1))
	copy(numsTmp, nums1)

	for i := 0; i < n+m; i++ {
		if start1 < m && n != 0 && m != 0 && start2 < n {
			if start2 >= n {
				nums1[i] = numsTmp[start1]
				continue
			}
			if numsTmp[start1] < nums2[start2] {
				nums1[i] = numsTmp[start1]
				start1++
				continue
			}

			if numsTmp[start1] >= nums2[start2] {
				nums1[i] = nums2[start2]
				start2++
				continue
			}
		}

		if n != 0 && start2 < n {
			nums1[i] = nums2[start2]
			start2++
			continue
		}

		if m != 0 {
			nums1[i] = numsTmp[start1]
			start1++
			continue
		}
	}

	fmt.Println(nums1)
}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)

	merge([]int{1}, 1, []int{}, 0)

	merge([]int{2, 0}, 1, []int{1}, 1)

	merge([]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3)
}
