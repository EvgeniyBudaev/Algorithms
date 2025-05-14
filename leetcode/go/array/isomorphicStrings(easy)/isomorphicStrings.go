package main

import (
	"fmt"
)

/* 205. Isomorphic Strings
https://leetcode.com/problems/isomorphic-strings/description/

Даны две строки s и t. Определите, изоморфны ли они.
Две строки s и t изоморфны, если символы в s можно заменить, чтобы получить t.
Все вхождения символа должны быть заменены другим символом с сохранением порядка символов. Никакие два символа не могут
сопоставляться одному и тому же символу, но символ может сопоставляться сам с собой.
Например, рассмотрим строки ACAB и XCXY.
Они изоморфны, поскольку мы можем отобразить 'A' —> 'X', 'B' —> 'Y'и 'C' —> 'C'.

Input: s = "egg", t = "add"
Output: true

Input: s = "foo", t = "bar"
Output: false

Input: s = "paper", t = "title"
Output: true
*/

func main() {
	fmt.Println(isIsomorphic("egg", "add")) // true
	fmt.Println(isIsomorphic("abc", "aaa")) // false
}

// isIsomorphic - определяет, изоморфны ли строки.
// time: O(n), space: O(n), где n - длина строки.
func isIsomorphic(s string, t string) bool {
	// Если длины строк не совпадают, то они не изоморфны.
	if len(s) != len(t) {
		return false
	}

	// Используем карту для отображения символов.
	charMap := make(map[byte]byte)
	// Используем карту для отслеживания уже сопоставленных символов.
	valuesMap := make(map[byte]bool)

	for i := range s { // Проходим по символам в строке s.
		sLetter := s[i] // Текущий символ в строке s.
		tLetter := t[i] // Текущий символ в строке t.

		v, ok := charMap[sLetter]
		if !ok {
			// Если символ уже существует, но не соответствует текущему символу строки t, возвращаем false.
			if _, saved := valuesMap[tLetter]; saved {
				return false
			}

			// Символ не существует, добавляем его в карту.
			valuesMap[tLetter] = true  // Добавляем в карту, чтобы отслеживать уже сопоставленные символы.
			charMap[sLetter] = tLetter // Добавляем в карту соответствия символов.
			continue                   // Переходим к следующему символу.
		}

		// Если символ уже существует, но не соответствует текущему символу строки t, возвращаем false.
		if v != tLetter {
			return false
		}
	}

	return true // Строки изоморфны.
}
