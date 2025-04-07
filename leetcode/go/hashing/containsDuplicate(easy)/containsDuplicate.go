package main

import "fmt"

/* 217. Contains Duplicate
https://leetcode.com/problems/contains-duplicate/description/

Учитывая числа целочисленного массива, верните true, если какое-либо значение встречается в массиве хотя бы дважды, и
верните false, если каждый элемент различен.

Input: nums = [1,2,3,1]
Output: true

Input: nums = [1,2,3,4]
Output: false

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true
*/

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1})) // true
}

// containsDuplicate проверяет, содержит ли массив повторяющиеся элементы.
func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}

	return false
}
