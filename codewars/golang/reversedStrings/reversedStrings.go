package main

import "fmt"

/* Reversed Strings
https://www.codewars.com/kata/5168bb5dfe9a00b126000018/train/go

Дополните решение так, чтобы оно перевернуло переданную в него строку.

Example:
'world'  =>  'dlrow'
'word'   =>  'drow'
*/

func main() {
	fmt.Println(Solution("world")) // "dlrow"
}

// Solution возвращает перевернутую строку.
// time: O(n), space: O(1)
func Solution(word string) string {
	str := []byte(word)
	left, right := 0, len(word)-1

	for left < right {
		str[left], str[right] = str[right], str[left]
		left++
		right--
	}

	return string(str)
}
