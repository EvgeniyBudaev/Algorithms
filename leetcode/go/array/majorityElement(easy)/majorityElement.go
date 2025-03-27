package main

import "fmt"

/* https://leetcode.com/problems/majority-element/description/

Учитывая массив nums размера n, верните элемент большинства.
Мажоритарным элементом является элемент, который появляется более ⌊n / 2⌋ раз. Вы можете предположить, что элемент
большинства всегда существует в массиве.

Input: nums = [3,2,3]
Output: 3

Input: nums = [2,2,1,1,1,2,2]
Output: 2

Input: nums = [3,3,4]
Output: 3
*/

func main() {
	arr := []int{3, 2, 3}
	fmt.Println(majorityElement(arr)) // 3
}

func majorityElement(nums []int) int {
	count := 0
	candidate := 0

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
		}
		if candidate == nums[i] {
			count++
		} else {
			count--
		}
	}

	return candidate
}
