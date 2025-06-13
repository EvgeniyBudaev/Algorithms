package main

import "fmt"

/* 658. Find K Closest Elements
https://leetcode.com/problems/find-k-closest-elements/description/

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

// findClosestElements возвращает k ближайших к x целых чисел в массиве.
// time: O(log(n-k) + k), space: O(1), где n - длина массива arr, k - количество ближайших чисел.
func findClosestElements(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-k

	for left < right {
		mid := (left + right) / 2 // Поиск середины между левым и правым указателями
		// Сравнение расстояния между элементом в середине и x с расстоянием между элементом, находящимся на k позициях
		if x-arr[mid] > arr[mid+k]-x { // Если расстояние до левого элемента больше, чем до правого элемента, двигаем левую границу
			left = mid + 1
		} else {
			right = mid
		}
	}

	// Возвращаем срез, начиная с позиции, найденной выше, и длиной k.
	return arr[left : left+k]
}
