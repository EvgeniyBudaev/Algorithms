package main

import "fmt"

/* https://leetcode.com/problems/combinations/description/

Учитывая два целых числа n и k, верните все возможные комбинации k чисел, выбранных из диапазона [1, n].
Вы можете вернуть ответ в любом порядке.

Input: n = 4, k = 2
Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
Пояснение: всего есть 4 комбинации на выбор 2 = 6.
Обратите внимание, что комбинации неупорядочены, т. е. [1,2] и [2,1] считаются одной и той же комбинацией.
*/

func main() {
	fmt.Println(combine(4, 2)) // [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
}

// combine генерирует все возможные комбинации k чисел из диапазона [1, n]
func combine(n int, k int) [][]int {
	var result [][]int // Здесь будем хранить все найденные комбинации

	// Внутренняя рекурсивная функция для генерации комбинаций
	var backtrack func(int, []int)
	backtrack = func(start int, current []int) {
		// Если текущая комбинация достигла нужной длины, добавляем в результат
		if len(current) == k {
			// Создаем копию текущей комбинации, чтобы избежать изменений в дальнейшем
			temp := make([]int, k)
			copy(temp, current)
			result = append(result, temp)
			return
		}

		// Генерируем комбинации, начиная с числа start
		for i := start; i <= n; i++ {
			// Добавляем текущее число в комбинацию
			current = append(current, i)
			// Рекурсивно вызываем для следующих чисел (чтобы избежать повторов)
			backtrack(i+1, current)
			// Удаляем последнее число (backtracking)
			current = current[:len(current)-1]
		}
	}

	backtrack(1, []int{}) // Начинаем с числа 1 и пустой комбинации
	return result
}
