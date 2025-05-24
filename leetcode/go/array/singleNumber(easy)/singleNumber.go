package main

import (
	"fmt"
	"sort"
)

/* 136. Single Number
https://leetcode.com/problems/single-number/description/

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
	fmt.Println(singleNumber([]int{2, 2, 1})) // 1
}

// singleNumber находит уникальное число в слайсе.
// time: O(n), space: O(1)
func singleNumber(nums []int) int {
	sort.Ints(nums) // time: O(n logn), space: O(1)

	// Цикл по всем элементам, кроме последнего. time: O(n), space: O(1)
	for i := 0; i < len(nums)-1; i += 2 {
		// Если текущий элемент не равен следующему, то текущий элемент является уникальным.
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}

	// Если все элементы встречаются дважды, то последний элемент будет уникальным.
	return nums[len(nums)-1]
}
