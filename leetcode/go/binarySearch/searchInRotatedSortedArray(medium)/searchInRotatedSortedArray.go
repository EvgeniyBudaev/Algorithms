package main

import "fmt"

/* 33. Search in Rotated Sorted Array
https://leetcode.com/problems/search-in-rotated-sorted-array/description/

Существует целочисленный массив nums, отсортированный по возрастанию (с разными значениями).

Перед передачей в вашу функцию nums, возможно, поворачивается с неизвестным индексом поворота k (1 <= k < nums.length),
так что результирующий массив имеет вид [nums[k], nums[k+1],... , nums[n-1], nums[0], nums[1], ..., nums[k-1]]
(с индексом 0). Например, [0,1,2,4,5,6,7] можно повернуть с опорным индексом 3 и превратить в [4,5,6,7,0,1,2].

Учитывая массив nums после возможного поворота и целочисленную цель, верните индекс цели, если он находится в nums, или
-1, если он не в nums.

Вы должны написать алгоритм с компиляцией времени выполнения O(log n).

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
*/

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(nums, 0)) // 4
}

// search ищет индекс цели в отсортированном массиве после возможного поворота.
// time: O(log n), space: O(1)
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2
		guess := nums[mid]
		leftNum, rightNum := nums[low], nums[high]

		if guess == target {
			return mid
		}

		// Если левая половина отсортирована
		if leftNum <= guess {
			// Если цель находится между левой и центральной числами
			if leftNum <= target && target < guess {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			// Правая половина отсортирована
			// Если цель находится между правой и центральной числами
			if guess < target && target <= rightNum {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}
