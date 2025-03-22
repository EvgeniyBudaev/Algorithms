package main

import "fmt"

/* Fixed window size

Учитывая целочисленный массив nums и целое число k, найдите сумму подмассива с наибольшей суммой, длина которой равна k.
*/

func main() {
	nums := []int{3, -1, 4, 12, -8, 5, 6}
	fmt.Println(findMaxSumSubarrayLengthK(nums, 4)) // 18 <- [3, -1, 4, 12,]
}

func findMaxSumSubarrayLengthK(nums []int, k int) int {
	curr := 0 // Текущая сумма подмассива длины k

	// Вычисляем сумму первого окна длины k
	for i := 0; i < k; i++ {
		curr += nums[i]
	}

	ans := curr // Инициализируем ответ суммой первого окна

	// Скользим окном по массиву
	for i := k; i < len(nums); i++ {
		curr += nums[i] - nums[i-k] // Добавляем nums[i] и удаляем nums[i-k]
		if curr > ans {             // Обновляем ans, если текущая сумма больше
			ans = curr
		}
	}

	return ans
}
