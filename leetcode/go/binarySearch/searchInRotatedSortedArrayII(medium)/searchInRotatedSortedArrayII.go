package main

import "fmt"

/* https://leetcode.com/problems/search-in-rotated-sorted-array-ii/description/

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

func search(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}

func searchBinarySearch(nums []int, target int) bool {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2

		if nums[mid] == target {
			return true
		}

		// Обработка дубликатов
		if nums[low] == nums[mid] {
			low++
			continue
		}

		// Определение отсортированной половины
		if nums[low] <= nums[mid] {
			// Левая половина отсортирована
			if nums[low] <= target && target <= nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			// Правая половина отсортирована
			if nums[mid] <= target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return false
}
