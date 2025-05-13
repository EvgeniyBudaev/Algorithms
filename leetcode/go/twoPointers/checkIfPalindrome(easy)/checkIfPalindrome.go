package main

import "fmt"

/*
Учитывая строку s, верните true, если это палиндром, и false в противном случае.
Строка является палиндромом, если она читается как вперед, так и назад одинаково.
*/

func main() {
	fmt.Println(checkIfPalindrome("racecar")) // true
	fmt.Println(checkIfPalindrome("aleba"))   // false
}

// checkIfPalindrome проверяет, является ли строка палиндромом.
// time: O(n), space: O(1)
func checkIfPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}
