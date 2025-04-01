package main

import (
	"fmt"
	"sort"
)

/* 1338. Reduce Array Size to The Half

Вам дан целочисленный массив. Вы можете выбрать набор целых чисел и удалить все вхождения этих целых чисел в массиве.
Верните минимальный размер набора, чтобы была удалена хотя бы половина целых чисел массива.

Input: arr = [3,3,3,3,5,5,5,2,2,7]
Output: 2
Объяснение: При выборе {3,7} будет создан новый массив [5,5,5,2,2] размером 5
(т.е. равный половине размера старого массива).
Возможные наборы размера 2: {3,5},{3,2},{5,2}.
Выбор набора {2,7} невозможен, так как при этом будет создан новый массив [3,3,3,3,5,5,5],
размер которого превышает половину размера старого массива.

Input: arr = [7,7,7,7,7,7]
Output: 1
Пояснение: Единственный возможный набор, который вы можете выбрать, — это {7}. Это сделает новый массив пустым.
*/

func main() {
	arr := []int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}
	fmt.Println(minSetSize(arr)) // 2
}

func minSetSize(arr []int) int {
	freqMap := make(map[int]int)

	// Подсчитываем частоту каждого числа
	for _, num := range arr {
		freqMap[num]++
	}

	// Создаем слайс частот и сортируем по убыванию
	freqs := make([]int, 0, len(freqMap))
	for _, count := range freqMap {
		freqs = append(freqs, count)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(freqs)))

	total := len(arr)
	half := total / 2
	sum := 0
	size := 0

	// Суммируем частоты, пока не достигнем половины
	for _, count := range freqs {
		sum += count
		size++
		if sum >= half {
			return size
		}
	}

	return size
}
