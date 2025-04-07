package main

import "fmt"

/* 448. Find All Numbers Disappeared in an Array
https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/

Учитывая массив nums из n целых чисел, где nums[i] находится в диапазоне [1, n], верните массив всех целых чисел в
диапазоне [1, n], которые не встречаются в nums.

Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]

Input: nums = [1,1]
Output: [2]
*/

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums)) // [5,6]
}

// findDisappearedNumbers - находит пропущенные числа в массиве nums.
func findDisappearedNumbers(nums []int) []int {
	numSet := make(map[int]bool)
	var result []int

	// Отмечаем все существующие номера на карте
	for _, num := range nums {
		numSet[num] = true
	}

	// Проверяем наличие пропущенных цифр от 1 до n
	for i := 1; i <= len(nums); i++ {
		if !numSet[i] {
			result = append(result, i)
		}
	}

	return result
}
