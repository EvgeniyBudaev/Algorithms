package main

import "fmt"

/* https://leetcode.com/problems/longest-increasing-subsequence/description/

Учитывая целочисленный массив nums, верните длину самого длинного, строго увеличивающегося последовательность.

Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Пояснение: Самая длинная возрастающая подпоследовательность — [2,3,7,101], поэтому длина равна 4.

Input: nums = [0,1,0,3,2,3]
Output: 4

Input: nums = [7,7,7,7,7,7,7]
Output: 1
*/

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})) // 4
}

// lengthOfLIS находит длину наибольшей возрастающей подпоследовательности
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Создаем массив для хранения длин НВП для каждой позиции
	result := make([]int, len(nums))
	for i := range result {
		result[i] = 1 // Каждый элемент сам по себе является подпоследовательностью длины 1
	}

	for i := 0; i < len(nums); i++ {
		// Проверяем все предыдущие элементы
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				// Обновляем длину НВП для текущей позиции
				result[i] = max(result[i], result[j]+1)
			}
		}
	}

	// Находим максимальное значение в массиве result
	maxLen := 0
	for _, v := range result {
		if v > maxLen {
			maxLen = v
		}
	}

	return maxLen
}
