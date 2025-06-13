package main

import (
	"fmt"
	"unicode"
)

/* Is it a palindrome?
https://www.codewars.com/kata/57a1fd2ce298a731b20006a4/train/go

Напишите функцию, которая проверяет, является ли заданная строка (без учета регистра) палиндромом.

Палиндром — это слово, число, фраза или другая последовательность символов, которая читается одинаково как в прямом,
так и в обратном порядке, например, «madam» или «racecar».
*/

func main() {
	fmt.Println(IsPalindrome("racecar")) // true
	fmt.Println(IsPalindrome("a"))       // true
	fmt.Println(IsPalindrome("Abba"))    // true
}

// IsPalindrome проверяет, является ли строка палиндромом.
// time: O(n), space: O(1)
func IsPalindrome(str string) bool {
	left, right := 0, len(str)-1

	for left < right {
		if unicode.ToLower(rune(str[left])) != unicode.ToLower(rune(str[right])) {
			return false
		}
		left++
		right--
	}

	return true
}
