package main

import (
	"fmt"
	"strings"
)

/* Shortest Word
https://www.codewars.com/kata/57cebe1dc6fdc20c57000ac9/train/go

Проще говоря, если задана строка слов, верните длину самого короткого из них.
Строка никогда не будет пустой, и вам не нужно учитывать различные типы данных.
*/

func main() {
	fmt.Println(FindShort("bitcoin take over the world maybe who knows perhaps")) // 3
}

// FindShort принимает строку и возвращает длину самого короткого слова.
// time: O(n), space: O(1), где n - длина строки
func FindShort(s string) int {
	words := strings.Fields(s)
	minLen := len(words[0])

	for _, word := range words {
		if len(word) < minLen {
			minLen = len(word)
		}
	}

	return minLen
}
