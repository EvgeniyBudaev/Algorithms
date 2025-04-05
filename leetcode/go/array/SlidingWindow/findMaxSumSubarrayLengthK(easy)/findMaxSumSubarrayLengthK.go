package main

import "fmt"

/* Fixed window size

Учитывая целочисленный массив nums и целое число k, найдите сумму подмассива с наибольшей суммой, длина которой равна k.
*/

func main() {
	nums := []int{3, -1, 4, 12, -8, 5, 6}
	fmt.Println(findMaxSumSubarrayLengthK(nums, 4)) // 18 <- [3, -1, 4, 12,]
}

// findMaxSumSubarrayLengthK находит сумму подмассива с наибольшей суммой, длина которой равна k.
func findMaxSumSubarrayLengthK(nums []int, k int) int {
	sum := 0 // Текущая сумма подмассива длины k

	// Вычисляем сумму первого окна длины k
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	result := sum // Инициализируем ответ суммой первого окна

	// Скользим окном по массиву
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k] // Добавляем nums[i] и удаляем nums[i-k]
		if sum > result {          // Обновляем ans, если текущая сумма больше
			result = sum
		}
	}

	return result
}
