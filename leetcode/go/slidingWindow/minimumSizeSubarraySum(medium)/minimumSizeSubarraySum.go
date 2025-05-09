package main

import (
	"fmt"
	"math"
)

/* 209. Minimum Size Subarray Sum
https://leetcode.com/problems/minimum-size-subarray-sum/description/

Учитывая массив положительных целых чисел nums и цель положительного целого числа, верните минимальную длину
подмассива, сумма которых больше или равна целевой. Если такого подмассива нет, вместо него верните 0.

Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Пояснение: Подмассив [4,3] имеет минимальную длину при условии ограничения задачи

Input: target = 4, nums = [1,4,4]
Output: 1

Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0
*/

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(7, nums)) // 2
}

// minSubArrayLen возвращает минимальную длину подмассива, сумма которой больше или равна целевой.
func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	minLength := math.MaxInt32 // Используем максимальное значение для инициализации
	subarraySum := nums[0]     // Начинаем с первого элемента

	for right < len(nums) {
		// Если текущая сумма больше или равна целевой
		if subarraySum >= target {
			// Обновляем минимальную длину, если текущая длина меньше
			minLength = min(minLength, right-left+1)
			// Уменьшаем сумму, убирая левый элемент
			subarraySum -= nums[left]
			// Сдвигаем левый указатель
			left++
		} else {
			// Увеличиваем окно, добавляя правый элемент
			right++
			// Если правый указатель не вышел за пределы массива, добавляем элемент в сумму
			if right < len(nums) {
				subarraySum += nums[right]
			}
		}
	}

	// Если minLength не изменился, возвращаем 0
	if minLength == math.MaxInt32 {
		return 0
	}

	return minLength
}
