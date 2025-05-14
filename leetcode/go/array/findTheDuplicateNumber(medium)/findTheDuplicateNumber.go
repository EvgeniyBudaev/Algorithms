package main

import (
	"fmt"
)

/* 287. Find the Duplicate Number
https://leetcode.com/problems/find-the-duplicate-number/description/

Дан массив целых чисел nums, содержащий n + 1 целых чисел, где каждое целое число находится в диапазоне [1, n]
включительно.
В nums есть только одно повторяющееся число, верните это повторяющееся число.
Вы должны решить проблему, не изменяя числа массива и используя только постоянное дополнительное пространство.

Input: nums = [1,3,4,2,2]
Output: 2

Input: nums = [3,1,3,4,2]
Output: 3
*/

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2})) // 2
}

// findDuplicate - возвращает целое повторяющееся число.
// time: O(n), space: O(1)
func findDuplicate(nums []int) int {
	for _, num := range nums {
		i := abs(num) // индекс в массиве
		// Если элемент уже посещался, значит это повторяющееся число.
		if nums[i] < 0 {
			return i
			// Если элемент еще не посещался, сделаем его отрицательным.
		} else {
			nums[i] = -nums[i]
		}
	}

	// Не найдено повторяющееся число.
	return 0
}

// abs - возвращает абсолютное значение числа.
// time: O(1), space: O(1)
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// findDuplicate - возвращает целое повторяющееся число.
// time complexity: O(n) - проход по массиву.
// space complexity: O(n) - используется вспомогательный массив для хранения посещенных чисел.
//func findDuplicate(nums []int) int {
//	visited := make([]bool, len(nums))
//
//	for _, num := range nums {
//		if !visited[num] {
//			visited[num] = true
//		} else {
//			return num
//		}
//	}
//
//	return -1
//}

// findDuplicate - возвращает целое повторяющееся число.
// time: O(n*log(n)) - сортировка массива.
// space: O(1) - не используется дополнительная память.
//func findDuplicate(nums []int) int {
//	sort.Ints(nums)
//	for i := 1; i < len(nums); i++ {
//		if nums[i] == nums[i-1] {
//			return nums[i]
//		}
//	}
//	return -1
//}
