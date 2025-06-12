package main

import (
	"fmt"
)

/* Find the smallest integer in the array
https://www.codewars.com/kata/55a2d7ebe362935a210000b2/train/go

Учитывая массив целых чисел, ваше решение должно найти наименьшее целое число.

Например:
Если задано [34, 15, 88, 2], ваше решение вернет 2
Если задано [34, -345, -1, 100], ваше решение вернет -345

Вы можете предположить, что предоставленный массив не будет пустым.
*/

func main() {
	numbers1 := []int{34, 15, 88, 2}
	fmt.Println(SmallestIntegerFinder(numbers1)) // 2

	numbers2 := []int{34, -345, -1, 100}
	fmt.Println(SmallestIntegerFinder(numbers2)) // -345
}

func SmallestIntegerFinder(numbers []int) int {
	minNumber := numbers[0]

	for _, num := range numbers {
		if num < minNumber {
			minNumber = num
		}
	}

	return minNumber
}
