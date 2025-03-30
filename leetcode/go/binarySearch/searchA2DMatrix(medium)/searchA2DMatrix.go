package main

import "fmt"

/* https://leetcode.com/problems/search-a-2d-matrix/description/

Вам дана целочисленная матрица размера m x n со следующими двумя свойствами:
Каждая строка сортируется в порядке неубывания.
Первое целое число каждой строки больше последнего целого числа предыдущей строки.
Учитывая целочисленную цель, верните true, если цель находится в матрице, или false в противном случае.
Вы должны написать решение с временной сложностью O(log(m * n)).

Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true

Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false
*/

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	fmt.Println(searchMatrix(matrix, 3)) // true
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows, cols := len(matrix), len(matrix[0])
	left, right := 0, rows*cols-1

	for left <= right {
		mid := (left + right) / 2
		row, col := mid/cols, mid%cols
		value := matrix[row][col]

		if value == target {
			return true
		} else if value < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
