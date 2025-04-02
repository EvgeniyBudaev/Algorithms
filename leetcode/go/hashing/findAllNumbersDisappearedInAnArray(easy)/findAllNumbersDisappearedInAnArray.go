package main

import "fmt"

/* https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/

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

func findDisappearedNumbers(nums []int) []int {
	numMap := make(map[int]bool)
	var result []int

	// Отмечаем все существующие номера на карте
	for _, num := range nums {
		numMap[num] = true
	}

	// Проверяем наличие пропущенных цифр от 1 до n
	for i := 1; i <= len(nums); i++ {
		if !numMap[i] {
			result = append(result, i)
		}
	}

	return result
}
