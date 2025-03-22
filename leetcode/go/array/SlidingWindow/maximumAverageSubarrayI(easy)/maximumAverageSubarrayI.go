package main

import "fmt"

/* https://leetcode.com/problems/maximum-average-subarray-i/description/
Fixed window size

Вам дан целочисленный массив nums, состоящий из n элементов, и целое число k.
Найдите непрерывный подмассив длиной k, имеющий максимальное среднее значение, и верните это значение.

Input: nums = [1,12,-5,-6,50,3], k = 4
Output: 12.75000
Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75

Input: nums = [5], k = 1
Output: 5.00000
*/

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	fmt.Println(findMaxAverage(nums, 4)) // 12.75
}

func findMaxAverage(nums []int, k int) float64 {
	sum := 0

	// Вычисляем сумму первых k элементов
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	// Инициализируем максимальное среднее значение
	maxAvg := float64(sum) / float64(k)

	// Используем sliding window для вычисления суммы остальных подмассивов
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k] // Добавляем новый элемент и убираем старый
		avg := float64(sum) / float64(k)
		if avg > maxAvg {
			maxAvg = avg
		}
	}

	return maxAvg
}
