package main

import (
	"fmt"
	"sort"
)

/* 1385. Find the Distance Value Between Two Arrays
https://leetcode.com/problems/find-the-distance-value-between-two-arrays/description/

Даны два целочисленных массива arr1 и arr2, и целое число d, вернуть значение расстояния между двумя массивами.
Значение расстояния определяется как количество элементов arr1[i], таких, что нет ни одного элемента arr2[j],
где |arr1[i]-arr2[j]| <= d.

Input: arr1 = [4,5,8], arr2 = [10,9,1,8], d = 2
Output: 2
Explanation:
For arr1[0]=4 we have:
|4-10|=6 > d=2
|4-9|=5 > d=2
|4-1|=3 > d=2
|4-8|=4 > d=2
For arr1[1]=5 we have:
|5-10|=5 > d=2
|5-9|=4 > d=2
|5-1|=4 > d=2
|5-8|=3 > d=2
For arr1[2]=8 we have:
|8-10|=2 <= d=2
|8-9|=1 <= d=2
|8-1|=7 > d=2
|8-8|=0 <= d=2
*/

func main() {
	arr1 := []int{4, 5, 8}
	arr2 := []int{10, 9, 1, 8}
	d := 2                                           // d - дистанция
	fmt.Println(findTheDistanceValue(arr1, arr2, d)) // 2
}

// findTheDistanceValue - возвращает значение расстояния между двумя массивами.
// time: O(n log n), space: O(1)
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2) // Сортируем arr2
	count := 0      // Инициализируем счетчик подходящих элементов

	// Проходим по каждому элементу x в arr1
	for _, x := range arr1 {
		// Используем бинарный поиск для поиска ближайшего элемента в arr2
		index := sort.SearchInts(arr2, x)

		// Проверяем, выполняется ли условие |x - y| > d для найденного элемента и его соседей
		valid := true // Флаг, указывающий на то, что элемент подходит
		// Если элемент находится в arr2, проверяем его соседей
		if index < len(arr2) {
			if abs(x-arr2[index]) <= d { // Если условие не выполняется, устанавливаем флаг в false
				valid = false
			}
		}
		// Если элемент находится в arr2, проверяем его соседей
		if index > 0 {
			if abs(x-arr2[index-1]) <= d { // Если условие не выполняется, устанавливаем флаг в false
				valid = false
			}
		}

		// Если элемент x подходит, увеличиваем счетчик
		if valid {
			count++
		}
	}

	// Возвращаем значение счетчика
	return count
}

// abs - возвращает абсолютное значение числа.
// time: O(1), space: O(1)
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
