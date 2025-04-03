package main

import (
	"fmt"
	"sort"
)

/* https://leetcode.com/problems/combination-sum-ii/description/

Учитывая набор номеров кандидатов (кандидатов) и целевое число (цель), найдите все уникальные комбинации среди
кандидатов, в которых сумма чисел кандидатов равна целевому значению.
Каждое число кандидатов можно использовать в комбинации только один раз.
Примечание. Набор решений не должен содержать повторяющихся комбинаций.

Input: candidates = [10,1,2,7,6,1,5], target = 8
Output: [[1,1,6], [1,2,5], [1,7], [2,6]]
*/

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	fmt.Println(combinationSum2(candidates, 8)) // [[1,1,6], [1,2,5], [1,7], [2,6]]
}

// combinationSum2 находит все уникальные комбинации чисел из candidates,
// которые в сумме дают target. Каждое число может использоваться только один раз.
func combinationSum2(candidates []int, target int) [][]int {
	// Сортируем кандидатов для удобства обработки дубликатов
	sort.Ints(candidates)
	var result [][]int

	// Внутренняя рекурсивная функция для поиска комбинаций
	var dfs func(int, int, []int)
	dfs = func(remainingTarget, start int, current []int) {
		// Если достигли целевой суммы
		if remainingTarget == 0 {
			// Добавляем копию текущей комбинации в результат
			temp := make([]int, len(current))
			copy(temp, current)
			result = append(result, temp)
			return
		}

		// Перебираем кандидатов, начиная с указанной позиции
		for i := start; i < len(candidates); i++ {
			// Пропускаем дубликаты (кроме первого вхождения)
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			// Если текущий кандидат превышает оставшуюся цель, прерываем цикл
			if candidates[i] > remainingTarget {
				break
			}

			// Добавляем кандидата в текущую комбинацию
			current = append(current, candidates[i])
			// Рекурсивно ищем комбинации для уменьшенной цели
			dfs(remainingTarget-candidates[i], i+1, current)
			// Удаляем кандидата из комбинации (backtracking)
			current = current[:len(current)-1]
		}
	}

	dfs(target, 0, []int{})
	return result
}
