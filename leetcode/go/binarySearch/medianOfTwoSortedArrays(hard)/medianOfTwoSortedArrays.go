package main

import (
	"fmt"
	"math"
)

/* https://leetcode.com/problems/median-of-two-sorted-arrays/description/
solution https://leetcode.com/problems/median-of-two-sorted-arrays/solutions/4070371/94-96-binary-search-two-pointers/

Учитывая два отсортированных массива nums1 и nums2 размера m и n соответственно, верните медиану двух отсортированных
массивов.
Общая сложность времени выполнения должна составлять O(log (m+n)).

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
*/

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2)) // 2
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Обеспечиваем, чтобы nums1 был более коротким массивом
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	low, high := 0, m

	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (m+n+1)/2 - partitionX

		// Определяем max и min для каждого раздела
		var maxX, maxY int
		if partitionX == 0 {
			maxX = math.MinInt32
		} else {
			maxX = nums1[partitionX-1]
		}

		if partitionY == 0 {
			maxY = math.MinInt32
		} else {
			maxY = nums2[partitionY-1]
		}

		var minX, minY int
		if partitionX == m {
			minX = math.MaxInt32
		} else {
			minX = nums1[partitionX]
		}

		if partitionY == n {
			minY = math.MaxInt32
		} else {
			minY = nums2[partitionY]
		}

		// Проверяем условие нахождения медианы
		if maxX <= minY && maxY <= minX {
			if (m+n)%2 == 0 {
				return float64(max(maxX, maxY)+min(minX, minY)) / 2.0
			} else {
				return float64(max(maxX, maxY))
			}
		} else if maxX > minY {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}

	panic("Input arrays are not sorted")
}
