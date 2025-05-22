package main

import (
	"fmt"
	"sort"
)

/* 561. Array Partition
https://leetcode.com/problems/array-partition/description/

Дан массив целых чисел nums из 2n целых чисел, сгруппируйте эти целые числа в n пар (a1, b1), (a2, b2), ..., (an, bn)
так, чтобы сумма min(ai, bi) для всех i была максимальной. Верните максимальную сумму.

Input: nums = [1,4,3,2]
Output: 4
Пояснение: Все возможные пары (без учета порядка элементов):
1. (1, 4), (2, 3) -> min(1, 4) + min(2, 3) = 1 + 2 = 3
2. (1, 3), (2, 4) -> min(1, 3) + min(2, 4) = 1 + 2 = 3
3. (1, 2), (3, 4) -> min(1, 2) + min(3, 4) = 1 + 3 = 4
Таким образом, максимально возможная сумма — 4.
*/

func main() {
	nums := []int{1, 4, 3, 2}
	fmt.Println(arrayPairSum(nums)) // 4
}

// arrayPairSum возвращает максимальную сумму пар в массиве чисел.
// time: O(n), space: O(1)
func arrayPairSum(nums []int) int {
	sort.Ints(nums) // Сортируем массив O(n*log(n)) [1, 2, 3, 4]
	sum := 0        // Сумма пар

	// Берем каждый второй элемент (начиная с 0). time: O(n), space: O(1)
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i] // Добавляем его к сумме пар
	}

	return sum // Возвращаем сумму пар
}
