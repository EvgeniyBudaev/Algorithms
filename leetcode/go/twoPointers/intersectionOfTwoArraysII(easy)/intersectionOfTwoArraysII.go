package main

import (
	"fmt"
	"sort"
)

/* 350. Intersection of Two Arrays II
https://leetcode.com/problems/intersection-of-two-arrays-ii/description/

Учитывая два целочисленных массива nums1 и nums2, верните массив их пересечения. Каждый элемент результата должен
появляться столько раз, сколько он отображается в обоих массивах, и вы можете возвращать результат в любом порядке.

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
Пояснение: [9,4] также принимается.
*/

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersect(nums1, nums2)) // [2,2]
}

// intersect - находит пересечение двух массивов nums1 и nums2.
// time: O(n), space: O(n)
func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1) // [1 1 2 2]
	sort.Ints(nums2) // [2 2]
	left, right := 0, 0
	result := make([]int, 0)

	for left < len(nums1) && right < len(nums2) {
		if nums1[left] < nums2[right] {
			left++
		} else if nums1[left] > nums2[right] {
			right++
		} else {
			result = append(result, nums1[left])
			left++
			right++
		}
	}

	return result
}
