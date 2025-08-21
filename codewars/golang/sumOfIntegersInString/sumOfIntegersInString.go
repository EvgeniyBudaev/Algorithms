package main

import (
	"fmt"
	"unicode"
)

/* Sum of integers in string
https://www.codewars.com/kata/598f76a44f613e0e0b000026/train/go

Ваша задача в этой ката — реализовать функцию, вычисляющую сумму целых чисел в строке.
Например, в строке "The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog" сумма целых чисел равна 3635.

Примечание: будут проверяться только положительные целые числа.
*/

func main() {
	fmt.Println(SumOfIntegersInString("The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog")) // 3635
}

// SumOfIntegersInString возвращает сумму целых чисел в строке.
// time: O(n), space: O(1), где n - длина строки
func SumOfIntegersInString(strng string) int {
	sum := 0
	currentNumber := 0
	inNumber := false

	for _, char := range strng {
		if unicode.IsDigit(char) {
			// Если это цифра, добавляем её к текущему числу
			digit := int(char - '0')
			currentNumber = currentNumber*10 + digit
			inNumber = true
		} else {
			// Если это не цифра и мы были внутри числа, добавляем число к сумме
			if inNumber {
				sum += currentNumber
				currentNumber = 0
				inNumber = false
			}
		}
	}

	// Добавляем последнее число, если строка заканчивается цифрой
	if inNumber {
		sum += currentNumber
	}

	return sum
}
