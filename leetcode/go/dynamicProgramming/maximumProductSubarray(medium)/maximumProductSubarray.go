package main

import (
	"fmt"
	"math"
)

/* https://leetcode.com/problems/maximum-product-subarray/description/

Учитывая целочисленный массив чисел, найдите подмассив который имеет самый большой продукт, и возвращает продукт.
Тестовые случаи генерируются таким образом, чтобы ответ соответствовал 32-битному целому числу.

Input: nums = [2,3,-2,4]
Output: 6
Пояснение: [2,3] имеет наибольшее произведение 6.

Input: nums = [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
*/

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4})) // 6 (подмассив [2,3])
}

// maxProduct находит максимальное произведение подмассива
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Инициализация переменных:
	// maxSoFar - максимальное произведение на текущий момент
	// minSoFar - минимальное произведение на текущий момент (нужно для отрицательных чисел)
	// result - итоговый результат
	maxSoFar := nums[0]
	minSoFar := nums[0]
	result := maxSoFar

	for i := 1; i < len(nums); i++ {
		current := nums[i]

		// Временная переменная для хранения нового maxSoFar перед обновлением
		tempMaxSoFar := math.Max(
			float64(current),
			math.Max(
				float64(maxSoFar*current),
				float64(minSoFar*current),
			),
		)

		// Обновляем minSoFar с учетом текущего числа
		minSoFar = int(math.Min(
			float64(current),
			math.Min(
				float64(maxSoFar*current),
				float64(minSoFar*current),
			),
		))

		// Обновляем maxSoFar
		maxSoFar = int(tempMaxSoFar)

		// Обновляем результат, если нашли большее произведение
		if maxSoFar > result {
			result = maxSoFar
		}
	}

	return result
}
