package main

import "fmt"

/* 1493. Longest Subarray of 1's After Deleting One Element
https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/description/

Учитывая двоичный массив nums, вам следует удалить из него один элемент.
Верните размер самого длинного непустого подмассива, содержащего только 1 в результирующем массиве. Верните 0, если
такого подмассива нет.

Input: nums = [1,1,0,1]
Output: 3
Объяснение: После удаления числа в позиции 2 [1,1,1] содержит 3 числа со значением 1.

Input: nums = [0,1,1,1,0,1,1,0,1]
Output: 5
Пояснение: После удаления числа в позиции 4, [0,1,1,1,1,1,0,1] самым длинным подмассивом со значением единиц будет [1,1,1,1,1].

Input: nums = [1,1,1]
Output: 2
Объяснение: Вам необходимо удалить один элемент.
*/

func main() {
	nums := []int{1, 1, 0, 1}
	fmt.Println(longestSubarray(nums)) // 3
}

// longestSubarray возвращает размер самого длинного подмассива, содержащего только 1 в результирующем массиве.
// time: O(n), space: O(1)
func longestSubarray(nums []int) int {
	left, result, zeroCount := 0, 0, 0 // left - левый указатель, result - результат, zeroCount - количество нулей

	for right := 0; right < len(nums); right++ { // Проходим по массиву
		if nums[right] == 0 { // Если ноль, то увеличиваем количество нулей
			zeroCount++
		}

		for zeroCount > 1 { // Если количество нулей превышает 1
			if nums[left] == 0 { // Если ноль, то уменьшаем количество нулей
				zeroCount--
			}
			left++ // Сдвигаем левый указатель
		}

		result = max(result, right-left) // Обновляем максимальную длину подмассива
	}

	return result // Возвращаем размер самого длинного подмассива
}
