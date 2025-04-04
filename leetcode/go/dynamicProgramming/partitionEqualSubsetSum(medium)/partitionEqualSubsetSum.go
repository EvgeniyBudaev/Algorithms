package main

import "fmt"

/* https://leetcode.com/problems/partition-equal-subset-sum/description/


Учитывая числа целочисленного массива, верните true, если вы можете разделить массив на два подмножества так, чтобы
сумма элементов в обоих подмножествах была равна или false в противном случае.

Input: nums = [1,5,11,5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11].

Input: nums = [1,2,3,5]
Output: false
Explanation: The array cannot be partitioned into equal sum subsets.
*/

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5})) // true (11 = 1 + 5 + 5)
}

// canPartition проверяет, можно ли разбить массив на две части с равной суммой
func canPartition(nums []int) bool {
	// Вычисляем общую сумму элементов
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// Если сумма нечетная - разбиение невозможно
	if totalSum%2 != 0 {
		return false
	}

	target := totalSum / 2
	// Создаем массив для динамического программирования
	dp := make([]bool, target+1)
	dp[0] = true // Базовый случай - сумма 0 всегда достижима

	// Перебираем все числа в массиве
	for _, num := range nums {
		// Идем от целевой суммы вниз до текущего числа
		for j := target; j >= num; j-- {
			// Если можно достичь суммы (j - num), то можно достичь и суммы j
			if dp[j-num] {
				dp[j] = true
			}
		}
	}

	return dp[target]
}
