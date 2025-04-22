package main

import "fmt"

/* 704. Binary Search
https://leetcode.com/problems/binary-search/description/

Учитывая массив целых чисел nums, отсортированный в порядке возрастания, и целочисленную цель, напишите функцию для
поиска цели в nums. Если цель существует, верните ее индекс. В противном случае верните -1.
Вы должны написать алгоритм со сложностью выполнения O(log n).

Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4

Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1
*/

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(nums, 9)) // 4
}

// search возвращает индекс элемента массива или -1, если элемент отсутствует в массиве.
func search(nums []int, target int) int {
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

	return -1
}
