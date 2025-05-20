package main

import "fmt"

/* 3. Longest Substring Without Repeating Characters
https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

Дана строка s. Найдите длину самой длинной подстрока без повторения символов.

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

// lengthOfLongestSubstring находит длину самой длинной подстрока без повторения символов.
// time: O(n), space: O(n)
func lengthOfLongestSubstring(s string) int {
	// Если строка пустая или состоит из одного символа, то длина самой длинной подстроки равна длине строки.
	if len(s) <= 1 {
		return len(s)
	}

	left, right, maxLength := 0, 0, 0 // left - левая граница, right - правая граница, maxLength - максимальная длина подстроки
	charMap := make(map[byte]bool)    // charMap - карта для хранения уникальных символов

	for right < len(s) {
		currentChar := s[right]                 // текущий символ
		if _, ok := charMap[currentChar]; !ok { // если символ не найден в charMap, то добавляем его в charMap
			charMap[currentChar] = true              // добавляем символ в charMap
			maxLength = max(maxLength, right-left+1) // вычисляем максимальную длину подстроки
			right++                                  // сдвигаем правую границу
		} else { // если символ найден в charMap, то удаляем символ с левой границы и сдвигаем левую границу
			delete(charMap, s[left])
			left++
		}
	}

	return maxLength // возвращаем максимальную длину подстроки
}
