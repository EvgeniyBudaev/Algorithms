package main

import "fmt"

/* https://leetcode.com/problems/number-of-longest-increasing-subsequence/description/

Учитывая целочисленный массив nums, верните количество самых длинных возрастающих подпоследовательностей.
Обратите внимание, что последовательность должна быть строго возрастающей.

Input: nums = [1,3,5,4,7]
Output: 2
Пояснение: Две самые длинные возрастающие подпоследовательности — это [1, 3, 4, 7] и [1, 3, 5, 7].

Input: nums = [2,2,2,2,2]
Output: 5
Пояснение: Длина самой длинной возрастающей подпоследовательности равна 1, существует 5 возрастающих
подпоследовательностей длины 1, поэтому выведите 5.
*/

func main() {
	fmt.Println(findNumberOfLIS([]int{1, 3, 5, 4, 7})) // 2
}

// findNumberOfLIS находит количество самых длинных возрастающих подпоследовательностей
func findNumberOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// lengths - хранит длины НВП, заканчивающихся на каждом индексе
	lengths := make([]int, len(nums))
	// counts - хранит количество НВП для каждой длины
	counts := make([]int, len(nums))

	// Инициализация: каждый элемент сам по себе является подпоследовательностью длины 1
	for i := range lengths {
		lengths[i] = 1
		counts[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				// Случай 1: нашли более длинную подпоследовательность
				if lengths[j]+1 > lengths[i] {
					lengths[i] = lengths[j] + 1
					counts[i] = counts[j]
				} else if lengths[j]+1 == lengths[i] {
					// Случай 2: нашли еще одну подпоследовательность такой же длины
					counts[i] += counts[j]
				}
			}
		}
	}

	// Находим максимальную длину
	maxLength := 0
	for _, l := range lengths {
		if l > maxLength {
			maxLength = l
		}
	}

	// Суммируем количество подпоследовательностей с максимальной длиной
	result := 0
	for i := range counts {
		if lengths[i] == maxLength {
			result += counts[i]
		}
	}

	return result
}
