package main

import (
	"fmt"
	"strings"
)

/* https://leetcode.com/problems/n-queens/description/

Загадка с n ферзями — это задача о том, как разместить n ферзей на шахматной доске размера n x n так, чтобы никакие два
ферзя не атаковали друг друга.

Учитывая целое число n, найдите все различные решения головоломки с n ферзями. Вы можете вернуть ответ в любом порядке.

Каждое решение содержит отдельную конфигурацию доски для размещения n ферзей, где «Q» и «.». оба указывают на ферзя и
пустое место соответственно.

Input: n = 4
Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
Пояснение: существует два различных решения головоломки с четырьмя ферзями, как показано выше.

Input: n = 1
Output: [["Q"]]
*/

func main() {
	fmt.Println(solveNQueens(4)) // [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
}

func solveNQueens(n int) [][]string {
	// Инициализируем шахматную доску
	chessBoard := make([][]string, n)
	for i := range chessBoard {
		chessBoard[i] = make([]string, n)
		for j := range chessBoard[i] {
			chessBoard[i][j] = "."
		}
	}

	var result [][]string

	// isValidQueen проверяет, можно ли поставить ферзя на позицию (row, col)
	isValidQueen := func(row, col int) bool {
		// Проверяем вертикаль вверх
		for i := 0; i < row; i++ {
			if chessBoard[i][col] == "Q" {
				return false
			}
		}
		// Проверяем диагональ влево-вверх
		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if chessBoard[i][j] == "Q" {
				return false
			}
		}
		// Проверяем диагональ вправо-вверх
		for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if chessBoard[i][j] == "Q" {
				return false
			}
		}
		return true
	}

	// backtrack рекурсивно ищет все возможные расстановки
	var backtrack func(int)
	backtrack = func(row int) {
		if row == n {
			// Добавляем текущую доску в результат
			board := make([]string, n)
			for i := range chessBoard {
				board[i] = strings.Join(chessBoard[i], "")
			}
			result = append(result, board)
			return
		}

		for col := 0; col < n; col++ {
			if isValidQueen(row, col) {
				chessBoard[row][col] = "Q" // Ставим ферзя
				backtrack(row + 1)         // Рекурсивно продолжаем
				chessBoard[row][col] = "." // Убираем ферзя (backtracking)
			}
		}
	}

	backtrack(0) // Начинаем с первой строки
	return result
}
