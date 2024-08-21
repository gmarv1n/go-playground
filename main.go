package main

import (
	"errors"
	"fmt"
	_ "runtime/pprof"
	"slices"
	"strconv"
	"strings"
)

func recursiveSum(ints []int) int {
	fmt.Print("Call! ")

	if len(ints) <= 1 {
		return ints[0]
	}

	divided := len(ints) / 2

	return recursiveSum(ints[0:divided]) + recursiveSum(ints[divided:])
}

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

func countKeyChanges(s string) int {
	if len(s) == 1 {
		return 0
	}
	start := 0
	end := 1

	cntr := 0

	bt := []byte(s)

	for end < len(bt) {
		first := bt[start]
		second := bt[end]

		if first == second || first == second-32 || first-32 == second {
			start++
			end++
			continue
		}

		cntr++
		start++
		end++
	}

	return cntr
}

func sortPeople(names []string, heights []int) []string {
	talls := map[int]string{}
	for i, height := range heights {
		talls[height] = names[i]
	}
	slices.Sort(heights)
	for i := len(heights) - 1; i >= 0; i-- {
		names[len(heights)-1-i] = talls[heights[i]]
	}

	return names
}

func reversePrefix(word string, ch byte) string {
	prefix := []byte{}
	bytes := []byte(word)

	for i := 0; i < len(word); i++ {
		if bytes[i] == ch {
			prefix = append(prefix, bytes[i])
			prefix = []byte(string(prefix))
			for j := len(prefix) - 1; j >= 0; j-- {
				bytes[len(prefix)-1-j] = prefix[j]
			}
			break
		}
		prefix = append(prefix, bytes[i])
	}
	return string(bytes)
}

func cellsInRange(s string) []string {
	cells := strings.Split(s, ":")

	result := []string{}

	startIndex, _ := strconv.Atoi(string(cells[0][1]))
	endIndex, _ := strconv.Atoi(string(cells[1][1]))

	startChar := cells[0][0]
	endChar := cells[1][0]

	for startChar <= endChar {
		for i := startIndex; i <= endIndex; i++ {
			result = append(result, fmt.Sprintf("%s%d", string(startChar), i))
		}
		startChar += 1
	}

	return result
}

func restoreString(s string, indices []int) string {
	res := make([]byte, len(s))
	for i, v := range indices {
		res[v] = s[i]
	}

	return string(res)
}

func decodeMsg(msg []rune, code []int) {
	res := make([]rune, len(msg))
	for i, v := range code {
		res[v] = msg[i]
	}

	fmt.Println(string(res))
}

type privetFunc func(in string) string

func main() {

	f := func(in string) string {
		fmt.Println(in)

		return in
	}

	d := (privetFunc)(f)("text")

	fmt.Println(d)

}

type SomeType struct {
	ID int
}

func (s *SomeType) Validate() error {
	return errors.New("error occurred")
}

type Validatable interface {
	Validate() error
}

func Validation(val interface{}) error {

	switch v := val.(type) {
	case Validatable:
		return v.Validate()
	default:
		fmt.Println(v)
	}

	return nil
}
