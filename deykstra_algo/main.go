package main

import (
	"fmt"
	"math"
)

func findShortest(graph map[string]map[string]int, first, end string) string {
	costs := map[string]int{}
	parents := map[string]string{}

	processed := map[string]bool{}

	parents[end] = ""

	costs[end] = math.MaxInt

	processed[first] = true

	for k, v := range graph[first] {
		costs[k] = v
		parents[k] = first
	}

	node := findLowesCostNode(costs, processed)

	for node != "" {
		cost := costs[node]
		neighbours := graph[node]
		for neighbourNode, cst := range neighbours {
			newCost := cost + cst
			if costs[neighbourNode] > newCost {
				costs[neighbourNode] = newCost
				parents[neighbourNode] = node
			}
		}
		processed[node] = true
		node = findLowesCostNode(costs, processed)
	}

	seq := []string{}

	seq = append(seq, end)

	for _, _ = range parents {
		seq = append(seq, parents[seq[len(seq)-1]])
	}

	return fmt.Sprintf("Min cost of going to %s is %d, sequence is: %s", end, costs[end], seq)
}

func findLowesCostNode(costs map[string]int, processed map[string]bool) string {
	lowest := math.MaxInt
	node := ""
	for i, v := range costs {
		if _, ok := processed[i]; ok {
			continue
		}
		if v < lowest {
			lowest = v
			node = i
		}
	}

	return node
}

func main() {
	graph := make(map[string]map[string]int)

	graph["start"] = map[string]int{"a": 6, "b": 2}
	graph["a"] = map[string]int{"finish": 1}
	graph["b"] = map[string]int{"a": 3, "finish": 5}
	graph["finish"] = map[string]int{}

	fmt.Println(findShortest(graph, "start", "finish"))
}
