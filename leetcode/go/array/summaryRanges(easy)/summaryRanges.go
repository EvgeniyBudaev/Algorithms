package main

import (
	"fmt"
	"strconv"
)

/* https://leetcode.com/problems/summary-ranges/description/

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
	r := summaryRanges(nums)
	fmt.Println(r)
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	result := []string{}
	start := nums[0]

	for i := 1; i <= len(nums); i++ {
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

	return result
}
