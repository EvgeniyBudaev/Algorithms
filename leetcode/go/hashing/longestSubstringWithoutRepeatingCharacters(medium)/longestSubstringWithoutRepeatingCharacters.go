package main

import "fmt"

/* 3. Longest Substring Without Repeating Characters
https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

Дана строка s, найдите длину самой длинной подстроки без повторения символов.

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
}

// lengthOfLongestSubstring - находит длину самой длинной подстрока без повторения символов.
// time: O(n), space: O(n)
func lengthOfLongestSubstring(s string) int {
	// Если строка пустая или содержит только один символ, возвращаем длину строки
	if len(s) <= 1 {
		return len(s)
	}

	left, right, maxLength := 0, 0, 0 // Индексы левого и правого края подстроки, максимальная длина
	charSet := make(map[byte]bool)    // Мапа для хранения уникальных символов

	for right < len(s) { // Проходим по строке
		currentChar := s[right]                 // Текущий символ
		if _, ok := charSet[currentChar]; !ok { // Если символ не найден в мапе
			charSet[currentChar] = true              // Добавляем его в мапу
			maxLength = max(maxLength, right-left+1) // Обновляем максимальную длину
			right++                                  // Переходим к следующему символу
		} else { // Если символ найден в мапе
			delete(charSet, s[left]) // Удаляем символ из мапы
			left++                   // Переходим к следующему символу
		}
	}

	return maxLength // Возвращаем максимальную длину
}
