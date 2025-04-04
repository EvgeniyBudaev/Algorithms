package main

import "fmt"

/* https://leetcode.com/problems/palindromic-substrings/description/

Дана строка s, верните количество палиндромных подстрок в ней.
Строка является палиндромом, если она читается как в прямом, так и в обратном направлении.
Подстрока — это непрерывная последовательность символов внутри строки.

Input: s = "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".

Input: s = "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
*/

func main() {
	fmt.Println(countSubstrings("abc")) // 3 ("a", "b", "c")
}

// countSubstrings подсчитывает количество палиндромных подстрок в строке
func countSubstrings(s string) int {
	totalCount := 0
	for i := 0; i < len(s); i++ {
		// Подсчет палиндромов нечетной длины
		totalCount += extendPalindrome(s, i, i)
		// Подсчет палиндромов четной длины
		totalCount += extendPalindrome(s, i, i+1)
	}
	return totalCount
}

// extendPalindrome расширяется от центра и подсчитывает палиндромы
func extendPalindrome(s string, left, right int) int {
	count := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		count++
		left--
		right++
	}
	return count
}
