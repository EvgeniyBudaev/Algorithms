package main

import "fmt"

/* https://leetcode.com/problems/permutations-ii/description/

Учитывая набор чисел nums, который может содержать дубликаты, возвращает все возможные уникальные перестановки в любом
порядке.

Input: nums = [1,1,2]
Output: [[1,1,2], [1,2,1], [2,1,1]]

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2})) // [[1,1,2], [1,2,1], [2,1,1]]
}

// permuteUnique возвращает все уникальные перестановки чисел с учетом дубликатов
func permuteUnique(nums []int) [][]int {
	var result [][]int

	// Внутренняя рекурсивная функция для генерации перестановок
	var backtrack func(int)
	backtrack = func(start int) {
		// Если дошли до конца массива, добавляем текущую перестановку
		if start == len(nums) {
			// Создаем копию текущей перестановки
			permutation := make([]int, len(nums))
			copy(permutation, nums)
			result = append(result, permutation)
			return
		}

		// Множество для отслеживания уже использованных чисел на текущей позиции
		seen := make(map[int]bool)

		for i := start; i < len(nums); i++ {
			// Пропускаем дубликаты
			if seen[nums[i]] {
				continue
			}
			seen[nums[i]] = true

			// Меняем местами элементы
			nums[start], nums[i] = nums[i], nums[start]
			// Рекурсивно генерируем перестановки для следующей позиции
			backtrack(start + 1)
			// Возвращаем элементы на место (backtracking)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}

	backtrack(0)
	return result
}
