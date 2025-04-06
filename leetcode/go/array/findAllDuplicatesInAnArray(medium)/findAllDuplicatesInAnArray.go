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

// Time O(n)
// Space O(1)
// findDuplicates находит дубликаты чисел.
func findDuplicates(nums []int) []int {
	sort.Ints(nums) // [1, 2, 2, 3, 3, 4, 7, 8]
	var result []int

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			result = append(result, nums[i])
		}
	}

	return result
}

// Invalid Space Complexity
//func findDuplicates(nums []int) []int {
//	set := make(map[int]bool) // Используем map для хранения уникальных значений
//	result := make([]int, 0)
//
//	for _, num := range nums {
//		if set[num] {
//			result = append(result, num)
//		} else {
//			set[num] = true
//		}
//	}
//
//	return result
//}

// Time O(n)
// Space O(1)
//func findDuplicates(nums []int) []int {
//	var result []int
//
//	for _, num := range nums {
//		n := abs(num)
//
//		if nums[n-1] < 0 { // Проверяем, отрицателен ли nums[n-1]
//			// Если да, значит, число n уже встречалось ранее (так как мы инвертировали его в прошлый раз), и это дубликат → добавляем n в результат
//			result = append(result, n)
//		}
//		// Если нет, инвертируем знак (nums[n-1] = -nums[n-1]), чтобы отметить, что число n уже встречалось.
//		nums[n-1] = -nums[n-1]
//	}
//
//	return result
//}
//
//func abs(n int) int {
//	if n < 0 {
//		return -n
//	}
//	return n
//}
