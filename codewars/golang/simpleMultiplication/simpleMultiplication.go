package main

import (
	"fmt"
)

/* Simple multiplication
https://www.codewars.com/kata/583710ccaa6717322c000105/train/go

В этом ката речь идет об умножении заданного числа на восемь, если это четное число, и на девять в противном случае.
*/

func main() {
	fmt.Println(SimpleMultiplication(2))
	fmt.Println(SimpleMultiplication(3))
}

// SimpleMultiplication возвращает произведение заданного числа на 8, если это четное число, и на 9 в противном случае.
// time: O(1), space: O(1)
func SimpleMultiplication(n int) int {
	if n%2 == 0 {
		return n * 8
	}
	return n * 9
}
