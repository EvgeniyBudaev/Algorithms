package main

import (
	"fmt"
)

/* 907. Sum of Subarray Minimums
https://leetcode.com/problems/sum-of-subarray-minimums/description/

Дан массив целых чисел arr, найти сумму min(b), где b пробегает каждый (смежный) подмассив arr.
Поскольку ответ может быть большим, верните ответ по модулю 109 + 7.

Input: arr = [3,1,2,4]
Output: 17
Пояснение:
Подмассивы — это [3], [1], [2], [4], [3,1], [1,2], [2,4], [3,1,2], [1,2,4], [3,1,2,4].
Минимумы — 3, 1, 2, 4, 1, 1, 2, 1, 1, 1.
Сумма — 17.
*/

func main() {
	arr := []int{3, 1, 2, 4}
	fmt.Println(sumSubarrayMins(arr)) // 17
}

// sumSubarrayMins находит сумму минимумов каждого подмассива.
// time: O(n), space: O(n)
func sumSubarrayMins(arr []int) int {
	sums := make([]int, len(arr)) // Суммы каждого подмассива
	var stack []int               // Стэк индексов

	for i, v := range arr {
		for len(stack) > 0 && v < arr[stack[len(stack)-1]] { // Сравниваем текущее число с верхним индексом стэка
			stack = stack[:len(stack)-1] // Если текущее число меньше верхнего индекса стэка, удаляем верхний индекс стэка
		}
		stack = append(stack, i) // Добавляем текущий индекс стэку

		if len(stack) == 1 { // Если стэк пустой, то сумма подмассива равна текущему числу
			sums[i] += (i + 1) * v // Сумма подмассива равна текущему числу
		} else { // Если стэк не пустой, то сумма подмассива равна сумме подмассива с верхним индексом стэка и текущему числу
			prev := stack[len(stack)-2]        // Предыдущий индекс стэка
			sums[i] += (i-prev)*v + sums[prev] // Сумма подмассива равна сумме подмассива с верхним индексом стэка и текущему числу
		}
	}

	var res int              // Результат
	for _, v := range sums { // Суммируем все суммы подмассивов
		res += v // Суммируем все суммы подмассивов
	}

	// 1e9 + 7 Модуль для избежания переполнения
	// Результат по модулю 109 + 7
	return res % (1e9 + 7)
}
