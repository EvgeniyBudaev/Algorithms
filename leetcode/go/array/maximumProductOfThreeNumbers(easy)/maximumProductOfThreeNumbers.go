package main

import (
	"fmt"
	"sort"
)

/* 628. Maximum Product of Three Numbers
https://leetcode.com/problems/maximum-product-of-three-numbers/description/

Дан целочисленный массив nums. Найдите три числа, произведение которых максимально, и верните максимальное произведение.

Input: nums = [1,2,3]
Output: 6

Input: nums = [1,2,3,4]
Output: 24

Input: nums = [-1,-2,-3]
Output: -6

Input: nums = [-100,-98,-1,2,3,4]
Output: 39200
*/

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(maximumProduct(nums)) // 6
}

// maximumProduct возвращает максимальное произведения трех чисел в массиве nums.
// time: O(n * log(n)), space: O(1)
func maximumProduct(nums []int) int {
	sort.Ints(nums) // Сортируем массив в порядке возрастания
	n := len(nums)  // Получаем длину массива

	// Вариант 1: три самых больших числа
	option1 := nums[n-1] * nums[n-2] * nums[n-3]

	// Вариант 2: два самых маленьких (возможно отрицательных) и самое большое
	option2 := nums[0] * nums[1] * nums[n-1]

	// Возвращаем максимальное из двух вариантов
	return max(option1, option2)
}
