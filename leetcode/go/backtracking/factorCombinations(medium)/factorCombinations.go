package main

import "fmt"

/* 254. Factor Combinations

Числа можно рассматривать как произведение своих множителей. Например,
8 = 2 х 2 х 2;
  = 2 х 4.
Напишите функцию, которая принимает целое число n и возвращает все возможные комбинации его множителей.

Input: 1
Output: []

Input: 37
Output:[]

Input: 12
Output: [[2, 6], 2, 2, 3], [3, 4]]
*/

func main() {
	fmt.Println(getFactors(12)) // [[2, 2, 3] [2, 6] [3, 4]]
}

// getFactors находит все возможные комбинации множителей числа n
// (где каждый множитель больше или равен предыдущему)
func getFactors(n int) [][]int {
	var result [][]int // Здесь будем хранить все найденные комбинации множителей

	// Внутренняя рекурсивная функция для поиска множителей
	var backtrack func(int, int, []int)
	backtrack = func(remain int, start int, current []int) {
		// Если остаток равен 1, проверяем и добавляем комбинацию
		if remain == 1 {
			if len(current) > 1 { // Исключаем тривиальный случай [n]
				// Создаем копию текущей комбинации
				combination := make([]int, len(current))
				copy(combination, current)
				result = append(result, combination)
			}
			return
		}

		// Перебираем возможные множители, начиная с 'start'
		for i := start; i <= remain; i++ {
			if remain%i == 0 { // Если i является делителем
				current = append(current, i)       // Добавляем множитель
				backtrack(remain/i, i, current)    // Рекурсивно ищем множители для остатка
				current = current[:len(current)-1] // Удаляем последний множитель (backtracking)
			}
		}
	}

	backtrack(n, 2, []int{}) // Начинаем с числа n, минимального множителя 2 и пустого списка
	return result
}
