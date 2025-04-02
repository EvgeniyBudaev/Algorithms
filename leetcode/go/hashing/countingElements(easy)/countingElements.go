package main

import (
	"fmt"
	"sort"
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
}

func countElements(arr []int) int {
	// Создаем map для подсчета количества каждого элемента
	countMap := make(map[int]int)
	for _, num := range arr {
		countMap[num]++
	}

	// Создаем slice с уникальными элементами
	uniqueNums := make([]int, 0, len(countMap))
	for num := range countMap {
		uniqueNums = append(uniqueNums, num)
	}

	// Если уникальных элементов меньше 3, возвращаем 0
	if len(uniqueNums) < 3 {
		return 0
	}

	// Сортируем уникальные элементы
	sort.Ints(uniqueNums)

	// Исключаем минимальный и максимальный элементы
	middleNums := uniqueNums[1 : len(uniqueNums)-1]

	// Суммируем количество оставшихся элементов
	total := 0
	for _, num := range middleNums {
		total += countMap[num]
	}

	return total
}
