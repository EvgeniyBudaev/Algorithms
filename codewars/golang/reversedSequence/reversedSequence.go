package main

import (
	"fmt"
)

/* Reversed sequence
https://www.codewars.com/kata/5a00e05cc374cb34d100000d/train/go

Создайте функцию, возвращающую массив целых чисел от n до 1, где n>0.

Пример: n=5 --> [5,4,3,2,1]
*/

func main() {
	fmt.Println(ReverseSeq(5)) // [5,4,3,2,1]
}

// ReverseSeq возвращает массив целых чисел от n до 1, где n>0
// time: O(n), space: O(n), n - длина массива
func ReverseSeq(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = n - i
	}
	return result
}
