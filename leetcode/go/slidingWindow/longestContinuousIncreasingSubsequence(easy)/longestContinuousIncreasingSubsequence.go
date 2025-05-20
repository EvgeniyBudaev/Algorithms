package main

import (
	"fmt"
)

/* 674. Longest Continuous Increasing Subsequence
https://leetcode.com/problems/longest-continuous-increasing-subsequence/description/

Дан несортированный массив целых чисел nums, вернуть длину самой длинной непрерывной возрастающей подпоследовательности
(т. е. подмассива). Подпоследовательность должна быть строго возрастающей.
Непрерывная возрастающая подпоследовательность определяется двумя индексами l и r (l < r),
такими, что она равна [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] и для каждого l <= i < r, nums[i] < nums[i + 1].

Input: nums = [1,3,5,4,7]
Output: 3
Объяснение: Самая длинная непрерывная возрастающая подпоследовательность — [1,3,5] с длиной 3.
Хотя [1,3,5,7] является возрастающей подпоследовательностью, она не является непрерывной, поскольку элементы 5 и 7
разделены элементом 4.
*/

func main() {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println(findLengthOfLCIS(nums)) // 3
}

// findLengthOfLCIS находит длину самой длинной непрерывной возрастающей подпоследовательности в массиве nums.
// time: O(n), space: O(1)
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 { // Если массив nums пуст, возвращаем 0
		return 0
	}

	left, maxLen := 0, 1 // left - левая граница текущей подпоследовательности, maxLen - максимальная длина найденной подпоследовательности

	for right := 1; right < len(nums); right++ {
		if nums[right] <= nums[right-1] { // Если текущий элемент не больше предыдущего, сбрасываем left
			left = right
		}
		maxLen = max(maxLen, right-left+1) // Обновляем maxLen, если текущая подпоследовательность длиннее
	}

	return maxLen // Возвращаем максимальную длину найденной подпоследовательности
}
