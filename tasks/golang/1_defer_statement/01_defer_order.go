package main

import "fmt"

func main() {
	// 3 -> 2 -> 1
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}
