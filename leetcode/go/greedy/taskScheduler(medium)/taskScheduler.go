package main

import (
	"fmt"
	"sort"
)

/* https://leetcode.com/problems/task-scheduler/description/

Вам дан массив задач ЦП, каждая из которых обозначена буквами от A до Z, и время охлаждения n. Каждый цикл или интервал
позволяет выполнить одну задачу. Задачи можно выполнять в любом порядке, но есть ограничение: одинаковые задачи должны
быть разделены как минимум n интервалами из-за времени остывания.
Верните минимальное количество интервалов, необходимое для выполнения всех задач.

Input: tasks = ["A","A","A","B","B","B"], n = 2
Output: 8

Пояснение: Возможная последовательность: A -> B -> ожидание -> A -> B -> ожидание -> A -> B.
После завершения задачи А вы должны подождать два цикла, прежде чем снова выполнить А. То же самое относится и к
задаче Б. В 3-м интервале ни А, ни Б сделать невозможно, поэтому вы бездельничаете. К 4-му циклу можно снова делать А,
так как прошло 2 интервала.
*/

func main() {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	fmt.Println(leastInterval(tasks, 2)) // 8
}

func leastInterval(tasks []byte, n int) int {
	// Создаем карту для подсчета частот задач
	taskMap := make(map[byte]int)
	for _, task := range tasks {
		taskMap[task]++
	}

	// Создаем слайс для хранения частот и сортируем по убыванию
	taskCounts := make([]int, 0, len(taskMap))
	for _, count := range taskMap {
		taskCounts = append(taskCounts, count)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(taskCounts)))

	// Вычисляем максимальное количество повторений
	maxCount := taskCounts[0]

	// Вычисляем количество "холостых" слотов
	idleSlots := (maxCount - 1) * n

	// Заполняем холостые слоты другими задачами
	for i := 1; i < len(taskCounts); i++ {
		idleSlots -= min(taskCounts[i], maxCount-1)
	}

	// Холостые слоты не могут быть отрицательными
	if idleSlots < 0 {
		idleSlots = 0
	}

	// Общее время = все задачи + холостые слоты
	return len(tasks) + idleSlots
}
