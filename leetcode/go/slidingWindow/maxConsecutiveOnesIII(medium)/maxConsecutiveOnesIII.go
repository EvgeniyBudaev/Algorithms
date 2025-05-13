package main

import "fmt"

/* 1004. Max Consecutive Ones III
https://leetcode.com/problems/max-consecutive-ones-iii/description/

Учитывая двоичный массив nums и целое число k, верните максимальное количество последовательных единиц в массиве,
если вы можете перевернуть не более k нулей.

Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
Output: 6
Explanation: [1,1,1,0,0,1,1,1,1,1,1]
Числа, выделенные жирным шрифтом, были перевернуты с 0 на 1. Самый длинный подмассив подчеркнут.

Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
Output: 10
Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
*/

func main() {
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	fmt.Println(longestOnes(nums, 2)) // 6
}

// longestOnes - возвращает максимальное количество последовательных единиц в массиве при условии, что
// можно перевернуть не более k нулей.
// time: O(n), space: O(1)
func longestOnes(nums []int, k int) int {
	left, zeroCount, result := 0, 0, 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 { // Если ноль, то увеличиваем счетчик нулей
			zeroCount++
		}
		for zeroCount > k { // Если превышено количество нулей, то сдвигаем левый указатель
			if nums[left] == 0 { // Если ноль, то уменьшаем счетчик нулей
				zeroCount--
			}
			left++
		}
		result = max(result, right-left+1)
	}

	return result // Возвращаем максимальную длину подмассива
}
