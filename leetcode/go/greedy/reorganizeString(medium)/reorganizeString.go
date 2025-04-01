package main

import (
	"fmt"
	"sort"
)

/* https://leetcode.com/problems/reorganize-string/description/
solution https://leetcode.com/problems/reorganize-string/submissions/1333724578/

Дана строка s, переставьте символы s так, чтобы любые два соседних символа не были одинаковыми.
Верните любую возможную перестановку s или верните "", если это невозможно.

Input: s = "aab"
Output: "aba"

Input: s = "aaab"
Output: ""
*/

func main() {
	fmt.Println(reorganizeString("aab")) // "aba"
}

func reorganizeString(s string) string {
	// Создаем карту для подсчета частот символов
	freq := make(map[rune]int)
	for _, char := range s {
		freq[char]++
	}

	// Создаем слайс для сортировки символов по частоте
	type charCount struct {
		char  rune
		count int
	}
	var chars []charCount
	for char, count := range freq {
		chars = append(chars, charCount{char, count})
	}

	// Сортируем по убыванию частоты
	sort.Slice(chars, func(i, j int) bool {
		return chars[i].count > chars[j].count
	})

	// Проверяем возможность реорганизации
	if chars[0].count > (len(s)+1)/2 {
		return ""
	}

	// Создаем результат и заполняем его
	result := make([]rune, len(s))
	position := 0

	for _, cc := range chars {
		for i := 0; i < cc.count; i++ {
			result[position] = cc.char
			position += 2
			if position >= len(result) {
				position = 1
			}
		}
	}

	return string(result)
}
