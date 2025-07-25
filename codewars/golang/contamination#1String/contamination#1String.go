package main

import (
	"fmt"
	"strings"
)

/* Contamination #1 -String-
https://www.codewars.com/kata/596fba44963025c878000039/train/go

ИИ заразил текст символом!!
Теперь этот текст полностью мутировал в этот символ.
Начиная с исходного текста и заданного символа, вернуть текст после мутации, чтобы все символы исходного текста были
заменены этим символом.
Если текст или символ пусты, вернуть пустую строку.
Никогда не будет случая, когда оба будут пустыми, поскольку ничего не происходит!!
Примечание: Символ — это строка длиной 1 или пустая строка.
*/

func main() {
	fmt.Println(Contamination("abc", "z")) // "zzz"
	fmt.Println(Contamination("", "z"))    // ""
}

// Contamination возвращает текст после мутации.
// time: O(n), space: O(n), где n - длина текста.
func Contamination(text, char string) string {
	return strings.Repeat(char, len(text))
}
