package main

import "fmt"

/* 78. Subsets
https://leetcode.com/problems/subsets/description/

Учитывая целочисленный массив чисел уникальных элементов, верните все возможные подмножества (набор мощности).
Набор решений не должен содержать повторяющихся подмножеств. Верните решение в любом порядке.

Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Input: nums = [0]
Output: [[],[0]]
*/

func main() {
	fmt.Println(subsets([]int{1, 2, 3})) // [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
}

// subsets возвращает все возможные подмножества (powerset) входного массива.
// time: O(2^n), где n - количество элементов в массиве.
// space: O(2^n), так как результат содержит все возможные подмножества.
func subsets(nums []int) [][]int {
	var result [][]int       // Инициализируем результирующий массив для хранения подмножеств
	var dfs func(int, []int) // Внутренняя рекурсивная функция для генерации подмножеств

	dfs = func(index int, current []int) {
		// Если дошли до конца массива, добавляем текущее подмножество
		if index == len(nums) {
			// Создаем копию текущего подмножества
			subset := make([]int, len(current))
			copy(subset, current)
			result = append(result, subset)
			return
		}

		// Вариант 1: включаем текущий элемент
		dfs(index+1, append(current, nums[index]))

		// Вариант 2: не включаем текущий элемент
		dfs(index+1, current)
	}

	dfs(0, []int{}) // Начинаем с пустого подмножества
	return result
}
