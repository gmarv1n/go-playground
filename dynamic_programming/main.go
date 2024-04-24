package main

import (
	"fmt"
)

func main() {
	// ### Longest substring

	fmt.Println("Longest substring length is:", longestSubstr("fish", "hish"))
	fmt.Println("Longest general substring length is:", longestGeneralSubstr("fish", "fosh"))

	// ### Max bag capacity

	bagCap := 4

	items := []Item{
		{Name: "guitar", Weight: 1, Price: 1500},
		{Name: "audiop", Weight: 4, Price: 3000},
		{Name: "laptop", Weight: 3, Price: 2000},
	}

	findMaxCap(bagCap, items)
}

// ### Longest substring

func longestGeneralSubstr(first, compare string) int {
	out := 0
	cell := make([][]int, len(compare))

	for row := 0; row < len(first); row++ {
		cell[row] = make([]int, len(compare))
		for col := 0; col < len(compare); col++ {
			if row == 0 {
				if first[row] == compare[col] {
					cell[row][col] += 1
					out++
				} else if first[row] != compare[col] && col > 0 {
					cell[row][col] = cell[row][col-1]
				}
				continue
			}
			if first[row] == compare[col] && col > 0 && row > 0 {
				cell[row][col] += cell[row-1][col-1] + 1
				out++
			}

			if first[row] != compare[col] && row > 0 {
				prevLetInRow := cell[row-1][col]
				prevLetInCol := 0

				if col > 0 {
					prevLetInCol = cell[row][col-1]
				}

				maximum := prevLetInRow
				if prevLetInRow < prevLetInCol {
					maximum = prevLetInCol
				}

				cell[row][col] = maximum
			}
		}
	}

	printCell(cell)

	return out
}

func longestSubstr(first, compare string) int {
	out := 0
	cell := make([][]int, len(compare))

	for row := 0; row < len(first); row++ {
		cell[row] = make([]int, len(compare))
		for col := 0; col < len(compare); col++ {
			if row == 0 {
				if first[row] == compare[col] {
					cell[row][col] += 1
				}
				continue
			}
			if first[row] == compare[col] && col > 0 {
				cell[row][col] += cell[row-1][col-1] + 1
				out++
			}
		}
	}

	printCell(cell)

	return out
}

func printCell(cell [][]int) {
	for i := range cell {
		for j := range cell[i] {
			fmt.Print(cell[i][j], " ")
		}
		fmt.Print("\n")
	}
}

// ### Max bag capacity

type Item struct {
	Name   string
	Weight int
	Price  int
}

func findMaxCap(bagCap int, items []Item) {
	table := make([][]int, len(items))
	for row := range table {
		table[row] = make([]int, bagCap)
	}

	weightMap := map[int]Item{}
	for _, item := range items {
		weightMap[item.Weight] = item
	}

	maxPricesItems := map[int][]Item{}

	maxPriceTotal := 0

	for i := 0; i < bagCap; i++ {
		if items[0].Weight <= i+1 {
			table[0][i] = items[0].Price
			maxPricesItems[items[0].Price] = []Item{items[0]}
			maxPriceTotal = items[0].Price
		}
	}

	for row := 1; row < len(items); row++ {
		for col := 0; col < len(table[row]); col++ {
			currCap := col + 1
			itemWeight := items[row].Weight

			prevMaxPrice := table[row-1][col]

			currMaxPrice := 0
			if currCap-itemWeight >= 0 {
				currMaxPrice = items[row].Price

				maxPricesItems[currMaxPrice] = []Item{items[row]}

				if addItem, ok := weightMap[currCap-itemWeight]; ok {
					if currCap-itemWeight != items[row].Weight {
						currMaxPrice += addItem.Price
						maxPricesItems[currMaxPrice] = []Item{items[row], addItem}
					}
				}
			}

			maxPrice := currMaxPrice

			if currMaxPrice < prevMaxPrice {
				maxPrice = prevMaxPrice
			}

			maxPriceTotal = maxPrice
			table[row][col] = maxPrice
		}
	}
	fmt.Println("Table is:")
	printTable(table, items)

	fmt.Println("Best items are:")
	fmt.Println(maxPricesItems[maxPriceTotal])
}

func printTable(tbl [][]int, items []Item) {
	for i := range tbl {
		fmt.Print(items[i].Name, " ")
		for j := range tbl[i] {
			fmt.Print(tbl[i][j], " ")
		}
		fmt.Print("\n")
	}
}
