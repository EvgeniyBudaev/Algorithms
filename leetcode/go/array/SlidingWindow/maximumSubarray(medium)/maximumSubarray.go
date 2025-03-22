package main

import "fmt"

/* https://leetcode.com/problems/maximum-subarray/description/

Учитывая целочисленный массив чисел, найдите подмассив с наибольшей суммой и вернуть ее сумму.

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: Подмассив [4,-1,2,1] имеет наибольшую сумму 6.

Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.

Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
*/

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums)) // 6
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	currentSum := maxSum

	for i := 1; i < len(nums); i++ {
		// Обновляем текущую сумму: либо продолжаем подмассив, либо начинаем новый
		currentSum = max(currentSum+nums[i], nums[i])
		// Обновляем максимальную сумму, если текущая сумма больше
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}
