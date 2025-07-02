package main

import (
	"fmt"
	"unicode"
)

/* Alternate capitalization
https://www.codewars.com/kata/59cfc000aeb2844d16000075/train/go

Дана строка, сделайте заглавными буквы, которые занимают четные и нечетные индексы по отдельности, и верните,
как показано ниже. Индекс 0 будет считаться четным.

Например, capitalize("abcdef") = ['AbCdEf', 'aBcDeF'].

Входными данными будет строка в нижнем регистре без пробелов.
*/

func main() {
	fmt.Println(Capitalize("abcdef")) // ['AbCdEf', 'aBcDeF']
}

// Capitalize делает заглавными буквы, которые занимают четные и нечетные индексы по отдельности.
// time: O(n), space: O(n), где n - длина строки
func Capitalize(st string) []string {
	first := []rune(st)

	for i := 0; i < len(first); i++ {
		if i%2 == 0 {
			first[i] = unicode.ToUpper(first[i])
		}
	}

	second := []rune(st)

	for i := 0; i < len(second); i++ {
		if i%2 != 0 {
			second[i] = unicode.ToUpper(second[i])
		}
	}

	return []string{string(first), string(second)}
}
