package main

import "fmt"

/* https://leetcode.com/problems/search-insert-position/description/

Учитывая отсортированный массив различных целых чисел и целевое значение, верните индекс, если цель найдена.
Если нет, верните индекс там, где он был бы, если бы он был вставлен по порядку.
Вы должны написать алгоритм со сложностью выполнения O(log n)

Input: nums = [1,3,5,6], target = 5
Output: 2

Input: nums = [1,3,5,6], target = 2
Output: 1

Input: nums = [1,3,5,6], target = 7
Output: 4
*/

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 5)) // 2
}

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2
		guess := nums[mid]

		if guess == target {
			return mid
		}
		if guess > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}
