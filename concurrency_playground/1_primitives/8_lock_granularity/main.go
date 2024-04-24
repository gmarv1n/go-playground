package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var cache map[string]string

func doSmth() {
	mutex.Lock()
	item := cache["key"]
	// fmt.Println(item) // NOT HERE, cause item is on the stack
	mutex.Unlock()
	fmt.Println(item)
}

func main() {

}
