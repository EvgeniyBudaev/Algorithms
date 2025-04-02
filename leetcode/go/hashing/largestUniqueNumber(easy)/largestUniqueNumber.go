package main

import "fmt"

/* 1133. Largest Unique Number

Учитывая массив целых чисел A, верните наибольшее целое число, которое встречается только один раз.
Если ни одно целое число не встречается один раз, верните -1.

Input: [5,7,3,9,4,9,8,3,1]
Output: 8
Explanation:
The maximum integer in the array is 9 but it is repeated. The number 8 occurs only once, so it's the answer.

Input: [9,9,8,8]
Output: -1
Explanation:
There is no number that occurs only once.
*/

func main() {
	nums := []int{5, 7, 3, 9, 4, 9, 8, 3, 1}
	fmt.Println(largestUniqueNumber(nums)) // 8
}

// largestUniqueNumber находит наибольшее уникальное число в массиве.
// Если такого числа нет, возвращает -1.
func largestUniqueNumber(nums []int) int {
	// Создаем мапу для подсчета количества вхождений каждого числа
	freq := make(map[int]int)

	// Заполняем мапу: ключ - число, значение - количество вхождений
	for _, num := range nums {
		freq[num]++
	}

	maxNum := -1 // Инициализируем максимальное значение как -1

	// Проходим по всем записям в мапе
	for num, count := range freq {
		// Если число встречается 1 раз и больше текущего максимума, обновляем максимум
		if count == 1 && num > maxNum {
			maxNum = num
		}
	}

	return maxNum
}
