package main

import (
	"fmt"
	"unicode"
)

/* https://leetcode.com/problems/check-if-the-sentence-is-pangram/description/

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

func checkIfPangram(sentence string) bool {
	if len(sentence) < 26 {
		return false
	}

	letters := make(map[rune]bool)
	for _, char := range sentence {
		if unicode.IsLetter(char) {
			letters[unicode.ToLower(char)] = true
		}
	}
	return len(letters) == 26
}
