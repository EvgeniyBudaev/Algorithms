package main

import (
	"fmt"
	"sort"
	"strings"
)

/* Where is my parent!?(cry)
https://www.codewars.com/kata/58539230879867a8cd00011c/train/go

Мамы устроили танцевальную вечеринку для детей в школе. На вечеринке присутствовали только мамы и их дети.
Все веселились на танцполе, как вдруг погас свет. Темная ночь, и никто не видел друг друга. Но вы летали рядом,
вы можете видеть в темноте и можете телепортировать людей куда угодно.

Условные обозначения:
- Заглавные буквы обозначают матерей, строчные — их детей, например, дети матери «А» — «aaaa».
- Входные данные функции: Строка содержит только буквы, заглавные буквы уникальны.

Задание:
Расположите всех людей в алфавитном порядке, где за матерями следуют их дети, например,
«aAbaBb» => «AaaBbb».
*/

func main() {
	fmt.Println(FindChildren("aAbaBb")) // "AaaBbb"
}

// FindChildren возвращает строку, отсортированную по алфавиту, где за матерями следуют их дети.
// time: O(n*log(n)), space: O(n), n - длина строки
func FindChildren(dancingBrigade string) string {
	strLower := strings.ToLower(dancingBrigade) // aababb
	strLw := strings.Split(strLower, "")        // [a a b a b b]
	sort.Strings(strLw)                         // [a a a b b b]
	ch := ""                                    // ch - это предыдущий символ

	for i, char := range strLw { // [a a a b b b]
		if char != ch {
			temp := char
			strLw[i] = strings.Title(char)
			ch = temp
		}
	}

	return strings.Join(strLw, "")
}
