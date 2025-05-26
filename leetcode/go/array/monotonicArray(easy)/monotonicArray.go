package main

import (
	"fmt"
)

/* 896. Monotonic Array
https://leetcode.com/problems/monotonic-array/description/

Массив монотонен, если он либо монотонно возрастает, либо монотонно убывает.
Массив nums монотонно возрастает, если для всех i <= j, nums[i] <= nums[j]. Массив nums монотонно убывает,
если для всех i <= j, nums[i] >= nums[j].
Если задан целочисленный массив nums, вернуть true, если заданный массив монотонен, или false в противном случае.

Input: nums = [1,2,2,3]
Output: true

Input: nums = [6,5,4,4]
Output: true

Input: nums = [1,3,2]
Output: false
*/

func main() {
	nums1 := []int{1, 2, 2, 3}
	fmt.Println(isMonotonic(nums1)) // true

	nums2 := []int{6, 5, 4, 4}
	fmt.Println(isMonotonic(nums2)) // true

	nums3 := []int{1, 3, 2}
	fmt.Println(isMonotonic(nums3)) // false
}

// isMonotonic возвращает true, если заданный массив монотонен, или false в противном случае.
// time: O(n), space: O(1)
func isMonotonic(nums []int) bool {
	isIncreasing := true // монотонно возрастает
	isDecreasing := true // монотонно убывает

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] { // если монотонно возрастает
			isDecreasing = false
		} else if nums[i] < nums[i-1] { // если монотонно убывает
			isIncreasing = false
		}
		// если не монотонно возрастает и не монотонно убывает
		if !isIncreasing && !isDecreasing {
			return false
		}
	}

	return true // если монотонно возрастает или монотонно убывает
}
