package main

import "fmt"

/* https://leetcode.com/problems/combination-sum-iii/description/

Найдите все допустимые комбинации k чисел, сумма которых равна n, такие, что выполняются следующие условия:
Используются только цифры от 1 до 9.
Каждое число используется не более одного раза.
Возвращает список всех возможных допустимых комбинаций. Список не должен содержать одну и ту же комбинацию дважды,
комбинации могут возвращаться в любом порядке.

Input: k = 3, n = 7
Output: [[1,2,4]]
Объяснение:
1 + 2 + 4 = 7
Других допустимых комбинаций нет

Input: k = 3, n = 9
Output: [[1,2,6],[1,3,5],[2,3,4]]
Объяснение:
1 + 2 + 6 = 9
1 + 3 + 5 = 9
2 + 3 + 4 = 9
Других допустимых комбинаций нет.
*/

func main() {
	fmt.Println(combinationSum3(3, 7)) // [[1,2,4]]
}

// combinationSum3 находит все уникальные комбинации k различных чисел от 1 до 9,
// сумма которых равна n. Каждое число используется не более одного раза.
func combinationSum3(k int, n int) [][]int {
	var result [][]int // Здесь будем хранить все найденные комбинации

	// Внутренняя рекурсивная функция для генерации комбинаций
	var backtrack func(int, int, []int)
	backtrack = func(currentDigit, sum int, elements []int) {
		// Базовые случаи остановки рекурсии
		if currentDigit > 9 || sum > n || len(elements) > k {
			return
		}

		// Если нашли подходящую комбинацию
		if sum == n && len(elements) == k {
			// Создаем копию элементов и добавляем в результат
			combination := make([]int, k)
			copy(combination, elements)
			result = append(result, combination)
			return
		}

		// Перебираем следующие цифры
		for i := currentDigit + 1; i <= 9; i++ {
			// Создаем новый слайс с добавленной цифрой
			newElements := append([]int{}, elements...)
			newElements = append(newElements, i)
			backtrack(i, sum+i, newElements)
		}
	}

	backtrack(0, 0, []int{}) // Начинаем с цифры 0 (чтобы первая была 1), суммы 0 и пустого набора
	return result
}
