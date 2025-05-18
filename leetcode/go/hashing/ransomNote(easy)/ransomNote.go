package main

import "fmt"

/* 383. Ransom Note
https://leetcode.com/problems/ransom-note/description/

Учитывая две строки ransomNote и magazine, верните true, если ransomNote можно создать с использованием букв из magazine,
и false в противном случае.
Каждое письмо в magazine можно использовать в ransomNote только один раз.

Input: ransomNote = "a", magazine = "b"
Output: false

Input: ransomNote = "aa", magazine = "ab"
Output: false

Input: ransomNote = "aa", magazine = "aab"
Output: true
*/

func main() {
	fmt.Println(canConstruct("a", "b"))    // false
	fmt.Println(canConstruct("aa", "aab")) // true
}

// canConstruct проверяет, можно ли построить строку ransomNote из букв строки magazine.
// time: O(n), space: O(1)
func canConstruct(ransomNote string, magazine string) bool {
	// Создаем map для подсчета количества каждой буквы в magazine
	letterCounts := make(map[rune]int)

	// Заполняем map: ключ - буква, значение - количество таких букв в magazine
	for _, letter := range magazine {
		letterCounts[letter]++
	}

	// Проверяем каждую букву в ransomNote
	for _, letter := range ransomNote {
		// Если буквы нет в magazine или ее недостаточно, возвращаем false
		if count, exists := letterCounts[letter]; !exists || count == 0 {
			return false
		}
		// Уменьшаем количество оставшихся букв
		letterCounts[letter]--
	}

	// Если все буквы ransomNote есть в magazine в достаточном количестве
	return true
}
