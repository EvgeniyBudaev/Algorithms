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

// findDuplicate возвращает целое повторяющееся число.
func findDuplicate(nums []int) int {
	visited := make([]bool, len(nums))

	for _, num := range nums {
		if !visited[num] {
			visited[num] = true
		} else {
			return num
		}
	}

	return -1
}

//func findDuplicate(nums []int) int {
//	sort.Ints(nums)
//	for i := 1; i < len(nums); i++ {
//		if nums[i] == nums[i-1] {
//			return nums[i]
//		}
//	}
//	return -1
//}
