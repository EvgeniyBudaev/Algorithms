package main

import (
	"fmt"
	"strings"
)

/* Reverse words

Examples:
"This is an example!" ==> "sihT si na !elpmaxe"
"double  spaces"      ==> "elbuod  secaps"
*/

func main() {
	fmt.Println(ReverseWords("This is an example!")) // "sihT si na !elpmaxe"
}

// ReverseWords переворачивает слова в строке.
// time: O(n), space: O(n)
func ReverseWords(str string) string {
	words := strings.Split(str, " ")

	for i, word := range words {
		words[i] = reverseString(word)
	}

	return strings.Join(words, " ")
}

// reverseString переворачивает строку.
// time: O(n), space: O(n)
func reverseString(s string) string {
	runes := []rune(s)
	left, right := 0, len(runes)-1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}
