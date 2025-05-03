package main

import "fmt"

/* 81. Search in Rotated Sorted Array II
https://leetcode.com/problems/search-in-rotated-sorted-array-ii/description/

Существует целочисленный массив nums, отсортированный в неубывающем порядке (не обязательно с разными значениями).
Перед передачей в вашу функцию nums вращается с неизвестным индексом поворота k (0 <= k < nums.length), так что
результирующий массив имеет вид [nums[k], nums[k+1],..., nums [n-1], nums[0], nums[1], ..., nums[k-1]] (с индексом 0).
Например, [0,1,2,4,4,4,5,6,6,7] можно повернуть с опорным индексом 5 и превратить в [4,5,6,6,7,0,1,2, 4,4].
Учитывая числа массива после вращения и целочисленную цель, верните true, если цель находится в числах, или ложь, если
она не в числах.
Вы должны максимально сократить общее количество этапов операции.

Input: nums = [2,5,6,0,0,1,2], target = 0
Output: true

Input: nums = [2,5,6,0,0,1,2], target = 3
Output: false
*/

func main() {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	fmt.Println(search(nums, 0)) // true
}

// search - проверяет, содержит ли массив целевую цель.
func search(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}

// searchBinarySearch - использует бинарный поиск для поиска целевой цели в отсортированном массиве с возможным поворотом.
func searchBinarySearch(nums []int, target int) bool {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2

		// Проверяем, совпадает ли целевая цель с числом в середине массива
		if nums[mid] == target {
			return true
		}

		// Обработка дубликатов
		if nums[low] == nums[mid] { // Если левый элемент равен среднему элементу, сдвигаем левую границу
			low++
			continue
		}

		// Определение отсортированной половины
		// Левая половина отсортирована
		if nums[low] <= nums[mid] {
			if nums[low] <= target && target <= nums[mid] { // Проверяем находится ли целевая цель в отсортированной левой половине
				// Цель находится в левой половине
				high = mid - 1
			} else {
				// Иначе, цель находится в правой половине
				low = mid + 1
			}
		} else {
			// Правая половина отсортирована
			if nums[mid] <= target && target <= nums[high] { // Проверяем находится ли целевая цель в отсортированной правой половине
				// Цель находится в правой половине
				low = mid + 1
			} else {
				// Иначе, цель находится в левой половине
				high = mid - 1
			}
		}
	}

	// Цель не найдена
	return false
}
