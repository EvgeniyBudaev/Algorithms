package main

import (
	"fmt"
	"math"
)

/* 356. Line Reflection

Учитывая n точек на двумерной плоскости, найдите, существует ли такая линия, параллельная оси Y, которая отражает данные
точки.

Input: [[1,1],[-1,1]]
Output: true

Input: [[1,1],[-1,-1]]
Output: false
*/

func main() {
	points1 := [][]int{{1, 1}, {-1, 1}}
	fmt.Println(isReflected(points1)) // true
}

// isReflected - проверяет, существует ли такая линия, параллельная оси Y, которая отражает данные точки.
// time: O(n), space: O(n)
func isReflected(points [][]int) bool {
	if len(points) == 0 { // Если нет точек, то линия отражения не существует
		return true
	}

	// Находим минимальный и максимальный X
	minX, maxX := math.MaxInt32, math.MinInt32
	for _, point := range points {
		x := point[0] // x-координата
		if x < minX { // Если x меньше минимального, то обновляем минимальное
			minX = x // Обновляем минимальное
		}
		if x > maxX { // Если x больше максимального, то обновляем максимальное
			maxX = x // Обновляем максимальное
		}
	}

	// Создаем множество точек для быстрой проверки
	pointSet := make(map[[2]int]bool)
	for _, point := range points { // Добавляем каждую точку в множество
		pointSet[[2]int{point[0], point[1]}] = true // Добавляем точку в множество
	}

	// Проверяем, есть ли отраженная точка для каждой исходной
	for point := range pointSet {
		x, y := point[0], point[1]            // x-координата и y-координата
		reflectedX := maxX + minX - x         // Отраженная x-координата
		if !pointSet[[2]int{reflectedX, y}] { // Если отраженная точка не существует, то линия отражения не существует
			return false
		}
	}

	// Если все точки отражены, то линия отражения существует
	return true
}
