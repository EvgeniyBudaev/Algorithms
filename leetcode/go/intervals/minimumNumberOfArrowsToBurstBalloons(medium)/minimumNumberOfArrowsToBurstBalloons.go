package main

import (
	"fmt"
	"sort"
)

/* 452. Minimum Number of Arrows to Burst Balloons
https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/description/

На плоской стене, представляющей плоскость XY, приклеены несколько сферических воздушных шаров.
Воздушные шары представлены в виде двумерного целочисленного массива точек, где points[i] = [xstart, xend]
обозначает воздушный шар, горизонтальный диаметр которого простирается между xstart и xend. Вы не знаете точных
координат y воздушных шаров.

Стрелы можно выпускать вертикально вверх (в положительном направлении y) из разных точек вдоль оси x.
Воздушный шар с xstart и xend лопается стрелой, выпущенной в x, если xstart <= x <= xend.
Нет ограничений на количество выпущенных стрел. Выпущенная стрела продолжает двигаться вверх бесконечно, лопая все
воздушные шары на своем пути.

Учитывая массив точек, верните минимальное количество стрел, которое необходимо выпустить, чтобы лопнуть все воздушные шары.

Example 1:
Input: points = [[10,16],[2,8],[1,6],[7,12]]
Output: 2
Объяснение: Воздушные шары можно лопнуть двумя стрелами:
- Выстрелите стрелой в точку x = 6, лопнув шары [2,8] и [1,6].
- Выстрелите стрелой в точку x = 11, лопнув шары [10,16] и [7,12].
*/

func main() {
	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	fmt.Println(findMinArrowShots(points)) // 2
}

// findMinArrowShots - функция, которая находит минимальное количество стрел, которое необходимо выпустить, чтобы лопнуть все воздушные шары.
// time complexity: O(n * log n)
// space complexity: O(1)
func findMinArrowShots(intervals [][]int) int {
	// Если нет воздушных шаров, то нет стрел
	if len(intervals) == 0 {
		return 0
	}

	// Сортируем интервалы по начальной точке
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// [[1, 6] [2, 8] [7, 12] [10, 16]]

	result := 1                  // Первый интервал всегда лопается стрелой
	lastInterval := intervals[0] // Последний интервал, который лопается стрелой

	for _, interval := range intervals {
		// Если интервал не пересекается с последним лопающим интервалом, то увеличиваем количество стрел
		if isOverlapping(lastInterval, interval) {
			lastInterval = overlapTwoIntervals(lastInterval, interval) // Иначе, обновляем последний лопающий интервал
			continue
		}
		lastInterval = interval // Иначе, обновляем последний лопающий интервал
		result++                // И увеличиваем количество стрел
	}

	return result
}

// isOverlapping - функция, которая проверяет, пересекаются ли два интервала.
func isOverlapping(a, b []int) bool {
	return max(a[0], b[0]) <= min(a[1], b[1])
}

// overlapTwoIntervals - функция, которая объединяет два интервала.
func overlapTwoIntervals(a, b []int) []int {
	return []int{max(a[0], b[0]), min(a[1], b[1])}
}
