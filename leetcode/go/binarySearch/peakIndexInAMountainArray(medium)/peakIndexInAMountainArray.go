package main

import "fmt"

/* https://leetcode.com/problems/peak-index-in-a-mountain-array/description/

Вам дан целочисленный горный массив длиной n, значения которого увеличиваются до пикового элемента, а затем уменьшаются.
Верните индекс пикового элемента.
Ваша задача — решить ее за время сложности O(log(n)).

Input: arr = [0,1,0]
Output: 1

Input: arr = [0,2,1,0]
Output: 1

Input: arr = [0,10,5,2]
Output: 1
*/

func main() {
	arr := []int{0, 1, 0}
	fmt.Println(peakIndexInMountainArray(arr)) // 1
}

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid] < arr[mid+1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
