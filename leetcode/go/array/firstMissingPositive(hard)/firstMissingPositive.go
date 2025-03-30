package main

import "fmt"

/* https://leetcode.com/problems/first-missing-positive/description/

Дан несортированный целочисленный массив nums. Верните наименьшее положительное целое число, которого нет в nums.
Вы должны реализовать алгоритм, который работает за время O(n) и использует вспомогательное пространство O(1).

Input: nums = [1,2,0]
Output: 3
Объяснение: Все числа в диапазоне [1,2] входят в массив

Input: nums = [3,4,-1,1]
Output: 2
Пояснение: 1 есть в массиве, но 2 отсутствует.

Input: nums = [7,8,9,11,12]
Output: 1
Объяснение: Пропущено наименьшее положительное целое число 1.
*/

func main() {
	nums := []int{1, 2, 0}
	fmt.Println(firstMissingPositive(nums)) // 3
}

func firstMissingPositive(nums []int) int {
	set := make(map[int]bool)

	for _, num := range nums {
		if num > 0 {
			set[num] = true
		}
	}

	i := 1
	for set[i] {
		i++
	}

	return i
}
