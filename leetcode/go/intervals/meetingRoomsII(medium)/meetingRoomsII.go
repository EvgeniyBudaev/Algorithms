package main

import (
	"fmt"
	"sort"
)

/* Meeting Rooms II

Дан двумерный целочисленный массив A размером N x 2, обозначающий временные интервалы различных встреч.

Где:
A[i][0] = время начала i-й встречи.
A[i][1] = время окончания i-й встречи.

Найдите минимальное количество конференц-залов, необходимое для проведения всех встреч.

Примечание: если встреча заканчивается в момент времени t, другая встреча, начинающаяся в момент времени t,
может использовать тот же конференц-зал.

Example 1:
Input: intervals = [(0,30),(5,10),(15,20)]
Output: 2
Explanation:
room1: (0,30)
room2: (5,10),(15,20)
*/

func main() {
	intervals := [][]int{{0, 30}, {5, 10}, {15, 20}}
	fmt.Println(minMeetingRooms(intervals)) // 2
}

// minMeetingRooms - находит минимальное количество конференц-залов, необходимое для проведения всех встреч.
// time complexity: O(n * log n)
// space complexity: O(n)
func minMeetingRooms(intervals [][]int) int {
	// point - точка начала или конца встречи
	type point struct {
		time  int // время начала или конца встречи
		delta int // дельта количества комнат: +1 если начало встречи, -1 если конец встречи
	}

	points := make([]point, 0, len(intervals)*2) // объединяем начало и конец встречи в одну точку

	// Создаем точки начала и конца событий
	for _, interval := range intervals {
		points = append(points, point{interval[0], 1})  // начало встречи
		points = append(points, point{interval[1], -1}) // конец встречи
	}
	// points [{0, 1} {30, -1} {5, 1} {10, -1} {15, 1} {20, -1}]

	// Сортируем точки:
	// 1. По времени
	// 2. Если время одинаковое, сначала идут окончания встреч (delta = -1)
	sort.Slice(points, func(i, j int) bool {
		// если время совпадает, то сначала идут окончания встреч (delta = -1)
		if points[i].time == points[j].time {
			return points[i].delta < points[j].delta
		}
		// если время не совпадает, то сначала идут началы встреч (delta = +1)
		return points[i].time < points[j].time
	})
	// points [{0, 1} {5, 1} {10, -1} {15, 1} {20, -1} {30, -1}]

	maxRooms := 0     // максимальное количество комнат
	currentRooms := 0 // текущее количество комнат

	for _, p := range points {
		currentRooms += p.delta // увеличиваем или уменьшаем количество комнат в зависимости от типа события
		// если количество комнат больше максимального, то обновляем максимальное количество комнат
		if currentRooms > maxRooms {
			maxRooms = currentRooms
		}
	}

	// возвращаем максимальное количество комнат
	return maxRooms
}
