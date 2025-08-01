package main

import (
	"fmt"
)

/* Factorial
https://www.codewars.com/kata/57a049e253ba33ac5e000212/train/go
*/

func main() {
	fmt.Println(Factorial(4)) // 24
}

// Factorial возвращает факториал числа n.
// time: O(n), space: O(n), n - число
func Factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * Factorial(n-1)
}
