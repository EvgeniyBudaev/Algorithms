package main

import "fmt"

/* 986. Interval List Intersections
https://leetcode.com/problems/interval-list-intersections/description/

Вам даны два списка закрытых интервалов, firstList и SecondList,
где firstList[i] = [starti, endi] и SecondList[j] = [startj, endj]. Каждый список интервалов попарно непересекающийся и
отсортирован.
Верните пересечение этих двух списков интервалов.
Замкнутый интервал [a, b] (с a <= b) обозначает набор действительных чисел x с a <= x <= b.
Пересечение двух закрытых интервалов представляет собой набор действительных чисел, которые либо пусты, либо
представлены в виде замкнутого интервала. Например, пересечение [1, 3] и [2, 4] — это [2, 3].

Input: firstList = [[0,2],[5,10],[13,23],[24,25]], secondList = [[1,5],[8,12],[15,24],[25,26]]
Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
*/

func main() {
	firstList := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	secondList := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	fmt.Println(intervalIntersection(firstList, secondList)) // [[1 2] [5 5] [8 10] [15 23] [24 24] [25 25]]
}

// intervalIntersection находит пересечение двух списков интервалов.
// time: O(n), space: O(n)
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var result [][]int // Результирующий список интервалов
	left, right := 0, 0

	for left < len(firstList) && right < len(secondList) { // Пока не закончатся интервалы в обоих списках
		first := firstList[left]          // Текущий интервал первого списка
		second := secondList[right]       // Текущий интервал второго списка
		start := max(first[0], second[0]) // Максимальное начало
		end := min(first[1], second[1])   // Минимальное окончание
		if start <= end {                 // Если начало меньше окончания, то добавляем интервал в результат
			result = append(result, []int{start, end})
		}
		// Усли окончание первого интервала меньше второго, то двигаем указатель первого интервала, иначе двигаем второй
		if first[1] < second[1] {
			left++
			// Усли окончание первого интервала больше второго, то двигаем указатель второго интервала
		} else if first[1] > second[1] {
			right++
		} else {
			left++
			right++
		}
	}

	return result // Список интервалов
}
