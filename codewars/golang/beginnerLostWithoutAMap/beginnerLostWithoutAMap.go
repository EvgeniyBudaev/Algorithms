package main

import "fmt"

/* Beginner - Lost Without a Map
https://www.codewars.com/kata/57f781872e3d8ca2a000007e/train/go

Дан массив целых чисел, вернуть новый массив, в котором каждое значение удвоено.

[1, 2, 3] --> [2, 4, 6]
*/

func main() {
	x := []int{1, 2, 3}
	fmt.Println(Maps(x)) // [2, 4, 6]
}

// Maps возвращает новый массив, в котором каждое значение удвоено.
// time: O(n), space: O(n), n - длина массива
func Maps(x []int) []int {
	result := make([]int, len(x))
	for i, num := range x {
		result[i] = num * 2
	}
	return result
}
