package main

import (
	"fmt"
)

/* 1. Two Sum
https://leetcode.com/problems/two-sum/description/

Учитывая массив целых чисел nums и целочисленную target, верните индексы двух чисел так, чтобы их сумма составляла
target. Вы можете предположить, что каждый вход будет иметь ровно одно решение, и вы не можете использовать один и тот
же элемент дважды. Вы можете вернуть ответ в любом порядке.

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Input: nums = [3,2,4], target = 6
Output: [1,2]

Input: nums = [3,3], target = 6
Output: [0,1]
*/

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0,1]
}

// twoSum возвращает индексы двух чисел так, чтобы их сумма составляла target.
// time: O(n), space: O(n)
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int) // Создаем map для хранения чисел и их индексов

	for i, num := range nums {
		diff := target - num // Вычисляем разницу между целевым значением и текущим числом
		// Проверяем, есть ли разница в map
		if index, ok := numMap[diff]; ok {
			// Если есть, возвращаем индексы
			return []int{index, i}
		}
		numMap[num] = i // Сохраняем текущее число и его индекс в map
	}

	return nil // Если решение не найдено, возвращаем nil
}
