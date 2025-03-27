package main

import "fmt"

/* https://leetcode.com/problems/find-all-duplicates-in-an-array/description/

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

func findDuplicates(nums []int) []int {
	set := make(map[int]bool) // Используем map для хранения уникальных значений
	result := make([]int, 0)
	for _, num := range nums {
		if set[num] {
			result = append(result, num)
		} else {
			set[num] = true
		}
	}
	return result
}
