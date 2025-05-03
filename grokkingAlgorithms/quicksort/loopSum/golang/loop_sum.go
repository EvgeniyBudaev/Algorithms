package main

import "fmt"

func main() {
	fmt.Println(sum([]int{1, 2, 3, 4})) // 10
}

// sum возвращает сумму элементов массива.
func sum(arr []int) int {
	total := 0
	for _, num := range arr {
		total += num
	}
	return total
}
