package main

import "fmt"

/*
Дан массив целых положительных чисел nums и целое число k. Найдите длину самого длинного подмассива, сумма которого
меньше или равна k.

Input: nums = [3, 1, 2, 7, 4, 2, 1, 1, 5], k = 8
Output: 4
*/

func main() {
	nums := []int{3, 1, 2, 7, 4, 2, 1, 1, 5}
	fmt.Println(findLengthSubarray(nums, 8)) // 4
}

// findLengthSubarray находит длину самого длинного подмассива, сумма которого меньше или равна k.
// time: O(n), space: O(1)
func findLengthSubarray(nums []int, k int) int {
	left, sum, result := 0, 0, 0 // Инициализируем левый указатель, сумму и результат

	for right := 0; right < len(nums); right++ { // Перебираем все элементы массива
		sum += nums[right] // Добавляем текущий элемент в сумму
		// Если сумма превышает k
		for sum > k {
			sum -= nums[left] // Убираем элемент из суммы
			left++            // Сдвигаем левый указатель
		}
		// Обновляем максимальную длину подмассива
		result = max(result, right-left+1)
	}

	return result // Возвращаем длину самого длинного подмассива
}
