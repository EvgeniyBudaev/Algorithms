package main

import (
	"fmt"
	"sort"
)

/* Sort Numbers
https://www.codewars.com/kata/5174a4c0f2769dd8b1000003/train/go

Завершите решение так, чтобы оно сортировало переданный массив чисел. Если функция принимает пустой массив или значение
null/nil, она должна вернуть пустой массив.

Example:
solution(c(1, 2, 3, 10, 5)) # should return [1, 2, 3, 5, 10]
solution(nil)              # should return []
*/

func main() {
	numbers := []int{1, 2, 10, 50, 5}
	fmt.Println(SortNumbers(numbers)) // [1 2 5 10 50]
	fmt.Println(SortNumbers([]int{})) // []
	fmt.Println(SortNumbers(nil))     // []
}

// SortNumbers сортирует переданный массив чисел.
// time: O(n log n), space: O(n), n - длина массива
func SortNumbers(numbers []int) []int {
	if numbers == nil || len(numbers) == 0 {
		return []int{}
	}
	sorted := make([]int, len(numbers))
	copy(sorted, numbers) // Создаем копию (чтобы не изменять исходный массив)
	sort.Ints(sorted)
	return sorted
}
