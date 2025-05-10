package main

import (
	"fmt"
	"math/rand"
)

/* 912. Sort an Array (Quick Sort)
https://leetcode.com/problems/sort-an-array/description/

Дан массив целых чисел nums, отсортируйте массив в порядке возрастания и верните его.

Вы должны решить задачу без использования каких-либо встроенных функций с временной сложностью O(nlog(n)) и с наименьшей
возможной пространственной сложностью

Input: nums = [5,2,3,1]
Output: [1,2,3,5]
Пояснение: После сортировки массива позиции некоторых чисел не меняются (например, 2 и 3), а позиции других чисел
меняются (например, 1 и 5).
*/

func main() {
	nums := []int{5, 2, 3, 1}
	fmt.Println(sortArray(nums)) // [1,2,3,5]
}

// sortArray - быстрая сортировка массива.
// time: O(nlog(n)), space: O(1)
func sortArray(nums []int) []int {
	qsort(nums, 0, len(nums))
	return nums
}

// getPivot - выбирает случайный pivot в диапазоне [l, r)
// time: O(1), space: O(1)
func getPivot(arr []int, l, r int) (pivot, pivotIdx int) {
	pivotIdx = rand.Intn(r-l) + l  // Генерируем случайный индекс в диапазоне [l, r)
	return arr[pivotIdx], pivotIdx // Возвращаем pivot и его индекс
}

// partition - разделяет массив на две части относительно pivot
// time: O(n), space: O(1)
func partition(arr []int, lIdx, rIdx int) int {
	pivot, pivotIdx := getPivot(arr, lIdx, rIdx) // Получаем pivot и его индекс

	// Перемещаем pivot в начало
	arr[lIdx], arr[pivotIdx] = arr[pivotIdx], arr[lIdx] // Меняем местами первый и pivot элементы
	l, r := lIdx+1, rIdx-1                              // Устанавливаем указатели на концы массива

	for l <= r {
		if arr[l] < pivot { // Если элемент слева меньше pivot, двигаем левый указатель
			l++
		} else if arr[r] > pivot { // Если элемент справа больше pivot, двигаем правый указатель
			r--
		} else { // Иначе меняем местами элементы и двигаем указатели
			arr[l], arr[r] = arr[r], arr[l] // Меняем местами элементы
			l++
			r--
		}
	}

	// Возвращаем pivot на правильную позицию
	arr[lIdx], arr[r] = arr[r], arr[lIdx]
	return r // Возвращаем индекс pivot после сортировки
}

// qsort - рекурсивная функция быстрой сортировки
// time: O(nlog(n)), space: O(1)
func qsort(nums []int, l, r int) {
	if l >= r {
		return
	}
	pivotIdx := partition(nums, l, r) // Получаем индекс pivot после сортировки
	qsort(nums, l, pivotIdx)          // Рекурсивно сортируем левую часть массива
	qsort(nums, pivotIdx+1, r)        // Рекурсивно сортируем правую часть массива
}
