package main

import "fmt"

/*
Дан массив целых положительных чисел nums и целое число k. Найдите длину самого длинного подмассива, сумма которого
меньше или равна k.

Input: nums = [3, 1, 2, 7, 4, 2, 1, 1, 5], k = 8
Output: 4
*/

func main() {
	nums := []int{3, 1, 2, 7, 4, 2, 1, 1, 5}
	fmt.Println(findLengthSubarray(nums, 8)) // 4
}

// findLengthSubarray находит длину самого длинного подмассива, сумма которого меньше или равна k.
func findLengthSubarray(nums []int, k int) int {
	left, sum, result := 0, 0, 0

	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		// Если сумма превышает k, сдвигаем left
		for sum > k {
			sum -= nums[left]
			left++
		}
		// Обновляем максимальную длину подмассива
		result = max(result, right-left+1)
	}

	return result
}
