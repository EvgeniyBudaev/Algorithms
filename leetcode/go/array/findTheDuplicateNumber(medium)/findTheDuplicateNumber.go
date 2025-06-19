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
// time: O(n), space: O(1)
func findDuplicate(nums []int) int {
	for _, num := range nums {
		i := abs(num)    // индекс в массиве
		if nums[i] < 0 { // Если элемент уже посещался, значит это повторяющееся число.
			return i
		} else { // Если элемент еще не посещался, сделаем его отрицательным.
			nums[i] = -nums[i]
		}
	}

	return 0 // Не найдено повторяющееся число.
}

// abs возвращает абсолютное значение числа.
// time: O(1), space: O(1)
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
