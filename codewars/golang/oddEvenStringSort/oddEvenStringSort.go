package main

import (
	"fmt"
)

/* Odd-Even String Sort
https://www.codewars.com/kata/580755730b5a77650500010c/train/go

Дана строка s. Ваша задача — вернуть другую строку, в которой символы s с чётными и нечётными индексами сгруппированы,
а группы разделены пробелами. Чётная группа идёт первой, за ней следует пробел, а затем нечётная часть.

Примеры
input: "CodeWars" => "CdWr oeas"
|||||||| |||| ||||
indexs: 01234567 0246 1357
Чётные индексы — 0, 2, 4, 6, поэтому первая группа — "CdWr".
Нечётные индексы — 1, 3, 5, 7, поэтому вторая группа — "oeas".
И последняя возвращаемая строка — "Cdwr oeas".
*/

func main() {
	fmt.Println(SortMyString("CodeWars")) // "CdWr oeas"
}

// SortMyString сортирует строку по чётным и нечётным индексам.
// time: O(n), space: O(n), n - длина строки
func SortMyString(s string) string {
	left, right := "", ""

	for i, val := range s {
		if i%2 == 0 {
			left += string(val)
		} else {
			right += string(val)
		}
	}

	return left + " " + right
}
