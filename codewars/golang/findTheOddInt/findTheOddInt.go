package main

import (
	"fmt"
)

/* Find the odd int
https://www.codewars.com/kata/54da5a58ea159efa38000836/train/go

Дан массив целых чисел. Найдите число, которое встречается нечётное количество раз.
Всегда будет только одно целое число, которое встречается нечётное количество раз.

Examples:
[7] should return 7, because it occurs 1 time (which is odd).
[0] should return 0, because it occurs 1 time (which is odd).
[1,1,2] should return 2, because it occurs 1 time (which is odd).
[0,1,0,1,0] should return 0, because it occurs 3 times (which is odd).
[1,2,2,3,3,3,4,3,3,3,2,2,1] should return 4, because it appears 1 time (which is odd
*/

func main() {
	seq := []int{1, 1, 2}
	fmt.Println(FindOdd(seq)) // 2
}

// FindOdd принимает массив целых чисел и возвращает число, которое встречается нечётное количество раз.
// time: O(n), space: O(n), n - длина массива
func FindOdd(seq []int) int {
	numMap := make(map[int]int)
	for _, num := range seq {
		numMap[num]++
	}
	for k, v := range numMap {
		if v%2 != 0 {
			return k
		}
	}
	return -1
}
