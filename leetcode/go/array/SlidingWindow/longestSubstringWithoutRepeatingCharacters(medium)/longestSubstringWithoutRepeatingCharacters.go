package main

import "fmt"

/* https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

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
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	left, right, maxLength := 0, 0, 0
	charMap := make(map[byte]bool)

	for right < len(s) {
		currentChar := s[right]
		if _, ok := charMap[currentChar]; !ok {
			charMap[currentChar] = true
			maxLength = max(maxLength, right-left+1)
			right++
		} else {
			delete(charMap, s[left])
			left++
		}
	}

	return maxLength
}
