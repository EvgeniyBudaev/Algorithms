package main

import (
	"fmt"
	"strings"
)

/* Fake Binary
https://www.codewars.com/kata/57eae65a4321032ce000002d/train/go

В данной строке цифр необходимо заменить любую цифру меньше 5 на «0», а любую цифру от 5 и выше — на «1».
Верните полученную строку.
*/

func main() {
	fmt.Println(FakeBin("45385593107843568")) // "01011110001100111"
}

// FakeBin заменяет цифры в строке на «0» или «1» в зависимости от их значения.
// time: O(n), space: O(n), n - длина строки
func FakeBin(x string) string {
	var result strings.Builder

	for _, char := range x {
		digit := int(char - '0')
		if digit < 5 {
			result.WriteByte('0')
		} else {
			result.WriteByte('1')
		}
	}

	return result.String()
}
