package main

import (
	"fmt"
)

/* What is between?
https://www.codewars.com/kata/55ecd718f46fba02e5000029/train/go

Завершите функцию, которая принимает два целых числа (a, b, где a < b) и возвращает массив всех целых чисел между
входными параметрами, включая их самих.

Example:
a = 1
b = 4
--> [1, 2, 3, 4]
*/

func main() {
	fmt.Println(Between(1, 4)) // [1, 2, 3, 4]
}

// Between возвращает массив всех целых чисел между двумя входными параметрами, включая их самих.
// time: O(n), space: O(n), где n - разница между a и b
func Between(a, b int) []int {
	var result []int

	for i := a; i <= b; i++ {
		result = append(result, i)
	}

	return result
}
