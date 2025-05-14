package main

import "fmt"

/* 41. First Missing Positive
https://leetcode.com/problems/first-missing-positive/description/

Дан несортированный целочисленный массив nums. Верните наименьшее положительное целое число, которого нет в nums.
Вы должны реализовать алгоритм, который работает за время O(n) и использует вспомогательное пространство O(1).

Input: nums = [1,2,0]
Output: 3
Объяснение: Все числа в диапазоне [1,2] входят в массив

Input: nums = [3,4,-1,1]
Output: 2
Пояснение: 1 есть в массиве, но 2 отсутствует.

Input: nums = [7,8,9,11,12]
Output: 1
Объяснение: Пропущено наименьшее положительное целое число 1.
*/

func main() {
	nums := []int{1, 2, 0}
	fmt.Println(firstMissingPositive(nums)) // 3
}

// firstMissingPositive - возвращает наименьшее положительное целое число, которого нет в nums.
// time: O(n), memory: O(1)
func firstMissingPositive(nums []int) int {
	n := len(nums) // Длина массива

	// Первый проход: расставляем числа на свои места
	for i := 0; i < n; i++ {
		// Пока число не на своем месте и валидно, перемещаем его
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1] // Меняем местами числа
		}
	}
	// nums: [1, 2, 0]

	// Второй проход: ищем первое число, не соответствующее индексу
	for i := 0; i < n; i++ {
		if nums[i] != i+1 { // Если число не на своем месте, возвращаем его индекс + 1
			return i + 1
		}
	}

	// Если все числа на месте, ответ n+1
	return n + 1
}

// firstMissingPositive - возвращает наименьшее положительное целое число, которого нет в nums.
// time: O(n), memory: O(n)
//func firstMissingPositive(nums []int) int {
//	set := make(map[int]bool)
//
//	for _, num := range nums {
//		if num > 0 {
//			set[num] = true
//		}
//	}
//
//	i := 1 // наименьшее положительное целое число
//	for set[i] {
//		i++
//	}
//
//	return i
//}
