package main

import (
	"fmt"
	"strconv"
)

/* 228. Summary Ranges
https://leetcode.com/problems/summary-ranges/description/

Вам дан отсортированный уникальный целочисленный массив чисел.
Диапазон [a,b] — это набор всех целых чисел от a до b (включительно).

Верните наименьший отсортированный список диапазонов, который точно охватывает все числа в массиве. То есть каждый
элемент nums охватывается ровно одним из диапазонов, и не существует целого числа x такого, что x находится в одном из
диапазонов, но не в nums.

Каждый диапазон [a,b] в списке должен выводиться как:
"a->b" if a != b
"a" if a == b

Input: nums = [0,1,2,4,5,7]
Output: ["0->2","4->5","7"]
Пояснение: Диапазоны:
[0,2] --> "0->2"
[4,5] --> "4->5"
[7,7] --> "7"
*/

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(summaryRanges(nums)) // [0->2 4->5 7]
}

// summaryRanges - возвращает наименьший отсортированный список диапазонов, который точно охватывает все числа в массиве nums.
// time: O(n), space: O(n)
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	var result []string // Результирующий список диапазонов
	start := nums[0]    // Начало текущего диапазона

	for i := 1; i <= len(nums); i++ {
		// Если числа идут подряд
		if i < len(nums) && nums[i]-nums[i-1] == 1 {
			continue // Продолжаем, если числа идут подряд
		}

		// Если диапазон состоит из одного числа
		if start == nums[i-1] {
			result = append(result, strconv.Itoa(start))
		} else {
			// Если диапазон состоит из нескольких чисел
			result = append(result, fmt.Sprintf("%d->%d", start, nums[i-1]))
		}

		// Обновляем начало диапазона
		if i < len(nums) {
			start = nums[i]
		}
	}

	return result // Возвращает наименьший отсортированный список диапазонов
}
