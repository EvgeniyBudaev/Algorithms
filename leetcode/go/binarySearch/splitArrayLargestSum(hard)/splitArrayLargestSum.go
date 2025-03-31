package main

import (
	"fmt"
	"math"
)

/* 410. Split Array Largest Sum

Учитывая целочисленный массив nums и целое число k, разбейте числа на k непустые подмассивы так, чтобы наибольшая сумма
любого подмассива была минимизирована.
Верните минимизированную наибольшую сумму разделения.
Подмассив — это непрерывная часть массива.

Input: nums = [7,2,5,10,8], k = 2
Output: 18
Объяснение: Существует четыре способа разбить числа на два подмассива.
Лучший способ — разделить его на [7,2,5] и [10,8], где наибольшая сумма среди двух подмассивов равна всего 18.

Input: nums = [1,2,3,4,5], k = 2
Output: 9
Объяснение: Существует четыре способа разбить числа на два подмассива.
Лучший способ — разбить его на [1,2,3] и [4,5], где наибольшая сумма среди двух подмассивов равна всего 9.
*/

func main() {
	nums := []int{7, 2, 5, 10, 8}
	fmt.Println(splitArray(nums, 2)) // 18
}

func splitArray(nums []int, k int) int {
	left := getMax(nums)
	right := sum(nums)
	ans := -1

	for left <= right {
		largestSum := (left + right) / 2

		if countSplitSubarray(largestSum, nums) <= k {
			ans = largestSum
			right = largestSum - 1
		} else {
			left = largestSum + 1
		}
	}

	return ans
}

func countSplitSubarray(largestSum int, nums []int) int {
	countSubarray := 1
	currentSum := 0

	for _, num := range nums {
		currentSum += num

		if currentSum > largestSum {
			countSubarray += 1
			currentSum = num
		}
	}

	return countSubarray
}

func getMax(nums []int) int {
	maxNum := math.MinInt32
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
