package main

import (
	"fmt"
	"slices"
)

func main() {

	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))

	fmt.Println(groupAnagrams([]string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"}))
}

func groupAnagrams(strs []string) [][]string {
	res := map[string][]string{}

	for _, str := range strs {
		currSorted := []rune(str)
		slices.Sort(currSorted)

		if _, ok := res[string(currSorted)]; ok {
			res[string(currSorted)] = append(res[string(currSorted)], str)
			continue
		}

		res[string(currSorted)] = []string{str}
	}

	result := [][]string{}

	for _, val := range res {
		result = append(result, val)
	}

	return result
}

// work but slow
func groupAnagramsBad(strs []string) [][]string {
	res := [][]string{}

	for len(strs) > 0 {
		curr := strs[0]

		currSorted := []rune(curr)
		slices.Sort(currSorted)

		foundIndexesMap := make(map[int]bool, len(strs))

		foundIndexesMap[0] = true

		iterRes := make([]string, 0, 2)
		iterRes = append(iterRes, curr)

		for i := 1; i < len(strs); i++ {
			compare := strs[i]
			compareSorted := []rune(compare)

			slices.Sort(compareSorted)

			if string(currSorted) == string(compareSorted) {
				iterRes = append(iterRes, compare)

				foundIndexesMap[i] = true
			}
		}

		newSlice := make([]string, 0, len(strs))
		for i, _ := range strs {
			if ok, _ := foundIndexesMap[i]; !ok {
				newSlice = append(newSlice, strs[i])
			}
		}

		res = append(res, iterRes)
		strs = newSlice
	}

	return res
}
