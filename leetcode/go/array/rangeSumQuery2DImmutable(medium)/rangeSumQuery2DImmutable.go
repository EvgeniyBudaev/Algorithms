package main

import "fmt"

/* 304. Range Sum Query 2D - Immutable
https://leetcode.com/problems/range-sum-query-2d-immutable/description/

Дана двумерная матрица matrix, обработайте несколько запросов следующего типа:

Вычислите сумму элементов matrix внутри прямоугольника, определенного его верхним левым углом (row1, col1) и нижним
правым углом (row2, col2).
Реализуйте класс NumMatrix:

NumMatrix(int[][] matrix) Инициализирует объект с целочисленной матрицей matrix.
int sumRegion(int row1, int col1, int row2, int col2) Возвращает сумму элементов matrix внутри прямоугольника,
определенного его верхним левым углом (row1, col1) и нижним правым углом (row2, col2).
Вам необходимо разработать алгоритм, в котором sumRegion работает со сложностью O(1).

Input
["NumMatrix", "sumRegion", "sumRegion", "sumRegion"]
[[[[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]]], [2, 1, 4, 3], [1, 1, 2, 2], [1, 2, 2, 4]]

Output
[null, 8, 11, 12]
*/

func main() {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	obj := Constructor(matrix)
	fmt.Println(obj.SumRegion(2, 1, 4, 3)) // Должно вернуть 8
}

type NumMatrix struct {
	prefixSum [][]int // Матрица префиксных сумм
}

// Constructor - конструктор.
// time: O(n * m)
// memory: O(n * m)
func Constructor(matrix [][]int) NumMatrix {
	// Проверяем на пустую матрицу
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{prefixSum: nil}
	}

	rows := len(matrix)    // Количество строк в матрице
	cols := len(matrix[0]) // Количество столбцов в матрице

	// Создаем матрицу префиксных сумм с дополнительными строкой и столбцом нулей
	prefixSum := make([][]int, rows+1)
	for i := range prefixSum {
		prefixSum[i] = make([]int, cols+1)
	}

	// Заполняем матрицу префиксных сумм
	for i := 1; i <= rows; i++ { // Перебираем все строки в матрице
		for j := 1; j <= cols; j++ { // Перебираем все ячейки в матрице
			// Вычисляем префиксную сумму для текущей ячейки
			prefixSum[i][j] = matrix[i-1][j-1] + prefixSum[i-1][j] + prefixSum[i][j-1] - prefixSum[i-1][j-1]
		}
	}

	// Возвращаем объект NumMatrix с матрицей префиксных сумм
	return NumMatrix{prefixSum: prefixSum}
}

// SumRegion - сумма прямоугольника.
// time: O(1)
// memory: O(1)
func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if this.prefixSum == nil {
		return 0
	}
	// Возвращаем сумму элементов матрицы внутри прямоугольника, определенного верхним левым углом (row1, col1)
	return this.prefixSum[row2+1][col2+1] - this.prefixSum[row1][col2+1] - this.prefixSum[row2+1][col1] + this.prefixSum[row1][col1]
}
