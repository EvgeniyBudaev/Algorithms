package main

import "fmt"

/* https://leetcode.com/problems/permutations/description/

Учитывая массив различных целых чисел, верните все возможные перестановки. Вы можете вернуть ответ в любом порядке.

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Input: nums = [0,1]
Output: [[0,1],[1,0]]

Input: nums = [1]
Output: [[1]]
*/

func main() {
	fmt.Println(permute([]int{1, 2, 3})) // [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
}

// permute возвращает все возможные перестановки чисел из входного массива
func permute(nums []int) [][]int {
	var result [][]int
	n := len(nums)

	// Вспомогательная функция для проверки наличия элемента в срезе
	contains := func(slice []int, num int) bool {
		for _, v := range slice {
			if v == num {
				return true
			}
		}
		return false
	}

	// Рекурсивная функция для генерации перестановок
	var backtrack func([]int)
	backtrack = func(current []int) {
		// Если текущая перестановка завершена, добавляем в результат
		if len(current) == n {
			// Создаем копию текущей перестановки
			permutation := make([]int, n)
			copy(permutation, current)
			result = append(result, permutation)
			return
		}

		// Перебираем все числа, которые еще не использованы
		for _, num := range nums {
			if !contains(current, num) {
				// Добавляем число в текущую перестановку
				current = append(current, num)
				// Рекурсивно продолжаем построение перестановки
				backtrack(current)
				// Удаляем последнее число (backtracking)
				current = current[:len(current)-1]
			}
		}
	}

	backtrack([]int{}) // Начинаем с пустой перестановки
	return result
}
