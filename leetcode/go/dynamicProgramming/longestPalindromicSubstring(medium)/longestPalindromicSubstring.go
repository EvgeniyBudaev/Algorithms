package main

import (
	"fmt"
	"math"
)

/* https://leetcode.com/problems/longest-palindromic-substring/description/

Учитывая строку s, вернуть самую длинную палиндромный подстрока в s.

Input: s = "babad"
Output: "bab"
Пояснение: «aba» также является допустимым ответом.

Input: s = "cbbd"
Output: "bb"
*/

func main() {
	fmt.Println(longestPalindrome("babad")) // "bab" или "aba"
}

// longestPalindrome находит наибольшую палиндромную подстроку в строке
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	// Функция для расширения вокруг центра
	expandAroundCenter := func(s string, left, right int) int {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		return right - left - 1
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		// Проверяем палиндромы нечетной длины (центр в i)
		odd := expandAroundCenter(s, i, i)
		// Проверяем палиндромы четной длины (центр между i и i+1)
		even := expandAroundCenter(s, i, i+1)
		// Выбираем максимальную длину
		maxLen := int(math.Max(float64(odd), float64(even)))

		// Если нашли палиндром больше текущего, обновляем границы
		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	// Возвращаем подстроку от start до end+1
	return s[start : end+1]
}
