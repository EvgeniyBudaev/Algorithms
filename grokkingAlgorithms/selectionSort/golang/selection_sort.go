package main

import "fmt"

func main() {
	s := []int{5, 3, 6, 2, 10}
	fmt.Println(selectionSort(s)) // [2, 3, 5, 6, 10]
}

// findSmallest возвращает индекс наименьшего элемента массива. O(n)
func findSmallest(arr []int) int {
	// Сохраняет наименьшее значение
	smallest := arr[0]

	// Сохраняет индекс наименьшего значения
	smallest_index := 0
	for i := 1; i < len(arr); i++ {
		// Проверяет наименьший элемент массива
		if arr[i] < smallest {
			smallest = arr[i]
			smallest_index = i
		}
	}
	return smallest_index
}

// selectionSort сортирует массив. O(n^2)
func selectionSort(arr []int) []int {
	size := len(arr)
	newArr := make([]int, size)
	for i := 0; i < size; i++ {
		// Находит наименьший элемент в массиве и добавляет его в новый массив
		smallest := findSmallest(arr)
		newArr[i] = arr[smallest]
		arr = append(arr[:smallest], arr[smallest+1:]...)
	}
	return newArr
}
