package main

import "fmt"

/* https://leetcode.com/problems/interval-list-intersections/description/
solution https://leetcode.com/problems/interval-list-intersections/solutions/4686775/simple-easy-solution-using-two-pointer/

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

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var result [][]int
	left, right := 0, 0

	for left < len(firstList) && right < len(secondList) {
		first := firstList[left]
		second := secondList[right]
		start := max(first[0], second[0])
		end := min(first[1], second[1])
		if start <= end {
			result = append(result, []int{start, end})
		}
		if first[1] < second[1] {
			left++
		} else if first[1] > second[1] {
			right++
		} else {
			left++
			right++
		}
	}

	return result
}
