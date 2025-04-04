package main

import (
	"fmt"
	"math"
)

/*
https://leetcode.com/problems/minimum-falling-path-sum/description/

Дана матрица n x n целых чисел, вернуть минимальную сумму любого нисходящего пути через матрицу.

Нисходящий путь начинается с любого элемента в первой строке и выбирает элемент в следующей строке, который находится
либо непосредственно под ним, либо по диагонали слева/справа. В частности, следующим элементом из позиции (row, col)
будет (row + 1, col - 1), (row + 1, col) или (row + 1, col + 1).

Input: matrix = [[2,1,3],[6,5,4],[7,8,9]]
Output: 13
Пояснение: Существуют два падающих пути с минимальной суммой, как показано на рисунке.
*/

func main() {
	matrix := [][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}
	fmt.Println(minFallingPathSum(matrix)) // 13 (путь 1 → 5 → 7)
}

// minFallingPathSum вычисляет минимальную сумму падающего пути в матрице
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}

	// Создаем копию матрицы для хранения промежуточных результатов
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		copy(dp[i], matrix[i])
	}

	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			// Инициализируем минимальную сумму как текущую ячейку + значение сверху
			minSum := dp[i-1][j] + matrix[i][j]

			// Проверяем левую верхнюю диагональ (если существует)
			if j > 0 {
				minSum = int(math.Min(float64(minSum), float64(dp[i-1][j-1]+matrix[i][j])))
			}

			// Проверяем правую верхнюю диагональ (если существует)
			if j < n-1 {
				minSum = int(math.Min(float64(minSum), float64(dp[i-1][j+1]+matrix[i][j])))
			}

			dp[i][j] = minSum
		}
	}

	// Находим минимальное значение в последней строке
	minTotal := dp[n-1][0]
	for j := 1; j < n; j++ {
		if dp[n-1][j] < minTotal {
			minTotal = dp[n-1][j]
		}
	}

	return minTotal
}
