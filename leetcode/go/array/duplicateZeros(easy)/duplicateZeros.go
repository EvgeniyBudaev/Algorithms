package main

import "fmt"

/* 1089. Duplicate Zeros
https://leetcode.com/problems/duplicate-zeros/description/

Учитывая целочисленный массив фиксированной длины arr, дублируйте каждое вхождение нуля, сдвигая оставшиеся элементы
вправо. Обратите внимание, что элементы, превышающие длину исходного массива, не записываются. Внесите вышеуказанные
изменения во входной массив и ничего не возвращайте.

Input: arr = [1,0,2,3,0,4,5,0]
Output: [1,0,0,2,3,0,0,4]
Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]

Input: arr = [1,2,3]
Output: [1,2,3]
Explanation: After calling your function, the input array is modified to: [1,2,3]
*/

func main() {
	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(arr)
	fmt.Println(arr) // [1,0,0,2,3,0,0,4]
}

// duplicateZeros дублирует каждое вхождение нуля, сдвигая оставшиеся элементы вправо.
// Элементы, превышающие длину исходного массива, не записываются.
// time: O(n), space: O(1)
func duplicateZeros(arr []int) {
	for i := 0; i < len(arr); i++ { // Проходим по массиву
		if arr[i] == 0 && len(arr) != 0 { // Если ноль и длина массива не равна 0
			copy(arr[i+1:], arr[i:len(arr)-1]) // Копируем элементы вправо
			i++                                // Увеличиваем индекс на 1
		}
	}
}
