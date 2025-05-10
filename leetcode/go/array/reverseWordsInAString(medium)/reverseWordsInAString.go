package main

import (
	"fmt"
	"slices"
	"strings"
)

/* 151. Reverse Words in a String
https://leetcode.com/problems/reverse-words-in-a-string/description/

Дана входная строка s, обратный порядок слов.
Слово определяется как последовательность непробельных символов. Слова в s будут разделены как минимум одним пробелом.
Верните строку слов в обратном порядке, объединенных одним пробелом.

Обратите внимание, что s может содержать начальные или конечные пробелы или несколько пробелов между двумя словами.
Возвращаемая строка должна содержать только один пробел, разделяющий слова. Не включайте никаких дополнительных пробелов.

Input: s = "the sky is blue"
Output: "blue is sky the"

Input: s = "  hello world  "
Output: "world hello"
*/

func main() {
	fmt.Println(reverseWords("the sky is blue")) // "blue is sky the"
}

// reverseWords - функция переворачивает слова в строке.
// time: O(n), space: O(n)
func reverseWords(s string) string {
	words := strings.Fields(s)      // Разбиваем строку на слова
	slices.Reverse(words)           // Переворачиваем порядок слов
	return strings.Join(words, " ") // Объединяем слова в строку с разделителем пробел
}
