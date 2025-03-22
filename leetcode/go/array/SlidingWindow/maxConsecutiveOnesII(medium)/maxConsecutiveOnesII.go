package main

import "fmt"

/* 487. Max Consecutive Ones II
https://github.com/EvgeniyBudaev/doocs_leetcode/blob/main/solution/0400-0499/0487.Max%20Consecutive%20Ones%20II/README.md

Учитывая двоичный массив nums, если не более одного можно перевернуть 0, верните 1 максимальное количество
последовательных чисел в массиве.

Input: nums = [1,0,1,1,0]
Output: 4
Объяснение: переверните первый 0, чтобы получить самую длинную последовательную 1.
     После переворачивания максимальное количество последовательных единиц равно 4.

Input: nums = [1,0,1,1,0,1]
Output: 4
*/

func main() {
	nums := []int{1, 0, 1, 1, 0}
	fmt.Println(findMaxConsecutiveOnes(nums)) // 4
}

func findMaxConsecutiveOnes(nums []int) int {
	left, zeroCount, ans, maxFlipOperations := 0, 0, 0, 1

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}
		for zeroCount > maxFlipOperations {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}

	return ans
}
