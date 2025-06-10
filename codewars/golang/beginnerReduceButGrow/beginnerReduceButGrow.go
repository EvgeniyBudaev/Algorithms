package main

import "fmt"

/* Beginner - Reduce but Grow
https://www.codewars.com/kata/57f780909f7e8e3183000078/train/go

Для заданного непустого массива целых чисел вернуть результат умножения значений по порядку.

Example:
[1, 2, 3, 4] => 1 * 2 * 3 * 4 = 24
*/

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(Grow(arr)) // 24
}

func Grow(arr []int) int {
	result := 1

	for _, num := range arr {
		result *= num
	}

	return result
}
