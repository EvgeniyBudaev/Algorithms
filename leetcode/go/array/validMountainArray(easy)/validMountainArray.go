package main

import "fmt"

/* 941. Valid Mountain Array
https://leetcode.com/problems/valid-mountain-array/

Учитывая массив целых чисел arr, верните true если и только если он является допустимым массивом гор.
Напомним, что arr является горным массивом тогда и только тогда, когда:

arr.length >= 3
Существует некоторое i с 0 < i < arr.length - 1 таким , что:
arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

Input: arr = [2,1]
Output: false

Input: arr = [3,5,5]
Output: false

Input: arr = [0,3,2,1]
Output: true
*/

func main() {
	arr := []int{0, 3, 2, 1}
	fmt.Println(validMountainArray(arr)) // true
}

// validMountainArray - проверяет, является ли массив горным массивом.
// time: O(n), space: O(1)
func validMountainArray(arr []int) bool {
	// Если массив меньше 3 элементов, это не горный массив.
	if len(arr) <= 2 {
		return false
	}

	left := 0             // Левый указатель
	right := len(arr) - 1 // Правый указатель

	for c := 0; c < len(arr); c++ { // Пока не достигнут конец массива
		// Если элементы возрастают, двигаемся вправо.
		if arr[left] < arr[left+1] {
			left++
		}
		// Если элементы убывают, двигаемся влево.
		if arr[right] < arr[right-1] {
			right--
		}
		// Если левый или правый указатель достигли конца, это не горный массив.
		if right == len(arr)-1 || left == 0 {
			return false
		}
		// Если левый и правый указатели встретились, это горный массив.
		if left == right {
			return true
		}
	}

	return false // Если не удалось найти горный массив, возвращаем false.
}
