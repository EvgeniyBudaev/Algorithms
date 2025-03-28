package main

import "fmt"

/* https://leetcode.com/problems/subarray-sum-equals-k/description/

Учитывая массив целых чисел nums и целое число k, верните общее количество подмассивов, сумма которых равна k.
Подмассив — это непрерывная непустая последовательность элементов массива.

Input: nums = [1,1,1], k = 2
Output: 2

Input: nums = [1,2,3], k = 3
Output: 2
*/

func main() {
	nums := []int{1, 1, 1}
	k := 2
	fmt.Println(subarraySum(nums, k)) // 2
}

func subarraySum(nums []int, k int) int {
	// Создаем map для хранения сумм и их количества
	sumMap := make(map[int]int)
	sumMap[0] = 1 // Инициализируем с суммой 0, встречающейся 1 раз

	sum := 0
	count := 0

	// Проходим по массиву
	for _, num := range nums {
		sum += num
		// Если (sum - k) существует в map, увеличиваем count
		if val, ok := sumMap[sum-k]; ok {
			count += val
		}
		// Обновляем map с текущей суммой
		sumMap[sum]++
	}

	return count
}
