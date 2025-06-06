package main

import "fmt"

/* 1314. Matrix Block Sum
https://leetcode.com/problems/matrix-block-sum/description/

Дана матрица mat размером m x n и целое число k, вернуть ответ матрицы, где каждый ответ[i][j] является суммой
всех элементов mat[r][c] для:

i - k <= r <= i + k,
j - k <= c <= j + k, и
(r, c) является допустимой позицией в матрице.

Input: mat = [[1,2,3],[4,5,6],[7,8,9]], k = 1
Output: [[12,21,16],[27,45,33],[24,39,28]]
*/

func main() {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	k := 1
	fmt.Println(matrixBlockSum(mat, k)) // [[12,21,16],[27,45,33],[24,39,28]]
}

// matrixBlockSum возвращает матрицу размером m x n, где каждый ответ[i][j] является суммой.
// time: O(m*n), space: O(m*n)
func matrixBlockSum(mat [][]int, k int) [][]int {
	r, c := len(mat), len(mat[0]) // r - количество строк, c - количество столбцов
	prefixMat := make([][]int, r) // матрица префиксов
	output := make([][]int, r)    // матрица вывода

	// Инициализируем матрицы префиксов и вывода
	for i := 0; i < r; i++ {
		prefixMat[i] = make([]int, c) // Инициализируем каждый столбец в каждой строке
		output[i] = make([]int, c)    // Инициализируем каждый столбец в каждой строке
	}

	// Создаем матрицу префиксов
	for i := 0; i < r; i++ {
		// Вычисляем сумму элементов в каждой строке
		for j := 0; j < c; j++ {
			// Сумма текущего элемента с предыдущим элементом в столбце
			// Если это первый элемент в столбце, то он не зависит от предыдущего элемента
			if j == 0 {
				prefixMat[i][j] = mat[i][j]
				// Если это не первый элемент в столбце, то сумма зависит от предыдущего элемента
			} else {
				prefixMat[i][j] = prefixMat[i][j-1] + mat[i][j]
			}
		}
	}

	// Создаем матрицу префиксов, используя предыдущие суммы в столбцах
	for j := 0; j < c; j++ {
		// Вычисляем сумму элементов в каждом столбце
		for i := 1; i < r; i++ {
			prefixMat[i][j] += prefixMat[i-1][j]
		}
	}

	var maxRow, maxCol, minRow, minCol int // максимальные и минимальные значения строк и столбцов
	for i := 0; i < r; i++ {               // Перебираем строки
		for j := 0; j < c; j++ { // Перебираем столбцы
			minRow = max(0, i-k)   // Вычисляем минимальное значение строки
			maxRow = min(i+k, r-1) // Вычисляем максимальное значение строки
			minCol = max(0, j-k)   // Вычисляем минимальное значение столбца
			maxCol = min(j+k, c-1) // Вычисляем максимальное значение столбца

			output[i][j] = prefixMat[maxRow][maxCol] // Записываем сумму элементов в текущую ячейку
			// Если есть элементы, которые нужно вычесть, то вычитаем их
			if minRow > 0 {
				output[i][j] -= prefixMat[minRow-1][maxCol]
			}
			// Если есть элементы, которые нужно вычесть, то вычитаем их
			if minCol > 0 {
				output[i][j] -= prefixMat[maxRow][minCol-1]
			}
			// Если есть элементы, которые нужно добавить, то добавляем их
			if minRow > 0 && minCol > 0 {
				output[i][j] += prefixMat[minRow-1][minCol-1]
			}
		}
	}

	// Возвращаем матрицу вывода
	return output
}
