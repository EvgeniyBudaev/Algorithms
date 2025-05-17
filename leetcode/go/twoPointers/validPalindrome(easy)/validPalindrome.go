package main

import (
	"fmt"
	"unicode"
)

/* 125. Valid Palindrome
https://leetcode.com/problems/valid-palindrome/description/

Фраза является палиндромом, если после преобразования всех прописных букв в строчные и удаления всех небуквенно-цифровых
символов она читается одинаково и вперед, и назад. Буквенно-цифровые символы включают буквы и цифры.
Учитывая строку s, верните true, если это палиндром, или false в противном случае.

Input: s = "A man, a plan, a canal: Panama"
Output: true
Объяснение: "amanaplanacanalpanama" палиндром.

Input: s = "race a car"
Output: false
Объяснение: "raceacar" не палиндром.

Input: s = " "
Output: true
Объяснение: s — это пустая строка "" после удаления небуквенно-цифровых символов.
Поскольку пустая строка читается одинаково и вперед, и назад, она является палиндромом.
*/

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
}

// isAlphaNumeric проверяет, является ли переданный символ буквенно-цифровым
// time: O(1), space: O(1)
func isAlphaNumeric(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

// isPalindrome проверяет, является ли строка s палиндромом после нормализации.
// time: O(n), где n - количество символов в строке, space: O(1)
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		skipLeft := !isAlphaNumeric(rune(s[left])) // пропускаем символы, которые не являются буквенно-цифровыми
		if skipLeft {
			left++
			continue
		}

		skipRight := !isAlphaNumeric(rune(s[right])) // пропускаем символы, которые не являются буквенно-цифровыми
		if skipRight {
			right--
			continue
		}

		// сравниваем символы, которые являются буквенно-цифровыми
		endsEqual := unicode.ToLower(rune(s[left])) == unicode.ToLower(rune(s[right]))
		if !endsEqual {
			return false
		}

		left++
		right--
	}

	return true // строка является палиндромом
}
