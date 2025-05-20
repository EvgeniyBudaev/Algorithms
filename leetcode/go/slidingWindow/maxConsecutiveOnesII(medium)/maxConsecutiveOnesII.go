package main

import "fmt"

/* 487. Max Consecutive Ones II
https://github.com/EvgeniyBudaev/doocs_leetcode/blob/main/solution/0400-0499/0487.Max%20Consecutive%20Ones%20II/README.md

Учитывая двоичный массив nums, если не более одного можно перевернуть 0, верните 1 максимальное количество
последовательных чисел в массиве.

Input: nums = [1,0,1,1,0]
Output: 4
Объяснение: переверните первый 0, чтобы получить самую длинную последовательную 1.
     После переворачивания максимальное количество последовательных единиц равно 4.

Input: nums = [1,0,1,1,0,1]
Output: 4
*/

func main() {
	nums := []int{1, 0, 1, 1, 0}
	fmt.Println(findMaxConsecutiveOnes(nums)) // 4
}

// findMaxConsecutiveOnes возвращает максимальное количество последовательных 1 в массиве при условии, что
// можно перевернуть 0 не более одного раза.
// time: O(n), space: O(1)
func findMaxConsecutiveOnes(nums []int) int {
	left, zeroCount, result, maxFlipOperations := 0, 0, 0, 1

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 { // Если ноль, то увеличиваем счетчик нулей
			zeroCount++
		}
		for zeroCount > maxFlipOperations { // Если превышено максимальное количество нулей, то сдвигаем левый указатель
			if nums[left] == 0 { // Если ноль, то уменьшаем счетчик нулей
				zeroCount--
			}
			left++
		}
		result = max(result, right-left+1)
	}

	return result // Возвращаем максимальное количество последовательных 1
}
