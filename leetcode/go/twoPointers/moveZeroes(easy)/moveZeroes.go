package main

import "fmt"

/* 283. Move Zeroes
https://leetcode.com/problems/move-zeroes/description/

Учитывая целочисленный массив nums, переместите все 0 в его конец, сохраняя относительный порядок ненулевых элементов.
Обратите внимание, что вы должны сделать это на месте, не копируя массив.

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

Input: nums = [0]
Output: [0]
*/

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums) // [1, 3, 12, 0, 0]
	fmt.Println(nums)
}

// moveZeroes -  перемещает все нули в конец массива nums, сохраняя относительный порядок ненулевых элементов.
// time: O(n), space: O(1)
func moveZeroes(nums []int) {
	left := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			nums[right], nums[left] = nums[left], nums[right]
			left++
		}
	}
}
