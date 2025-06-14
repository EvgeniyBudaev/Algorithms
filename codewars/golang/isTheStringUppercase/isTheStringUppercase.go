package main

import (
	"fmt"
	"strings"
)

/* Is the string uppercase?
https://www.codewars.com/kata/56cd44e1aa4ac7879200010b/train/go

Строка в верхнем регистре?
Задача:
Создайте метод, чтобы проверить, состоит ли строка из ВСЕХ ЗАГЛАВНЫХ букв.

Примеры (ввод -> вывод):
"c" -> False
"C" -> True
"hello I AM DONALD" -> False
"HELLO I AM DONALD" -> True
"ACSKLDFJSgSKLDFJSKLDFJ" -> False
"ACSKLDFJSGSKLDFJSKLDFJ" -> True
*/

func main() {
	m1 := MyString("ACSKLDFJSGSKLDFJSKLDFJ")
	fmt.Println(m1.IsUpperCase()) // true

	m2 := MyString("ABCDEFGHIJKLMNOPQRSTUVWXYz")
	fmt.Println(m2.IsUpperCase()) // false

	m3 := MyString("False")
	fmt.Println(m3.IsUpperCase()) // false

	m4 := MyString("C")
	fmt.Println(m4.IsUpperCase()) // true
}

type MyString string

// IsUpperCase проверяет, состоит ли строка из ВСЕХ ЗАГЛАВНЫХ букв.
// time: O(n), space: O(1)
func (s MyString) IsUpperCase() bool {
	return s == MyString(strings.ToUpper(string(s)))
}
