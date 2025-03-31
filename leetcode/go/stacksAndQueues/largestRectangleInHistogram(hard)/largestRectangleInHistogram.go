package main

import "fmt"

/* https://leetcode.com/problems/largest-rectangle-in-histogram/description/

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

func largestRectangleArea(heights []int) int {
	stack := make([]int, 0)
	maxArea := 0
	index := 0

	for index < len(heights) {
		if len(stack) == 0 || heights[index] >= heights[stack[len(stack)-1]] {
			stack = append(stack, index)
			index++
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			height := heights[top]
			width := index
			if len(stack) > 0 {
				width = index - stack[len(stack)-1] - 1
			}
			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		height := heights[top]
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
