package main

import "fmt"

/* 242. Valid Anagram
https://leetcode.com/problems/valid-anagram/description/

Учитывая две строки s и t, верните true, если t является анаграммой s, и false в противном случае.
Анаграмма — это слово или фраза, образованная путем перестановки букв другого слова или фразы, обычно с использованием
всех исходных букв ровно один раз.

Input: s = "anagram", t = "nagaram"
Output: true

Input: s = "rat", t = "car"
Output: false
*/

func main() {
	fmt.Println(isAnagram("anagram", "nagaram")) // true
}

// isAnagram проверяет, являются ли строки анаграммами.
// time: O(n), space: O(n)
func isAnagram(s string, t string) bool {
	// Если длины строк разные, они не могут быть анаграммами
	if len(s) != len(t) {
		return false
	}

	// Создаем map для подсчета символов
	charCount := make(map[rune]int)

	// Подсчитываем количество каждого символа в первой строке
	for _, char := range s {
		charCount[char]++
	}

	// Проверяем вторую строку
	for _, char := range t {
		// Если символа нет в первой строке или его недостаточно
		if count, exists := charCount[char]; !exists || count <= 0 {
			return false
		}
		// Уменьшаем счетчик символа
		charCount[char]--
	}

	return true // Все символы в обеих строках присутствуют одинаковое количество раз
}
