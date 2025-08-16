package main

import (
	"fmt"
	"strings"
)

/* Jaden Casing Strings
https://www.codewars.com/kata/5390bac347d09b7da40006f6/train/go

Джейден Смит, сын Уилла Смита, звезда таких фильмов, как «Каратэ-пацан» (2010) и «После нашей эры» (2013).
Джейден также известен своей философией, которую он распространяет в Твиттере. В Твиттере он практически всегда пишет
каждое слово с заглавной буквы. Для простоты вам придётся писать каждое слово с заглавной буквы; посмотрите, как должны
писаться сокращения, в примере ниже.

Ваша задача — преобразовать строки в то, как их написал бы Джейден Смит. Эти строки — настоящие цитаты Джейдена Смита,
но они написаны не с заглавной буквы, как он написал изначально.

Пример:
Not Jaden-Cased: "How can mirrors be real if our eyes aren't real"
Jaden-Cased:     "How Can Mirrors Be Real If Our Eyes Aren't Real"
*/

func main() {
	fmt.Println(ToJadenCase("How can mirrors be real if our eyes aren't real")) // "How Can Mirrors Be Real If Our Eyes Aren't Real"
}

// ToJadenCase преобразует строку в Jaden Case.
// time: O(n), space: O(n), n - длина строки
func ToJadenCase(str string) string {
	return strings.Title(str)
}
