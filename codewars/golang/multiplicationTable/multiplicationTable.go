package main

import (
	"fmt"
)

/* Multiplication table
https://www.codewars.com/kata/534d2f5b5371ecf8d2000a08/train/go

Ваша задача — создать таблицу умножения N×N, размер которой указан в параметре.

Например, если заданный размер равен 3:
1 2 3
2 4 6
3 6 9
*/

func main() {
	fmt.Println(MultiplicationTable(3)) // [[1,2,3],[2,4,6],[3,6,9]]
}

// MultiplicationTable создает таблицу умножения.
// time: O(n^2), space: O(n^2), n - size
func MultiplicationTable(size int) [][]int {
	table := make([][]int, size) // // Создаем двумерный слайс нужного размера
	// [[] [] []]
	for i := range table { // Заполняем слайс нулями
		table[i] = make([]int, size)
	}
	// [[0, 0, 0] [0, 0, 0] [0, 0, 0]]

	// Заполняем таблицу умножения
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			table[i][j] = (i + 1) * (j + 1)
		}
	}

	return table // Возвращаем таблицу
}
