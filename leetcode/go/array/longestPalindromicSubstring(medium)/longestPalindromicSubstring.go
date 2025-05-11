package main

import "fmt"

/* 5. Longest Palindromic Substring
https://leetcode.com/problems/longest-palindromic-substring/description/

Для заданной строки s вернуть самую длинную палиндромную подстроку в s.

Input: s = "babad"
Output: "bab"
Пояснение: «aba» также является допустимым ответом.
*/

func main() {
	fmt.Println(longestPalindrome("babad")) // "aba"
}

// longestPalindrome - возвращает самую длинную палиндромную подстроку в строке s.
// time: O(n^2), space: O(1)
func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		// Проверяем палиндромы нечетной длины
		len1 := expandAroundCenter(s, i, i)
		// Проверяем палиндромы четной длины
		len2 := expandAroundCenter(s, i, i+1)
		// Выбираем максимальную длину
		maxLen := max(len1, len2)

		// Если нашли палиндром длиннее текущего
		if maxLen > end-start {
			start = i - (maxLen-1)/2 // Вычитаем половину от длины палиндрома, чтобы получить начало палиндрома
			end = i + maxLen/2       // Прибавляем половину от длины палиндрома, чтобы получить конец палиндрома
		}
	}

	// Возвращаем палиндром
	return s[start : end+1]
}

// expandAroundCenter - возвращает длину палиндрома, который находится вокруг центра left и right.
// time: O(n), space: O(1)
func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] { // Проверяем, что символы с обоих сторон совпадают
		left--
		right++
	}
	return right - left - 1 // Возвращаем длину палиндрома
}
