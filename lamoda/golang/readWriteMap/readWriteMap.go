package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m   = map[string]int{"a": 1}
	rwm sync.RWMutex
)

func main() {
	go read()
	time.Sleep(1 * time.Second)
	go write()
	time.Sleep(1 * time.Second)
}

// read читает из мапы значение "a".
// time: O(1), space: O(1)
func read() {
	for {
		rwm.RLock()
		fmt.Println(m["a"])
		rwm.RUnlock()
	}
}

// write записывает в мапу значение "a".
// time: O(1), space: O(1)
func write() {
	for {
		rwm.Lock()
		m["a"] = 2
		rwm.Unlock()
	}
}
