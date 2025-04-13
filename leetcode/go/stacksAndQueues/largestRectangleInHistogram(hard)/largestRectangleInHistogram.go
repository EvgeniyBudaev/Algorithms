package main

import "fmt"

/* 84. Largest Rectangle in Histogram
https://leetcode.com/problems/largest-rectangle-in-histogram/description/

Учитывая массив целых чисел, представляющих высоту столбца гистограммы, где ширина каждого столбца равна 1, верните
площадь самого большого прямоугольника в гистограмме.

Input: heights = [2,1,5,6,2,3]
Output: 10
Пояснение: Выше представлена гистограмма, ширина каждого столбца которой равна 1.
Самый большой прямоугольник показан в красной области, его площадь = 10 единиц.
*/

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(heights)) // 10
}

// largestRectangleArea вычисляет площадь самого большого прямоугольника, который можно построить на гистограмме.
func largestRectangleArea(heights []int) int {
	stack := make([]int, 0) // Стек для хранения индексов столбцов гистограммы
	maxArea := 0            // Максимальная площадь прямоугольника
	index := 0              // Текущий индекс в массиве heights

	// Первый проход по всем столбцам гистограммы
	for index < len(heights) {
		// Если стек пуст или текущий столбец выше или равен последнему в стеке
		if len(stack) == 0 || heights[index] >= heights[stack[len(stack)-1]] {
			stack = append(stack, index) // Добавляем индекс текущего столбца в стек
			index++
		} else {
			// Если текущий столбец ниже, извлекаем последний индекс из стека
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// Вычисляем высоту прямоугольника - высота извлеченного столбца
			height := heights[top]
			// Ширина прямоугольника:
			// - если стек пуст, ширина равна текущему индексу (все предыдущие столбцы были выше)
			// - иначе ширина = расстояние между текущим индексом и предыдущим в стеке минус 1
			width := index
			if len(stack) > 0 {
				width = index - stack[len(stack)-1] - 1
			}
			// Вычисляем площадь и обновляем максимум при необходимости
			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}
	}

	// Второй проход - обрабатываем оставшиеся элементы в стеке
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		height := heights[top]
		// Для оставшихся элементов ширина равна длине всей гистограммы минус позиция предыдущего элемента
		width := index
		if len(stack) > 0 {
			width = index - stack[len(stack)-1] - 1
		}
		area := height * width
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}
