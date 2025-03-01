package main

import "fmt"

/* https://leetcode.com/problems/max-consecutive-ones/description/

Для двоичного массива nums вернуть максимальное количество последовательных 1 в массиве.

Example 1
Input: nums = [1,1,0,1,1,1]
Output: 3

Example 2
Input: nums = [1,0,1,1,0,1]
Output: 2
*/

func main() {
	nums := []int{1, 1, 0, 1, 1, 1}
	r := findMaxConsecutiveOnes(nums)
	fmt.Println(r) // 3
}

func findMaxConsecutiveOnes(nums []int) int {
	left, zeroCount, ans := 0, 0, 0

	for right := 0; right < len(nums); right++ {
		// Если ноль, увеличиваем счетчик подсчета нулей
		if nums[right] == 0 {
			zeroCount++
		}

		// Если количество нулей больше 0, двигаем левый указатель
		for zeroCount > 0 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		// Обновляем максимальную длину последовательных единиц
		currentWindowSize := right - left + 1

		if currentWindowSize > ans {
			ans = currentWindowSize
		}
	}

	return ans
}
