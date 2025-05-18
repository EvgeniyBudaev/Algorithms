package main

import (
	"fmt"
	"sort"
)

/* 128. Longest Consecutive Sequence
https://leetcode.com/problems/longest-consecutive-sequence/description/

Учитывая несортированный массив целых чисел nums, верните длину самой длинной последовательной последовательности
элементов.
Вы должны написать алгоритм, который работает за время O(n).

Input: nums = [100,4,200,1,3,2]
Output: 4
Объяснение: Самая длинная последовательная последовательность элементов — [1,2,3,4]. Следовательно, его длина равна 4

Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9
*/

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums)) // 4
}

// longestConsecutive находит длину самой длинной последовательности последовательных чисел.
// time: O(n), space: O(1)
func longestConsecutive(nums []int) int {
	// Если массив пуст, возвращаем 0.
	if len(nums) == 0 {
		return 0
	}

	// Сортируем массив по возрастанию
	sort.Ints(nums) // time: O(n log n), space: O(log n) для сортировки
	// Ищем самую длинную последовательность
	return search(nums)
}

// search проходит по отсортированному массиву и находит самую длинную последовательность.
// time: O(n), space: O(1)
func search(nums []int) int {
	maxScore := 1  // Максимальная длина последовательности
	currScore := 1 // Текущая длина последовательности

	for i := 1; i < len(nums); i++ {
		// Пропускаем дубликаты
		if nums[i-1] == nums[i] {
			continue
		}

		// Проверяем, является ли текущее число следующим в последовательности
		if nums[i] == nums[i-1]+1 {
			currScore++ // Увеличиваем текущую длину последовательности
			continue
		}

		// Обновляем максимальную длину если текущая последовательность закончилась
		if currScore > maxScore {
			maxScore = currScore
		}
		currScore = 1 // Сбрасываем счетчик для новой последовательности
	}

	// Возвращаем максимальную длину (на случай если самая длинная последовательность в конце массива)
	if currScore > maxScore {
		return currScore
	}
	return maxScore
}
