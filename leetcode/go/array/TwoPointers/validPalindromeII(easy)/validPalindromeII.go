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
	s1 := "aba"
	r1 := validPalindrome(s1)
	fmt.Println(r1)
	s2 := "abca"
	r2 := validPalindrome(s2)
	fmt.Println(r2)
	s3 := "abc"
	r3 := validPalindrome(s3)
	fmt.Println(r3)
}

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
