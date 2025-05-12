package main

import (
	"fmt"
)

/* 349. Intersection of Two Arrays
https://leetcode.com/problems/intersection-of-two-arrays/description/

Даны два целочисленных массива nums1 и nums2, вернуть массив их пересечения. Каждый элемент в результате должен быть
уникальным, и вы можете вернуть результат в любом порядке.

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
*/

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2)) // [2]
}

// intersection возвращает массив пересечения элементов двух массивов nums1 и nums2.
// time: O(n+m), space: O(n)
func intersection(nums1 []int, nums2 []int) []int {
	// Создаем множества из массивов
	set1 := make(map[int]bool)
	for _, num := range nums1 {
		set1[num] = true
	}

	set2 := make(map[int]bool)
	for _, num := range nums2 {
		set2[num] = true
	}

	// Находим пересечение множеств
	var result []int
	for num := range set1 {
		if set2[num] { // Если элемент присутствует в обоих множествах, добавляем его в результат
			result = append(result, num)
		}
	}

	return result // Возвращаем результат
}
