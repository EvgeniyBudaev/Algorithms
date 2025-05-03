package main

import "fmt"

/* 240. Search a 2D Matrix II
https://leetcode.com/problems/search-a-2d-matrix-ii/description/

Напишите эффективный алгоритм, который ищет целевое значение в целочисленной матрице размером m x n. Эта матрица имеет
следующие свойства:
Целые числа в каждой строке сортируются по возрастанию слева направо.
Целые числа в каждом столбце сортируются по возрастанию сверху вниз.

Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
Output: true

Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
Output: false
*/

func main() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	fmt.Println(searchMatrix(matrix, 5)) // true
}

// searchMatrix возвращает true, если цель находится в матрице, иначе false.
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	for _, row := range matrix {
		// Используем бинарный поиск для каждой строки
		left, right := 0, len(row)-1
		for left <= right {
			mid := (left + right) / 2
			// Если текущий элемент равен целевому значению, возвращаем true
			if row[mid] == target {
				return true
				// Если текущий элемент меньше целевого значения, продолжаем поиск в правой половине строки
			} else if row[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}
