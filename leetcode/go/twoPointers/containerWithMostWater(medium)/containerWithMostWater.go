package main

import (
	"fmt"
)

/* 11. Container With Most Water
https://leetcode.com/problems/container-with-most-water/description/

Вам дан целочисленный массив высотой длины n. Нарисовано n вертикальных линий, конечными точками которых являются
(i, 0) и (i, height[i]).
Найдите две линии, которые вместе с осью X образуют контейнер, в котором содержится больше всего воды.
Верните максимальное количество воды, которое может хранить контейнер.
Обратите внимание, что вы не можете наклонять контейнер

Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Пояснение: Вышеупомянутые вертикальные линии представлены массивом [1,8,6,2,5,4,8,3,7].
В этом случае максимальная площадь воды (синяя секция), которую может содержать контейнер, равна 49.

Input: height = [1,1]
Output: 1
*/

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(nums)) // 49
}

// maxArea вычисляет максимальную площадь контейнера, образованного двумя вертикальными линиями и осью X,
// где каждая линия задана высотой height[i], а расстояние между линиями определяет ширину контейнера.
// time: O(n), space: O(1)
func maxArea(height []int) int {
	maxAreaContainer := 0 // Максимальная площадь воды
	left := 0             // Левый указатель
	right := len(height) - 1

	for left < right {
		// Находим минимальную высоту между двумя линиями
		minHeight := min(height[left], height[right])
		// Вычисляем ширину
		width := right - left
		// Вычисляем площадь и обновляем maxArea, если текущая площадь больше
		currentArea := width * minHeight
		maxAreaContainer = max(maxAreaContainer, currentArea)
		// Перемещаем указатель с меньшей высотой
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxAreaContainer // Возвращаем максимальную площадь
}
