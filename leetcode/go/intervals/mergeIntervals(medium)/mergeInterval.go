package main

import (
	"fmt"
	"sort"
)

/* 56. Merge Intervals
https://leetcode.com/problems/merge-intervals/description/

Дан массив интервалов, где intervals[i] = [starti, endi], объединить все перекрывающиеся интервалы и вернуть массив
неперекрывающихся интервалов, которые покрывают все интервалы во входных данных.

Example 1:
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Пояснение: Поскольку интервалы [1,3] и [2,6] перекрываются, объединим их в [1,6].
*/

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals)) // [[1,6],[8,10],[15,18]]
}

// merge принимает массив интервалов и объединяет все перекрывающиеся интервалы.
// time complexity: O(n*log(n))
// space complexity: O(1)
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// Сортируем интервалы по начальному значению. O(n*log(n))
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]} // Результат, начинаем с первого интервала

	// O(n)
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]       // Текущий интервал
		last := result[len(result)-1] // Последний интервал в результате

		// если текущий интервал и последний в ответе пересекаются,
		// значит объединяем их, иначе добавляем интервал к ответу и это значит,
		// что ни один интервал, который имеет точку начала меньше текущего интервала
		// не будет пересечен ни с одним лежащим правее и не с текущим
		if isOverlapping(last, current) {
			result[len(result)-1] = mergeTwoIntervals(last, current)
		} else {
			result = append(result, current)
		}
	}

	return result
}

// isOverlapping проверяет, пересекаются ли два интервала.
func isOverlapping(a, b []int) bool {
	return max(a[0], b[0]) <= min(a[1], b[1])
}

// mergeTwoIntervals объединяет два интервала.
func mergeTwoIntervals(a, b []int) []int {
	// Интервалы обязательно должны пересекаться
	return []int{a[0], max(a[1], b[1])}
}
