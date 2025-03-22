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
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	left, maxInt := 0, 0
	charMap := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		current := s[right]
		if prevIndex, exists := charMap[current]; exists {
			left = max(prevIndex+1, left)
		}
		maxInt = max(maxInt, right-left+1)
		charMap[current] = right
	}

	return maxInt
}
