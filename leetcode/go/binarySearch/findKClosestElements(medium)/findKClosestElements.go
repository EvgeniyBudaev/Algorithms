package main

import "fmt"

/* https://leetcode.com/problems/find-k-closest-elements/description/

Учитывая отсортированный массив целых чисел arr, два целых числа k и x, верните k ближайших к x целых чисел в массиве.
Результат также следует отсортировать по возрастанию.

Целое число a ближе к x, чем целое число b, если:
|a - x| < |b - x|, or
|a - x| == |b - x| and a < b

Input: arr = [1,2,3,4,5], k = 4, x = 3
Output: [1,2,3,4]

Input: arr = [1,2,3,4,5], k = 4, x = -1
Output: [1,2,3,4]
*/

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(findClosestElements(arr, 4, 3)) // [1,2,3,4]
}

func findClosestElements(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-k

	for left < right {
		mid := (left + right) / 2
		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return arr[left : left+k]
}
