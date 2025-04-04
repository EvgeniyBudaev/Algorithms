package main

import "fmt"

/* https://leetcode.com/problems/partition-to-k-equal-sum-subsets/description/

Учитывая целочисленный массив nums и целое число k, верните true, если можно разделить этот массив на k непустые
подмножества, все суммы которых равны.

Input: nums = [4,3,2,3,5,2,1], k = 4
Output: true
Explanation: It is possible to divide it into 4 subsets (5), (1, 4), (2,3), (2,3) with equal

Input: nums = [1,2,3,4], k = 3
Output: false
*/

func main() {
	nums := []int{4, 3, 2, 3, 5, 2, 1}
	fmt.Println(canPartitionKSubsets(nums, 4)) // true ([5], [4,1], [3,2], [3,2])
}

// canPartitionKSubsets проверяет, можно ли разбить массив на K подмножеств с равными суммами
func canPartitionKSubsets(nums []int, k int) bool {
	// Вычисляем общую сумму элементов
	total := 0
	for _, num := range nums {
		total += num
	}

	// Если сумма не делится на k - разбиение невозможно
	if total%k != 0 {
		return false
	}

	target := total / k
	visited := make([]bool, len(nums))

	// Вспомогательная рекурсивная функция
	var canPartition func(start, numberOfSubsets, currentSum int) bool
	canPartition = func(start, numberOfSubsets, currentSum int) bool {
		// Базовый случай: все подмножества найдены
		if numberOfSubsets == 1 {
			return true
		}

		// Нашли подмножество, ищем следующее
		if currentSum == target {
			return canPartition(0, numberOfSubsets-1, 0)
		}

		// Перебираем элементы, начиная с индекса start
		for i := start; i < len(nums); i++ {
			if !visited[i] && currentSum+nums[i] <= target {
				visited[i] = true
				// Пробуем добавить текущий элемент в подмножество
				if canPartition(i+1, numberOfSubsets, currentSum+nums[i]) {
					return true
				}
				// Откатываем изменения (backtracking)
				visited[i] = false
			}
		}

		return false
	}

	return canPartition(0, k, 0)
}
