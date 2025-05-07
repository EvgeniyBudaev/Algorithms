package main

import (
	"fmt"
)

/* 189. Rotate Array
https://leetcode.com/problems/rotate-array/description/

Дан целочисленный массив nums. Повернуть массив вправо на k шагов, где k — неотрицательное число.

Example 1:
Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]

Example 2:
Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]
*/

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	k1 := 3
	rotate(nums1, k1)
	fmt.Println(nums1) // [5,6,7,1,2,3,4]

	nums2 := []int{-1, -100, 3, 99}
	k2 := 2
	rotate(nums2, k2)
	fmt.Println(nums2) // [3,99,-1,-100]
}

// rotate — поворачивает массив от индекса 0 до индекса len(nums)-1.
// time complexity: O(n) — проходим по всем элементам.
// space complexity: O(1) — не используем дополнительную память.
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n // Нормализуем k, чтобы он был в пределах длины массива
	if k == 0 {
		return // Если k = 0, вращение не нужно
	}
	rotateSubArray(nums, 0, n) // Поворачиваем весь массив. // [7, 6, 5, 4, 3, 2, 1]
	rotateSubArray(nums, 0, k) // [5, 6, 7, 4, 3, 2, 1]
	rotateSubArray(nums, k, n) // [5, 6, 7, 1, 2, 3, 4]
}

// rotateSubArray — поворачивает массив от индекса i до индекса j (не включительно).
// time complexity: O(n) — проходим по всем элементам.
// space complexity: O(1) — не используем дополнительную память.
func rotateSubArray(nums []int, i, j int) []int {
	j -= 1 // ЧТобы не включать последний индекс в поворот.
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}
