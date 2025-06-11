package main

import (
	"fmt"
)

/* Century From Year
https://www.codewars.com/kata/5a3fe3dde1ce0e8ed6000097/train/go

Введение.
Первый век охватывает период с 1 по 100 год включительно, второй век — с 101 по 200 год включительно и т. д.

Задание.
Дан год, верните век, в котором он находится.

Examples:
1705 --> 18
1900 --> 19
1601 --> 17
2000 --> 20
2742 --> 28
*/

func main() {
	fmt.Println(century(1705)) // 18
	fmt.Println(century(2000)) // 20
	fmt.Println(century(89))   // 1
}

// century возвращает век, в котором находится год.
// time: O(1), space: O(1)
func century(year int) int {
	if year%100 == 0 {
		return year / 100
	}

	return (year / 100) + 1
}
