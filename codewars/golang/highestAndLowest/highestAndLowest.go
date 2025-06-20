package main

import (
	"fmt"
	"strconv"
	"strings"
)

/* Highest and Lowest
https://www.codewars.com/kata/554b4ac871d6813a03000035/train/go

В этом небольшом задании вам дана строка чисел, разделенных пробелами, и вам нужно вернуть наибольшее и наименьшее число.

Examples:
HighAndLow("1 2 3 4 5") // return "5 1"
HighAndLow("1 2 -3 4 5") // return "5 -3"
HighAndLow("1 9 3 4 -5") // return "9 -5"
*/

func main() {
	fmt.Println(HighAndLow("1 9 3 4 -5"))                   // "9 -5"
	fmt.Println(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4")) // "42 -9"
	fmt.Println(HighAndLow("1 2 3"))                        // "3 1"
}

// HighAndLow возвращает наибольшее и наименьшее число.
// time: O(n), space: O(1)
func HighAndLow(in string) string {
	numbers := strings.Split(in, " ") // Разбиваем строку на части по пробелам

	// Инициализируем min и max первым числом
	minNumber, _ := strconv.Atoi(numbers[0])
	maxNumber := minNumber

	// Проходим по всем числам
	for _, numStr := range numbers {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			continue // Пропускаем некорректные значения
		}
		if num < minNumber {
			minNumber = num
		}
		if num > maxNumber {
			maxNumber = num
		}
	}

	return fmt.Sprintf("%v %v", maxNumber, minNumber) // Форматируем результат
}
