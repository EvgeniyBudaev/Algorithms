package main

import (
	"fmt"
	"strconv"
	"strings"
)

/* CSV representation of array
https://www.codewars.com/kata/5a34af40e1ce0eb1f5000036/train/go

Создайте функцию, которая возвращает CSV-представление двумерного числового массива.

Example:

input:
   [[ 0, 1, 2, 3, 4 ],
    [ 10,11,12,13,14 ],
    [ 20,21,22,23,24 ],
    [ 30,31,32,33,34 ]]

output:
     '0,1,2,3,4\n'
    +'10,11,12,13,14\n'
    +'20,21,22,23,24\n'
    +'30,31,32,33,34'
*/

func main() {
	input := [][]int{{0, 1, 2, 3, 4}, {10, 11, 12, 13, 14}, {20, 21, 22, 23, 24}, {30, 31, 32, 33, 34}}
	fmt.Println(ToCsvText(input)) // '0,1,2,3,4\n' +'10,11,12,13,14\n' +'20,21,22,23,24\n' +'30,31,32,33,34'
}

// ToCsvText возвращает CSV-представление двумерного числового массива.
// time: O(n), space: O(n), где n - количество элементов в массиве
func ToCsvText(array [][]int) string {
	var result []string

	for _, row := range array {
		// Преобразуем числа в строки
		strRow := make([]string, len(row))
		for i, num := range row {
			strRow[i] = strconv.Itoa(num)
		}
		// Объединяем через запятую
		result = append(result, strings.Join(strRow, ","))
	}

	// Добавляем перенос строки между строками
	return strings.Join(result, "\n")
}
