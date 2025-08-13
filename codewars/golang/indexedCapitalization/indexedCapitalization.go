package main

import (
	"fmt"
	"unicode"
)

/* Indexed capitalization
https://www.codewars.com/kata/59cfc09a86a6fdf6df0000f1/train/go

Даны строка строчных букв и массив целочисленных индексов. Сделать все буквы с заданными индексами заглавными.
Если индекс выходит за пределы строки, его следует игнорировать.

Примеры:
"abcdef", [1,2,5] ==> "aBCdeF"
"abcdef", [1,2,5,100] ==> "aBCdeF" // Индекса 100 нет.
*/

func main() {
	arr := []int{1, 2, 5}
	fmt.Println(Capitalize("abcdef", arr)) // "aBCdeF"
}

// Capitalize делает заглавными буквы с заданными индексами.
// time: O(n), space: O(n), где n - длина строки
func Capitalize(st string, arr []int) string {
	runes := []rune(st)
	for _, index := range arr {
		if index >= 0 && index < len(runes) {
			runes[index] = unicode.ToUpper(runes[index])
		}
	}
	return string(runes)
}
