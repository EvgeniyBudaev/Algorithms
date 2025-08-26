package main

import (
	"fmt"
	"strings"
)

/* Band name generator
https://www.codewars.com/kata/59727ff285281a44e3000011/train/go

Моя подруга хочет придумать новое название для своей группы. Ей нравятся группы, использующие формулу:
«The» + существительное с заглавной первой буквой, например:

"dolphin" -> "The Dolphin"

Однако, когда существительное НАЧИНАЕТСЯ и ЗАКАНЧИВАЕТСЯ на одну и ту же букву, она предпочитает повторять его дважды и
соединять их первой и последней буквами, образуя одно слово (БЕЗ «The» в начале), например:

"alaska" -> "Alaskalaska"

Завершите функцию, которая принимает существительное как строку и возвращает предпочитаемое ею название группы в виде
строки.
*/

func main() {
	fmt.Println(bandNameGenerator("dolphin"))     // "The Dolphin"
	fmt.Println(bandNameGenerator("alaska"))      // "Alaskalaska"
	fmt.Println(bandNameGenerator("tart"))        // "Tartart"
	fmt.Println(bandNameGenerator("step-father")) // "The Step-Father"
}

// bandNameGenerator возвращает предпочитаемое название группы.
// time: O(n), space: O(n), n - длина слова
func bandNameGenerator(word string) string {
	if len(word) > 0 && word[0] == word[len(word)-1] {
		return strings.Title(word) + word[1:]
	}
	return "The " + strings.Title(word)
}
