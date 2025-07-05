package main

import (
	"fmt"
	"strings"
)

/* MakeUpperCase
https://www.codewars.com/kata/57a0556c7cb1f31ab3000ad7/train/go

Напишите функцию, которая преобразует входную строку в верхний регистр.
*/

func main() {
	fmt.Println(MakeUpperCase("hello")) // "HELLO"
}

// MakeUpperCase преобразует строку в верхний регистр.
// time: O(n), space: O(n), где n - длина строки
func MakeUpperCase(str string) string {
	return strings.ToUpper(str)
}
