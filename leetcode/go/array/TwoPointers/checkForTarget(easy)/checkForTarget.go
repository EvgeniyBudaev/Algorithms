package main

import "fmt"

/*
Учитывая отсортированный массив уникальных целых чисел и целевое целое число, верните true, если существует пара чисел,
сумма которых равна целевому значению, и false в противном случае. Эта задача аналогична задаче «Две суммы».
(В режиме Two Sum входные данные не сортируются).

Например, если заданы числа = [1, 2, 4, 6, 8, 9, 14, 15] и цель = 13, верните true, потому что 4 + 9 = 13.
*/

func main() {
	fmt.Println(checkForTarget([]int{1, 2, 4, 6, 8, 9, 14, 15}, 13)) // true
}

// checkForTarget проверяет, существует ли в отсортированном массиве nums пара чисел, сумма которых равна заданному целевому значению target.
func checkForTarget(nums []int, target int) bool {
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return true
		} else if sum > target {
			right--
		} else {
			left++
		}
	}

	return false
}
