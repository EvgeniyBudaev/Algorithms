package main

import "fmt"

/* 713. Subarray Product Less Than K
https://leetcode.com/problems/subarray-product-less-than-k/description/

Учитывая массив положительных целых чисел nums и целое число k, верните количество подмассивов, в которых произведение
всех элементов в подмассиве строго меньше k.

Input: nums = [10,5,2,6], k = 100
Output: 8
Пояснение: 8 подмассивов с произведением меньше 100:
[10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6]
Обратите внимание, что [10, 5, 2] не включены, поскольку произведение 100 не строго меньше k.

Input: nums = [1,2,3], k = 0
Output: 0
*/

func main() {
	nums := []int{10, 5, 2, 6}
	fmt.Println(numSubarrayProductLessThanK(nums, 100)) // 8
}

// numSubarrayProductLessThanK возвращает количество подмассивов, в которых произведение всех элементов в подмассиве строго меньше k.
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	left, curr, result := 0, 1, 0

	for right := 0; right < len(nums); right++ {
		curr *= nums[right]
		// Если произведение превышает или равно k, сдвигаем левую границу окна
		for curr >= k {
			curr /= nums[left]
			left++
		}
		// Добавляем количество подмассивов, заканчивающихся на right
		result += right - left + 1
	}

	return result
}
