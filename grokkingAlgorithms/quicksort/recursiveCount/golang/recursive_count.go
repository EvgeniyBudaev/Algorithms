package main

import "fmt"

func main() {
	fmt.Println(count([]int{0, 1, 2, 3, 4, 5})) // 6
}

// count возвращает количество элементов в массиве.
func count(list []int) int {
	if len(list) == 0 {
		return 0
	}
	return 1 + count(list[1:])
}
