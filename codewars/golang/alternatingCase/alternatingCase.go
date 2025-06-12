package main

import (
	"fmt"
	"unicode"
)

/* altERnaTIng cAsE <=> ALTerNAtiNG CaSe
https://www.codewars.com/kata/56efc695740d30f963000557/train/go

ALTERNATING CASE <=> ALTERNATING CASE
Определите String.prototype.toAlternatingCase (или аналогичную функцию/метод,
например to_alternating_case/toAlternatingCase/ToAlternatingCase на выбранном вами языке;
подробности см. в исходном решении) так, чтобы каждая строчная буква стала заглавной,
а каждая заглавная буква стала строчной. Например:

ToAlternatingCase("hello world"); // => "HELLO WORLD"
ToAlternatingCase("HELLO WORLD"); // => "hello world"
ToAlternatingCase("hello WORLD"); // => "HELLO world"
ToAlternatingCase("HeLLo WoRLD"); // => "hEllO wOrld"
ToAlternativeCase("12345"); // => "12345" (Non-alphabetical characters are unaffected)
ToAlternativeCase("1a2b3c4d5e"); // => "1A2B3C4D5E"
ToAlternativeCase("String.prototype.toAlternatingCase"); // => "sTRING.PROTOTYPE.TOaLTERNATINGcASE"
*/

func main() {
	fmt.Println(ToAlternatingCase("HeLLo WoRLD")) // "hEllO wOrld"
	fmt.Println(ToAlternatingCase("1a2b3c4d5e"))  // "1A2B3C4D5E"
}

// ToAlternatingCase возвращает строку, в которой каждая строчная буква стала заглавной, а каждая заглавная буква стала строчной.
// time: O(n), space: O(n), где n - длина строки
func ToAlternatingCase(str string) string {
	letters := []rune(str)

	for i := 0; i < len(letters); i++ {
		if unicode.IsLetter(letters[i]) {
			if unicode.IsUpper(letters[i]) {
				letters[i] = unicode.ToLower(letters[i])
			} else {
				letters[i] = unicode.ToUpper(letters[i])
			}
		}
	}

	return string(letters)
}
