package main

import (
	"fmt"
	"sort"
)

/* 18. 4Sum
https://leetcode.com/problems/4sum/description/

Дан массив nums из n целых чисел, вернуть массив всех уникальных четверок [nums[a], nums[b], nums[c], nums[d]] таких, что:

0 <= a, b, c, d < n
a, b, c и d различны.
nums[a] + nums[b] + nums[c] + nums[d] == target
Вы можете вернуть ответ в любом порядке.

Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
*/

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println(fourSum(nums, target)) // [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
}

// fourSum - возвращает массив всех уникальных четверок [nums[a], nums[b], nums[c], nums[d]] таких, что nums[a] + nums[b] + nums[c] + nums[d] == target.
// time: O(n^3), space: O(n)
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)    // Сортируем массив для удобства работы с двумя указателями [-2 -1 0 0 1 2]
	var result [][]int // Результат
	n := len(nums)     // Длина массива

	for i := 0; i < n-3; i++ { // Первый указатель
		// Пропускаем дубликаты для первого числа
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n-2; j++ { // Второй указатель
			// Пропускаем дубликаты для второго числа
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left := j + 1  // Левый указатель
			right := n - 1 // Правый указатель

			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right] // Сумма

				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					// Пропускаем дубликаты для третьего числа
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					// Пропускаем дубликаты для четвертого числа
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
					// Если сумма меньше целевого значения, двигаем левый указатель вправо
				} else if sum < target {
					left++
					// Если сумма больше целевого значения, двигаем правый указатель влево
				} else {
					right--
				}
			}
		}
	}

	return result // Возвращаем массив всех уникальных четверок
}
