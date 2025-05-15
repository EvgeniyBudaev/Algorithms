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

// containsDuplicate - проверяет, содержит ли массив повторяющиеся элементы.
// time: O(n), space: O(n)
func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool) // Мапа для хранения уже встреченных значений

	for _, num := range nums {
		// Если значение уже встречалось, возвращаем true
		if seen[num] {
			return true
		}
		seen[num] = true
	}

	return false // Если ни одно значение не повторилось, возвращаем false
}
