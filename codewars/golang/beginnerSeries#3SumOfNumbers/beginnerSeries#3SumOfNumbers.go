package main

import "fmt"

// Beginner Series #3 Sum of Numbers

func main() {
	fmt.Println(GetSum(1, 0))  // 1 (1 + 0 = 1)
	fmt.Println(GetSum(-1, 2)) // 2 (-1 + 0 + 1 + 2 = 2)
}

// GetSum возвращает сумму всех целых чисел между двумя заданными числами (включая их).
// time: O(n), space: O(1)
func GetSum(a, b int) int {
	if a > b {
		return GetSum(b, a)
	}

	sum := 0

	for i := a; i <= b; i++ {
		sum += i
	}

	return sum
}
