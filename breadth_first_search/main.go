package main

import "fmt"

func breadthFirstSearch(name string, findName string, graph map[string][]string) {
	searched := map[string]bool{}
	queue := []string{}

	queue = append(queue, name)

	personsChecked := 0

	for len(queue) > 0 {
		if _, ok := searched[queue[0]]; ok {
			queue = queue[1:]
			continue
		}
		if queue[0] != findName {
			queue = append(queue, graph[queue[0]]...)
			searched[queue[0]] = true
			queue = queue[1:]
			personsChecked++
			continue
		}
		if queue[0] == findName {
			fmt.Println(fmt.Sprintf(
				"Success: You've started from %s, and %s foud! Persons Checked: %d",
				name,
				findName,
				personsChecked,
			))
			return
		}
	}

	fmt.Println(fmt.Sprintf("Fail: You've started from %s, and %s was not foud!", name, findName))
}

func main() {
	graph := map[string][]string{}
	graph["Vasya"] = []string{"Kolya", "Gena", "Serezha"}
	graph["Kolya"] = []string{"Petya", "Vasgen"}
	graph["Gena"] = []string{"Ksusha", "Natasha", "Oleg"}
	graph["Serezha"] = []string{"Vera", "Igor"}
	graph["Natasha"] = []string{"Sasha", "Vlad"}
	graph["Vasgen"] = []string{"Ashot", "Alisa"}

	breadthFirstSearch("Vasya", "Ashot", graph)
	breadthFirstSearch("Vasya", "Dima", graph)
}
