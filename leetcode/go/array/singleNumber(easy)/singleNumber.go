package main

import (
	"fmt"
	"sort"
)

/* https://leetcode.com/problems/single-number/description/

Учитывая непустой массив целых чисел nums, каждый элемент появляется дважды, кроме одного. Найдите этого единственного.
Вы должны реализовать решение с линейной сложностью времени выполнения и использовать только постоянное дополнительное
пространство.

Input: nums = [2,2,1]
Output: 1

Input: nums = [4,1,2,1,2]
Output: 4

Input: nums = [1]
Output: 1
*/

func main() {
	nums := []int{2, 2, 1}
	fmt.Println(singleNumber(nums)) // 1
}

func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}
