package main

import (
	"fmt"
)

/* Number of Decimal Digits
https://www.codewars.com/kata/58fa273ca6d84c158e000052/train/go

Определите общее количество цифр в целом числе (n>=0), переданном на вход функции. Например, 9 — однозначное число,
66 — двузначное, а 128685 — шестизначное. Будьте внимательны, чтобы избежать переполнений и потерь.

Все входные данные будут допустимыми.
*/

func main() {
	fmt.Println(Digits(0))                    // 1
	fmt.Println(Digits(9))                    // 1
	fmt.Println(Digits(66))                   // 2
	fmt.Println(Digits(18446744073709551615)) // 20
}

// Digits возвращает количество цифр в числе n.
// time: O(log n), space: O(1)
func Digits(n uint64) int {
	if n == 0 {
		return 1
	}

	count := 0
	for n > 0 {
		n /= 10
		count++
	}

	return count
}
