package main

import (
	"fmt"
	"unicode"
)

/* Simple string characters
https://www.codewars.com/kata/5a29a0898f27f2d9c9000058/train/go

В этой задаче вам дана строка, и ваша задача — вернуть список целых чисел с подробным указанием количества заглавных и
строчных букв, цифр и специальных символов (всех остальных).

Порядок следующий: заглавные буквы, строчные буквы, цифры и специальные символы.

"*'&ABCDabcde12345" --> [ 4, 5, 5, 3 ]
*/

func main() {
	fmt.Println(Solve("*'&ABCDabcde12345")) // [ 4, 5, 5, 3 ]
}

// Solve возвращает список целых чисел с подробным указанием количества заглавных и строчных букв, цифр и специальных символов.
// time: O(n), space: O(1), где n - длина строки
func Solve(s string) []int {
	upper := 0   // Количество заглавных букв
	lower := 0   // Количество строчных букв
	digit := 0   // Количество цифр
	special := 0 // Количество специальных символов

	for _, char := range s {
		if unicode.IsUpper(char) {
			upper++
		} else if unicode.IsLower(char) {
			lower++
		} else if unicode.IsDigit(char) {
			digit++
		} else {
			special++
		}
	}

	return []int{upper, lower, digit, special}
}
