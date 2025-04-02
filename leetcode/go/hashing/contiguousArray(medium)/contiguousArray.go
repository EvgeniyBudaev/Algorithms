package main

import "fmt"

/* https://leetcode.com/problems/contiguous-array/description/

Учитывая числа двоичного массива, верните максимальную длину непрерывного подмассива с равным количеством 0 и 1.

Input: nums = [0,1]
Output: 2
Пояснение: [0, 1] — это самый длинный непрерывный подмассив с равным количеством 0 и 1.

Input: nums = [0,1,0]
Output: 2
Пояснение: [0, 1] (или [1, 0]) — это самый длинный непрерывный подмассив с равным количеством 0 и 1.
*/

func main() {
	nums := []int{0, 1}
	fmt.Println(findMaxLength(nums)) // 2
}

func findMaxLength(nums []int) int {
	maxLen := 0 // Текущая максимальная длина подмассива
	// count - Отслеживает разницу между количеством единиц и нулей в рассматриваемой части массива.
	// Если count положительно, это означает, что больше единиц, чем нулей; если отрицательно — наоборот.
	count := 0
	countMap := make(map[int]int)
	countMap[0] = -1

	for i, num := range nums {
		if num == 1 {
			count++
		} else {
			count--
		}

		// Поиск подмассивов с равным количеством нулей и единиц
		// Если да, то вычисляем длину текущего подмассива как разницу между текущим индексом i и индексом последнего
		// вхождения такого же count (map[count]). Обновляем maxLen, если текущая длина больше предыдущего максимума.
		if prevIndex, exists := countMap[count]; exists {
			currentLen := i - prevIndex
			if currentLen > maxLen {
				maxLen = currentLen
			}
		} else {
			// Если нет, то сохраняем текущий индекс i в map под ключом count.
			countMap[count] = i
		}
	}

	return maxLen
}
