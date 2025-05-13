package main

import (
	"fmt"
	"sort"
)

/* 442. Find All Duplicates in an Array
https://leetcode.com/problems/find-all-duplicates-in-an-array/description/

Учитывая целочисленный массив nums длины n, где все целые числа nums находятся в диапазоне [1, n] и каждое целое число
встречается один или два раза, верните массив всех целых чисел, которые встречаются дважды.
Вы должны написать алгоритм, который работает за время O(n) и использует только постоянное дополнительное пространство.

Input: nums = [4,3,2,7,8,2,3,1]
Output: [2,3]

Input: nums = [1,1,2]
Output: [1]

Input: nums = [1]
Output: []
*/

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(nums)) // [2,3]
}

// findDuplicates находит дубликаты чисел.
// time: O(n), space: O(1)
func findDuplicates(nums []int) []int {
	sort.Ints(nums)  // [1, 2, 2, 3, 3, 4, 7, 8]
	var result []int // Создаем пустой слайс для хранения дубликатов

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] { // Если текущий элемент равен предыдущему, добавляем его в результат
			result = append(result, nums[i])
		}
	}

	return result // Возвращаем список дубликатов
}
