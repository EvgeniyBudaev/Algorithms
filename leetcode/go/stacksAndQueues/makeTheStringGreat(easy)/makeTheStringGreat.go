package main

import (
	"fmt"
	"unicode"
)

/* 1544. Make The String Great
https://leetcode.com/problems/make-the-string-great/description/

Дана строка s, состоящая из строчных и прописных английских букв.
Хорошая строка — это строка, в которой нет двух соседних символов s[i] и s[i + 1], где:
0 <= i <= s.length - 2
s[i] — строчная буква, а s[i + 1] — та же буква, но в верхнем регистре или наоборот.
Чтобы сделать строку хорошей, вы можете выбрать два соседних символа, которые делают строку плохой, и удалить их.
Вы можете продолжать делать это, пока строка не станет хорошей.
Верните строку после того, как исправили ее. Ответ гарантированно будет уникальным при заданных ограничениях.
Обратите внимание, что пустая строка тоже подойдет.

Input: s = "leEeetcode"
Output: "leetcode"
Explanation: На первом этапе, либо вы выберете i = 1, либо i = 2, в результате оба варианта «leEeetcode» будут сокращены
до «leetcode».

Input: s = "abBAcC"
Output: ""
Explanation: У нас есть много возможных сценариев, и все они ведут к одному и тому же ответу. Например:
"abBAcC" --> "aAcC" --> "cC" --> ""
"abBAcC" --> "abBA" --> "aA" --> ""

Input: s = "s"
Output: "s"
*/

func main() {
	fmt.Println(makeGood("leEeetcode")) // "leetcode"
}

// makeGood возвращает хорошую строку.
func makeGood(s string) string {
	if len(s) == 0 {
		return ""
	}

	var stack []rune // Используем слайс для стека символов

	for _, char := range s {
		if len(stack) > 0 {
			lastChar := stack[len(stack)-1]
			// Если символы являются противоположными регистрами
			if char != lastChar && unicode.ToLower(char) == unicode.ToLower(lastChar) {
				stack = stack[:len(stack)-1] // Удаляем последний элемент
				continue
			}
		}
		stack = append(stack, char)
	}

	return string(stack)
}
