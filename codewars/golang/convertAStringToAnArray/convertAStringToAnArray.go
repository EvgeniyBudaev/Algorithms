package main

import (
	"fmt"
	"strings"
)

/* Convert a string to an array
https://www.codewars.com/kata/57e76bc428d6fbc2d500036d/train/go

Напишите функцию для разбиения строки и преобразования ее в массив слов.

Examples (Input ==> Output):
"Robin Singh" ==> ["Robin", "Singh"]
*/

func main() {
	fmt.Println(StringToArray("Robin Singh")) // ["Robin", "Singh"]
}

// StringToArray разбивает строку на массив слов.
// time: O(n), space: O(n), где n - длина строки.
func StringToArray(str string) []string {
	return strings.Split(str, " ") // or return strings.Fields(str)
}
