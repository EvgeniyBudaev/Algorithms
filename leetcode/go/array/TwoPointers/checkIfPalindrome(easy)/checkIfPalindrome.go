package main

import "fmt"

/*
Учитывая строку s, верните true, если это палиндром, и false в противном случае.
Строка является палиндромом, если она читается как вперед, так и назад одинаково.
*/

func main() {
	r1 := checkIfPalindrome("racecar")
	fmt.Println(r1)
	r2 := checkIfPalindrome("aleba")
	fmt.Println(r2)
}

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
