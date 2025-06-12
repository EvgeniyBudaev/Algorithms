package main

import (
	"fmt"
	"strings"
)

/* Remove String Spaces
https://www.codewars.com/kata/57eae20f5500ad98e50002c5/train/go

Напишите функцию, которая удаляет пробелы из строки, а затем возвращает полученный результат.

Examples:
"8 j 8   mBliB8g  imjB8B8  jl  B" -> "8j8mBliB8gimjB8B8jlB"
"8 8 Bi fk8h B 8 BB8B B B  B888 c hl8 BhB fd" -> "88Bifk8hB8BB8BBBB888chl8BhBfd"
"8aaaaa dddd r     " -> "8aaaaaddddr"
*/

func main() {
	fmt.Println(NoSpace("8 j 8   mBliB8g  imjB8B8  jl  B")) // "88Bifk8hB8BB8BBBB888chl8BhBfd"
}

func NoSpace(word string) string {
	return strings.ReplaceAll(word, " ", "")
}
