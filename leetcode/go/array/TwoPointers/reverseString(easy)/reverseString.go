package main

import "fmt"

/* https://leetcode.com/problems/reverse-string/description/

Напишите функцию, которая переворачивает строку. Входная строка задается как массив символов s.

Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]

Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]
*/

func main() {
	s1 := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s1)
	fmt.Println(string(s1)) // ["o","l","l","e","h"]
}

func reverseString(s []byte) {
	left, right := 0, len(s)-1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
