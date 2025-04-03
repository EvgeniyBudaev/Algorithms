package main

import (
	"fmt"
	"strconv"
)

/* https://algo.monster/liteproblems/320

Учитывая строковое слово, верните список всех возможных обобщенных сокращений слова. Верните ответ в любом порядке.

Input: word = "word"
Output: ["4","3d","2r1","2rd","1o2","1o1d","1or1","1ord","w3","w2d","w1r1","w1rd","wo2","wo1d","wor1","word"]

Input: word = "a"
Output: ["1","a"]
*/

func main() {
	fmt.Println(generateAbbreviations("word")) // ["4","3d","2r1","2rd","1o2","1o1d","1or1","1ord","w3","w2d","w1r1","w1rd","wo2","wo1d","wor1","word"]
}

// generateAbbreviations генерирует все возможные сокращения для слова
func generateAbbreviations(word string) []string {
	var abbreviations []string
	var current []string

	var dfs func(string, []string)
	dfs = func(remaining string, currentAbbr []string) {
		// Если строка закончилась, добавляем текущее сокращение в результат
		if remaining == "" {
			abbreviations = append(abbreviations, joinStrings(currentAbbr))
			return
		}

		// Пробуем сокращения разной длины
		for i := 1; i <= len(remaining); i++ {
			// Добавляем числовое сокращение
			currentAbbr = append(currentAbbr, strconv.Itoa(i))

			// Если остались символы после сокращения
			if i < len(remaining) {
				// Добавляем следующий символ
				currentAbbr = append(currentAbbr, string(remaining[i]))
				// Продолжаем рекурсию для оставшейся части строки
				dfs(remaining[i+1:], currentAbbr)
				// Backtracking - удаляем последний символ
				currentAbbr = currentAbbr[:len(currentAbbr)-1]
			} else {
				// Если символы закончились, завершаем рекурсию
				dfs("", currentAbbr)
			}

			// Backtracking - удаляем числовое сокращение
			currentAbbr = currentAbbr[:len(currentAbbr)-1]
		}

		// Вариант без сокращения - оставляем первый символ как есть
		currentAbbr = append(currentAbbr, string(remaining[0]))
		dfs(remaining[1:], currentAbbr)
		currentAbbr = currentAbbr[:len(currentAbbr)-1]
	}

	dfs(word, current)
	return abbreviations
}

// joinStrings объединяет слайс строк в одну строку
func joinStrings(s []string) string {
	var result string
	for _, v := range s {
		result += v
	}
	return result
}
