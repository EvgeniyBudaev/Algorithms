package main

import "fmt"

/* 48. Rotate Image
https://leetcode.com/problems/rotate-image/description/

Вам дана матрица n x n 2D, представляющая изображение, поверните изображение на 90 градусов (по часовой стрелке).

Вам нужно повернуть изображение на месте, что означает, что вам нужно напрямую изменить входную матрицу 2D.
Не выделяйте другую матрицу 2D и не выполняйте поворот.

Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]
*/

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	fmt.Println("matrix: ", matrix) // [[7,4,1],[8,5,2],[9,6,3]]
}

// rotate - функция, которая поворачивает матрицу на 90 градусов по часовой стрелке.
// time complexity: O(n^2), space complexity: O(1)
func rotate(matrix [][]int) {
	n := len(matrix) // Размер матрицы

	// Транспонирование матрицы
	for i := 0; i < n; i++ {
		// Перестановка элементов в столбцах
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Переворот каждой строки
	for i := 0; i < n; i++ {
		// Перестановка элементов в строках
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}
