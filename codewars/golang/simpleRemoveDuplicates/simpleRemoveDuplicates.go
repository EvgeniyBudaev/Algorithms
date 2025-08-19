package main

import (
	"fmt"
)

/* Simple remove duplicates
https://www.codewars.com/kata/5ba38ba180824a86850000f7/train/go

Удалить дубликаты из списка целых чисел, сохранив последнее (самое правое) вхождение каждого элемента.

Пример:
Для входных данных: [3, 4, 4, 3, 6, 3]

удалить 3 элемента с индексом 0
удалить 4 элемента с индексом 1
удалить 3 элемента с индексом 3
Ожидаемый результат: [4, 6, 3]
*/

func main() {
	arr := []int{3, 4, 4, 3, 6, 3}
	fmt.Println(Solve(arr)) // [4, 6, 3]
}

// Solve возвращает список целых чисел без дубликатов.\
// time: O(n), space: O(n), где n - длина массива
func Solve(arr []int) []int {
	lastIndex := make(map[int]int)

	// Запоминаем последние индексы каждого элемента
	for i, num := range arr {
		lastIndex[num] = i
	}

	result := make([]int, 0)

	// Добавляем элементы, которые встречаются на своих последних позициях
	for i, num := range arr {
		if lastIndex[num] == i {
			result = append(result, num)
		}
	}

	return result
}
