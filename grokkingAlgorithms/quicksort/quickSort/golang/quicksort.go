package main

import "fmt"

func main() {
	fmt.Println(quicksort([]int{10, 5, 2, 3})) // [2, 3, 5, 10]
}

// quicksort быстрая сортировка массива O(n*log(n))
func quicksort(list []int) []int {
	// Базовый случай
	if len(list) < 2 {
		return list
		// Рекурсивный случай
	} else {
		pivot := list[0] // Опорный элемент

		var less = []int{}    // Подмассив всех элементов меньше опорного
		var greater = []int{} // Подмассив всех элементов больше опорного

		// Фильтруем массив, разбивая его на два подмассива
		for _, num := range list[1:] {
			// Если элемент меньше опорного, добавляем его в подмассив "меньше", иначе - в подмассив "больше"
			if pivot > num {
				less = append(less, num)
			} else {
				greater = append(greater, num)
			}
		}

		less = append(quicksort(less), pivot)
		greater = quicksort(greater)
		return append(less, greater...)
	}
}
