package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

/* Your order, please
https://www.codewars.com/kata/55c45be3b2079eccff00010f/train/go

Ваша задача — отсортировать заданную строку. Каждое слово в строке будет содержать одно число. Это число определяет
позицию слова в результатах.

Примечание: Числа могут быть от 1 до 9. Поэтому первым словом будет 1 (а не 0).

Если входная строка пуста, вернуть пустую строку. Слова во входной строке будут содержать только допустимые
последовательные числа.
*/

func main() {
	fmt.Println(Order("is2 Thi1s T4est 3a")) // "Thi1s is2 3a T4est"
}

// Order возвращает строку, отсортированную по порядку чисел в строке.
// time: O(n log n), space: O(n), n - длина строки
func Order(sentence string) string {
	if sentence == "" {
		return ""
	}

	words := strings.Fields(sentence) // [is2 Thi1s T4est 3a]
	sort.Slice(words, func(i, j int) bool {
		return extractNumber(words[i]) < extractNumber(words[j])
	})
	return strings.Join(words, " ")
}

// extractNumber извлекает число из слова.
// time: O(n), space: O(1), n - длина слова
func extractNumber(word string) int {
	for _, char := range word {
		if unicode.IsDigit(char) {
			return int(char - '0')
		}
	}
	return 0
}
