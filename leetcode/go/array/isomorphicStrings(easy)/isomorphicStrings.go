package main

import (
	"fmt"
)

/* https://leetcode.com/problems/isomorphic-strings/description/

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
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charMap := make(map[byte]byte)

	for i := range s {
		if v, ok := charMap[s[i]]; !ok {
			charMap[s[i]] = t[i]
		} else {
			if v != t[i] {
				return false
			}
		}
	}

	return true
}
