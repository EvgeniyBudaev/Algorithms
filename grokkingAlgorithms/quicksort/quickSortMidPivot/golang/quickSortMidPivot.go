package main

import "fmt"

func main() {
	fmt.Println(quickSort([]int{10, 5, 2, 3})) // [2, 3, 5, 10]
}

// quickSort быстрая сортировка массива O(n*log(n))
// time: O(n*log(n)), space: O(log(n))
func quickSort(array []int) []int {
	// Базовый случай: если массив содержит 0 или 1 элемент
	if len(array) < 2 {
		return array
	}

	// Выбираем опорный элемент (pivot)
	pivotIndex := len(array) / 2 // Здесь мы выбираем опорный элемент как средний элемент массива
	pivot := array[pivotIndex]   // Получаем значение опорного элемента

	// Создаем слайсы для элементов меньше и больше опорного
	var less, greater []int

	// Распределяем элементы по слайсам
	for i, num := range array {
		if i == pivotIndex { // Пропускаем опорный элемент
			continue
		}
		if num <= pivot { // Если элемент меньше или равен опорному, добавляем его в слайс "меньше"
			less = append(less, num)
		} else { // Иначе, добавляем его в слайс "больше"
			greater = append(greater, num)
		}
	}

	// Рекурсивно сортируем подмассивы и объединяем результат
	return append(
		append(quickSort(less), pivot),
		quickSort(greater)...,
	)
}

// quickSort быстрая сортировка массива O(n*log(n)) (in-place)
// time: O(n*log(n)), space: O(log(n))
//func quickSort(arr []int) []int {
//	if len(arr) < 2 {
//		return arr
//	}
//
//	left, right := 0, len(arr)-1
//	pivotIndex := len(arr) / 2                                // Опорный элемент
//	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex] // перемещаем опорный элемент в конец массива
//
//	for i := range arr {
//		if arr[i] < arr[right] {
//			arr[left], arr[i] = arr[i], arr[left]
//			left++
//		}
//	}
//
//	arr[left], arr[right] = arr[right], arr[left]
//
//	quickSort(arr[:left])
//	quickSort(arr[left+1:])
//
//	return arr
//}
