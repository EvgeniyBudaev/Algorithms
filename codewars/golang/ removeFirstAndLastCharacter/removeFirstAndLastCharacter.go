package main

import (
	"fmt"
)

/* Remove First and Last Character
https://www.codewars.com/kata/56bc28ad5bdaeb48760009b0/train/go

Это довольно просто. Ваша цель — создать функцию, которая удаляет первый и последний символы строки.
Вам дан один параметр — исходная строка. Вам не нужно беспокоиться о строках, содержащих менее двух символов.
*/

func main() {
	fmt.Println(RemoveChar("country")) // "ountr"
}

// RemoveChar удаляет первый и последний символы строки.
// time: O(n), space: O(n)
func RemoveChar(word string) string {
	return word[1 : len(word)-1]
}
