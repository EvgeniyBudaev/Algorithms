package main

import "fmt"

/* 852. Peak Index in a Mountain Array
https://leetcode.com/problems/peak-index-in-a-mountain-array/description/

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

// peakIndexInMountainArray возвращает индекс пикового элемента в массиве arr.
// time: O(log n), space: O(1)
func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2 // Находим индекс среднего элемента
		// Проверяем, является ли текущий элемент пиковым элементом
		if arr[mid] < arr[mid+1] { // Если текущий элемент меньше следующего, двигаемся вправо
			left = mid + 1
			// Если текущий элемент больше следующего, двигаемся влево
		} else {
			right = mid - 1
		}
	}

	// Возвращаем индекс пикового элемента
	return left
}
