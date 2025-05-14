package main

import "fmt"

/* 2215. Find the Difference of Two Arrays

Учитывая два целочисленных массива nums1 и nums2 с индексом 0, верните ответ в виде списка размером 2, где:

ответ[0] — это список всех различных целых чисел в nums1, которых нет в nums2.
ответ[1] — это список всех различных целых чисел в nums2, которых нет в nums1.
Обратите внимание, что целые числа в списках могут возвращаться в любом порядке.

Input: nums1 = [1,2,3], nums2 = [2,4,6]
Output: [[1,3],[4,6]]
Объяснение:
Для nums1 nums1[1] = 2 присутствует в индексе 0 в nums2, тогда как nums1[0] = 1 и nums1[2] = 3 отсутствуют в nums2.
Следовательно, ответ[0] = [1,3].
Для nums2 nums2[0] = 2 присутствует в индексе 1 nums1, тогда как nums2[1] = 4 и nums2[2] = 6 отсутствуют в nums2.
Следовательно, ответ[1] = [4,6].
*/

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{2, 4, 6}
	fmt.Println(findDifference(arr1, arr2)) // [[1,3],[4,6]]
}

// findDifference - возвращает списки элементов, которых нет в другом масиве.
// time: O(n), space: O(n)
func findDifference(nums1 []int, nums2 []int) [][]int {
	s1 := make(map[int]bool) // Множество элементов первого массива
	s2 := make(map[int]bool) // Множество элементов второго массива

	// Заполняем множества
	for _, num := range nums1 {
		s1[num] = true
	}
	for _, num := range nums2 {
		s2[num] = true
	}

	// Удаляем элементы, которые есть в обоих множествах
	for num := range s1 {
		if s2[num] {
			delete(s1, num)
			delete(s2, num)
		}
	}

	// Преобразуем множества обратно в слайсы
	result1 := make([]int, 0, len(s1))
	for num := range s1 {
		result1 = append(result1, num)
	}

	result2 := make([]int, 0, len(s2))
	for num := range s2 {
		result2 = append(result2, num)
	}

	// Возвращаем результат
	return [][]int{result1, result2}
}
