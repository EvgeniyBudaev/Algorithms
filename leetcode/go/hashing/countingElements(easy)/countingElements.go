package main

import (
	"fmt"
)

/* 1426. Counting Elements

Учитывая целочисленный массив arr, подсчитайте, сколько элементов x существует так, что x + 1 также находится в arr.
Если в arr есть дубликаты, посчитайте их отдельно.

Input: arr = [1,2,3]
Output: 2
Пояснение: 1 и 2 учитываются, потому что 2 и 3 находятся в arr.

Input: arr = [1,1,3,3,5,5,7,7]
Output: 0
Пояснение: Никакие числа не учитываются, поскольку в arr нет цифр 2, 4, 6 или 8.

Input: arr = [1,3,2,3,5,0]
Output: 3
Пояснение: 0, 1 и 2 учитываются, потому что 1, 2 и 3 находятся в arr.

Input: arr = [1,1,2,2]
Output: 2
Пояснение: Две единицы засчитываются, потому что 2 находится в arr.

Input: arr = [1,1,2]
Output: 2
Объяснение: Обе единицы учитываются, поскольку в массиве есть 2.
*/

func main() {
	fmt.Println(countElements([]int{1, 2, 3})) // 2
	fmt.Println(countElements([]int{1, 1, 2})) // 2
}

// countElements подсчитывает количество элементов x в массиве, таких что x + 1 также находится в массиве.
func countElements(arr []int) int {
	// Создаем множество (map) для быстрой проверки наличия элементов в массиве.
	elementSet := make(map[int]bool)
	for _, num := range arr {
		elementSet[num] = true
	}

	count := 0
	// Проходим по всем элементам массива и проверяем, есть ли (x + 1) в множестве.
	for _, num := range arr {
		if elementSet[num+1] {
			count++
		}
	}

	return count
}
