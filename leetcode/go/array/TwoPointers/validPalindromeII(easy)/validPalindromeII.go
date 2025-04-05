package main

import "fmt"

/* https://leetcode.com/problems/valid-palindrome-ii/description/

Учитывая строку s, верните true, если s может быть палиндромом после удаления из нее не более одного символа.

Input: s = "aba"
Output: true

Input: s = "abca"
Output: true
Объяснение: Вы можете удалить символ 'c'.

Input: s = "abc"
Output: false
*/

func main() {
	fmt.Println(validPalindrome("aba"))  // true
	fmt.Println(validPalindrome("abca")) // true
	fmt.Println(validPalindrome("abc"))  // false
}

// validPalindrome проверяет, можно ли сделать строку палиндромом, удалив один символ.
func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
		}
		left++
		right--
	}
	return true
}

// isPalindrome проверяет, является ли строка палиндромом.
func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
