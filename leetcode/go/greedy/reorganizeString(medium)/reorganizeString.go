package main

import (
	"fmt"
	"sort"
)

/* 767. Reorganize String
https://leetcode.com/problems/reorganize-string/description/
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

// reorganizeString переставляет символы в строке так, чтобы любые два соседних символа не были одинаковыми.
// time: O(n), space: O(n)
func reorganizeString(s string) string {
	// Создаем карту для подсчета частот символов
	freq := make(map[rune]int)
	for _, char := range s {
		freq[char]++
	}

	// Создаем слайс для сортировки символов по частоте
	type charCount struct {
		char  rune // Символ
		count int  // Частота
	}
	var chars []charCount           // Создаем слайс для хранения символов и их частот
	for char, count := range freq { // Добавляем символы и их частоты в слайс
		chars = append(chars, charCount{char, count}) // Добавляем символ и его частоту
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
	position := 0 // Индекс для заполнения результата

	for _, cc := range chars { // Делаем перебор для каждого символа в отсортированном слайсе
		for i := 0; i < cc.count; i++ { // Для каждого символа заполняем его в результирующей строке
			result[position] = cc.char // Заполняем символ в результирующей строке
			position += 2              // Переходим к следующему индексу для заполнения
			// Если мы достигли конца результирующей строки, начинаем с начала
			if position >= len(result) {
				position = 1
			}
		}
	}

	// Возвращаем реорганизованную строку
	return string(result)
}
