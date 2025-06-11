package main

import (
	"fmt"
)

/* Invert values
https://www.codewars.com/kata/5899dc03bc95b1bf1b0000ad/train/gogo

Дано множество чисел, вернуть аддитивное обратное для каждого. Каждое положительное становится отрицательным,
а отрицательное становится положительным.

[1, 2, 3, 4, 5] --> [-1, -2, -3, -4, -5]
[1, -2, 3, -4, 5] --> [-1, 2, -3, 4, -5]
[] --> []
*/

func main() {
	arr1 := []int{1, -2, 3, -4, 5}
	fmt.Println(Invert(arr1)) // [-1, 2, -3, 4, -5]

	arr2 := []int{1, 2, 3, 4, 5}
	fmt.Println(Invert(arr2)) // [-1, -2, -3, -4, -5]

	fmt.Println(Invert(nil)) // []
}

// Invert возвращает обратное значение для каждого элемента массива.
// time: O(n), space: O(n)
func Invert(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	inverted := make([]int, len(arr))
	for i, num := range arr {
		inverted[i] = -num
	}

	return inverted
}
