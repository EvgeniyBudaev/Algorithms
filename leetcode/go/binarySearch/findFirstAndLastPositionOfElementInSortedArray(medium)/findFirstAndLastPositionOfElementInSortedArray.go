package main

import (
	"fmt"
	"slices"
	"sort"
)

/* 34. Find First and Last Position of Element in Sorted Array
https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/

Дан массив целых чисел nums, отсортированный в неубывающем порядке, найти начальную и конечную позицию заданного
целевого значения.
Если target не найден в массиве, вернуть [-1, -1].
Необходимо написать алгоритм со сложностью выполнения O(log n).

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
*/

func main() {
	arr := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange(arr, target)) // [3,4]
}

// searchRange - поиск начальной и конечной позиции целевого значения в массиве.
// time complexity: O(log n), space complexity: O(1)
func searchRange(nums []int, target int) []int {
	// Проверка на пустой массив
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	// Поиск стартовой позиции
	start, ok := slices.BinarySearch(nums, target)
	if !ok {
		return []int{-1, -1}
	}

	// Поиск конечной позиции
	end := sort.Search(len(nums), func(i int) bool {
		return nums[i] > target // Поиск позиции, где элемент больше target
	})

	end--

	return []int{start, end}
}
