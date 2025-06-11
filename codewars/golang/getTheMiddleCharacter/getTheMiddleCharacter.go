package main

import (
	"fmt"
)

/* Get the Middle Character
https://www.codewars.com/kata/56747fd5cb988479af000028/train/go

Вам будет дана непустая строка. Ваша задача — вернуть средний символ(ы) строки.
Если длина строки нечетная, верните средний символ.
Если длина строки четная, верните средние 2 символа.

Examples:
"test" --> "es"
"testing" --> "t"
"middle" --> "dd"
"A" --> "A"
*/

func main() {
	fmt.Println(GetMiddle("test"))    // "es"
	fmt.Println(GetMiddle("testing")) // "t"
}

// GetMiddle возвращает средний символ(ы) строки.
// time: O(1), space: O(1)
func GetMiddle(s string) string {
	str := []rune(s)     // Преобразуем строку в слайс символов
	half := len(str) / 2 // Половина длины строки

	if len(str)%2 == 0 {
		return string(str[half-1 : half+1])
	} else {
		return string(str[half])
	}
}
