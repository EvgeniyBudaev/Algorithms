package main

import "fmt"

/* 118. Pascal's Triangle
https://leetcode.com/problems/pascals-triangle/description/

Дано целое число numRows, вернуть первые numRows треугольника Паскаля.
В треугольнике Паскаля каждое число является суммой двух чисел, расположенных непосредственно над ним, как показано:

Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
*/

func main() {
	numRows := 5
	fmt.Println(generate(numRows)) // [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
}

// generate возвращает первые numRows треугольника Паскаля.
// time: O(numRows^2) - количество итераций в цикле в цикле
// space: O(numRows^2) - количество итераций в цикле в цикле
func generate(numRows int) [][]int {
	var triangle [][]int // Треугольник Паскаля
	if numRows == 0 {    // Если numRows равно 0, то возвращаем пустой треугольник Паскаля
		return triangle
	}

	triangle = append(triangle, []int{1}) // Первый ряд треугольника Паскаля всегда состоит из одного числа 1

	for i := 1; i < numRows; i++ {
		prevRow := triangle[i-1]   // Предыдущий ряд треугольника Паскаля
		var newRow []int           // Текущий ряд треугольника Паскаля
		newRow = append(newRow, 1) // Первый элемент текущего ряда всегда равен 1

		for j := 1; j < len(prevRow); j++ { // Для каждого элемента в предыдущем ряду
			newRow = append(newRow, prevRow[j-1]+prevRow[j]) // Сумма двух элементов, расположенных непосредственно над ним
		}

		newRow = append(newRow, 1)          // Последний элемент текущего ряда всегда равен 1
		triangle = append(triangle, newRow) // Добавляем текущий ряд в треугольник Паскаля
	}

	// Возвращаем треугольник Паскаля
	return triangle
}
