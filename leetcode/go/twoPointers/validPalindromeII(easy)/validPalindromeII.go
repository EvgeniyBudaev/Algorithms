package main

import "fmt"

/* 680. Valid Palindrome II
https://leetcode.com/problems/valid-palindrome-ii/description/

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
// time: O(n), space: O(1)
func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right { // Сравниваем символы с обоих концов строки.
		// Если символы не равны, проверяем, можно ли сделать строку палиндромом, удалив один символ.
		if s[left] != s[right] {
			return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
		}
		left++
		right--
	}
	return true // Если цикл завершился успешно, строка является палиндромом.
}

// isPalindrome проверяет, является ли строка палиндромом.
// time: O(n), space: O(1)
func isPalindrome(s string, left, right int) bool {
	for left < right {
		// Если символы не равны, строка не является палиндромом.
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true // Если цикл завершился успешно, строка является палиндромом.
}
