package main

import (
	"fmt"
	"unicode"
)

/* https://leetcode.com/problems/valid-palindrome/description/

Фраза является палиндромом, если после преобразования всех прописных букв в строчные и удаления всех небуквенно-цифровых
символов она читается одинаково и вперед, и назад. Буквенно-цифровые символы включают буквы и цифры.
Учитывая строку s, верните true, если это палиндром, или false в противном случае.

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
*/

func main() {
	s := "A man, a plan, a canal: Panama"
	r := isPalindrome(s)
	fmt.Println(r)
}

func isAlphaNumeric(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		skipLeft := !isAlphaNumeric(rune(s[left]))
		if skipLeft {
			left++
			continue
		}

		skipRight := !isAlphaNumeric(rune(s[right]))
		if skipRight {
			right--
			continue
		}

		endsEqual := unicode.ToLower(rune(s[left])) == unicode.ToLower(rune(s[right]))
		if !endsEqual {
			return false
		}

		left++
		right--
	}

	return true
}
