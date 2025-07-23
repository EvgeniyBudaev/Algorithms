package main

import "fmt"

func main() {
	first := make([]int, 10)
	fmt.Println(len(first), cap(first)) // 10 10

	//second := first[8:9]
	//fmt.Println(second, len(second), cap(second)) // [0] 1 2
	//second = append(second, 5)
	//fmt.Println(first[9]) // 5

	third := first[8:9:9]
	fmt.Println(third, len(third), cap(third)) // [0] 1 1

	third = append(third, 5)
	fmt.Println(first[9]) // 0
}
