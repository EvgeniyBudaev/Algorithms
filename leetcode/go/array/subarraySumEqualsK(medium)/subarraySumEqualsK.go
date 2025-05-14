package main

import "fmt"

/* 560. Subarray Sum Equals K
https://leetcode.com/problems/subarray-sum-equals-k/description/

Учитывая массив целых чисел nums и целое число k, верните общее количество подмассивов, сумма которых равна k.
Подмассив — это непрерывная непустая последовательность элементов массива.

Input: nums = [1,1,1], k = 2
Output: 2

Input: nums = [1,2,3], k = 3
Output: 2
*/

func main() {
	nums := []int{1, 2, 3}
	k := 3
	fmt.Println(subarraySum(nums, k)) // 2
}

// subarraySum - находит количество подмассивов с суммой, равной k.
// time: O(n), space: O(n)
func subarraySum(nums []int, k int) int {
	sumMap := make(map[int]int) // Создаем map для хранения сумм и их количества
	sumMap[0] = 1               // Инициализируем с суммой 0, встречающейся 1 раз

	sum := 0   // Текущая сумма
	count := 0 // Количество подмассивов с суммой k

	// Проходим по массиву
	for _, num := range nums {
		sum += num // Добавляем текущее число к текущей сумме
		// Если (sum - k) существует в map, увеличиваем count
		if val, ok := sumMap[sum-k]; ok {
			count += val
		}
		// Обновляем map с текущей суммой
		sumMap[sum]++
	}

	// Возвращаем количество подмассивов с суммой k
	return count
}
