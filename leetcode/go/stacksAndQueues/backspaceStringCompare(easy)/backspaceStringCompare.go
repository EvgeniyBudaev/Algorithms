package main

import "fmt"

/* 844. Backspace String Compare
https://leetcode.com/problems/backspace-string-compare/description/

Учитывая две строки s и t, верните true, если они равны, когда обе вводятся в пустые текстовые редакторы. '#' означает
символ возврата.
Обратите внимание, что после возврата пустого текста текст останется пустым.

Input: s = "ab#c", t = "ad#c"
Output: true
Пояснение: И s, и t становятся "ac".

Input: s = "ab##", t = "c#d#"
Output: true
Пояснение: И s, и t становятся "".

Input: s = "a#c", t = "b"
Output: false
Пояснение: s становится «c», а t становится «b».
*/

func main() {
	fmt.Println(backspaceCompare("ab#c", "ad#c")) // true
}

// backspaceCompare проверяет, равны ли две строки, когда обе вводятся в пустые текстовые редакторы, после символа возврата.
func backspaceCompare(s string, t string) bool {
	build := func(str string) string {
		var stack []rune
		for _, char := range str {
			if char != '#' {
				stack = append(stack, char)
			} else if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		}
		return string(stack) // ac ac for "ab#c", "ad#c"
	}

	return build(s) == build(t)
}
