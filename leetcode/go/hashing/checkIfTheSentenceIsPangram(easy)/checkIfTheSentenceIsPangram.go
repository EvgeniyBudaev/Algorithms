package main

import (
	"fmt"
	"unicode"
)

/* 1832. Check if the Sentence Is Pangram
https://leetcode.com/problems/check-if-the-sentence-is-pangram/description/

Панграмма – это предложение, в котором каждая буква английского алфавита встречается хотя бы один раз.
Учитывая строковое предложение, содержащее только строчные буквы английского языка, верните true, если предложение
является панграммой, или false в противном случае.

Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
Output: true

Input: sentence = "leetcode"
Output: false
*/

func main() {
	fmt.Println(checkIfPangram("thequickbrownfoxjumpsoverthelazydog")) // true
	fmt.Println(checkIfPangram("leetcode"))                            // false
}

// checkIfPangram - проверяет, является ли предложение панграммой.
// time: O(n), space: O(1)
func checkIfPangram(sentence string) bool {
	// Проверяем, что предложение содержит хотя бы 26 уникальных букв
	if len(sentence) < 26 {
		return false
	}

	letters := make(map[rune]bool) // Создаем карту для хранения уникальных букв
	for _, char := range sentence {
		if unicode.IsLetter(char) { // Если символ является буквой
			letters[unicode.ToLower(char)] = true
		}
	}

	return len(letters) == 26 // Если количество уникальных букв равно 26, то это панграмма
}
