package main

import "fmt"

/* https://leetcode.com/problems/word-search/description/

Учитывая сетку символов m x n и строковое слово, верните true, если слово существует в сетке.
Слово может быть составлено из букв последовательно соседних ячеек, где соседние ячейки являются соседними по
горизонтали или по вертикали. Одна и та же буквенная ячейка не может использоваться более одного раза.

Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true
*/

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word)) // true
}

// exist проверяет, существует ли слово в матрице символов
func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])

	// Внутренняя рекурсивная функция для поиска слова
	var backtrack func(int, int, int) bool
	backtrack = func(i, j, k int) bool {
		// Если прошли все буквы слова - успех
		if k == len(word) {
			return true
		}
		// Проверка границ и совпадения символа
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != word[k] {
			return false
		}

		// Временно заменяем символ, чтобы избежать повторного использования
		temp := board[i][j]
		board[i][j] = 0

		// Проверяем все 4 направления
		result := backtrack(i+1, j, k+1) ||
			backtrack(i-1, j, k+1) ||
			backtrack(i, j+1, k+1) ||
			backtrack(i, j-1, k+1)

		// Восстанавливаем исходное значение
		board[i][j] = temp
		return result
	}

	// Перебираем все клетки как стартовые позиции
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}

	return false
}
