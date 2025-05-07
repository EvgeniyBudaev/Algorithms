package main

import (
	"fmt"
	"sort"
)

/* 252. Meeting Rooms

Учитывая массив временных интервалов встреч, где intervals[i] = [starti, endi], определите, может ли человек
присутствовать на всех собраниях.

Input: intervals = [[0,30],[5,10],[15,20]]
Output: false

Input: intervals = [[7,10],[2,4]]
Output: true
*/

func main() {
	intervals := [][]int{{7, 10}, {2, 4}}
	fmt.Println(canAttendMeetings(intervals)) // true
}

// canAttendMeetings - функция для определения, может ли человек присутствовать на всех собраниях.
// time complexity: O(n*log(n)) - из-за сортировки
// space complexity: O(1)
func canAttendMeetings(intervals [][]int) bool {
	// Сортируем интервалы по начальному времени
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	}) // [[2,4],[7,10]]

	// Проверяем, есть ли пересечения интервалов
	for i := 1; i < len(intervals); i++ {
		// Если конец предыдущего интервала больше начала текущего интервала, значит они пересекаются
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}

	return true
}
