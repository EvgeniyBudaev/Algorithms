package main

import (
	"fmt"
	"strings"
)

/* 1249. Minimum Remove to Make Valid Parentheses
https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/description/

Дана строка s из '(' , ')' и строчных английских символов.

Ваша задача — удалить минимальное количество скобок ('(' или ')', в любых позициях ) так, чтобы результирующая строка
скобок была допустимой и возвращала любую допустимую строку.

Формально строка скобок допустима тогда и только тогда, когда:

Это пустая строка, содержит только строчные символы или
Она может быть записана как AB (A, объединенная с B), где A и B — допустимые строки, или
Она может быть записана как (A), где A — допустимая строка.

Input: s = "lee(t(c)o)de)"
Output: "lee(t(c)o)de"
Пояснение: «lee(t(co)de)», «lee(t(c)ode)» также будут приняты.

Input: s = "a)b(c)d"
Output: "ab(c)d"
*/

func main() {
	s := "lee(t(c)o)de)"
	fmt.Println(minRemoveToMakeValid(s)) // "lee(t(c)o)de"
}

// minRemoveToMakeValid- удалить минимальное количество скобок ('(' или ')', в любых позициях ) так,
// чтобы результирующая строка скобок была допустимой.
// time: O(n), space: O(n)
func minRemoveToMakeValid(s string) string {
	runes := []rune(s)      // Преобразуем строку в слайс рун для изменения
	stack := make([]int, 0) // храним индексы открывающих скобок

	for i, char := range runes {
		if char == '(' {
			stack = append(stack, i) // запоминаем позицию '('
		} else if char == ')' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1] // нашли пару, удаляем из стека
			} else {
				runes[i] = 0 // помечаем для удаления лишнюю ')'
			}
		}
	}

	// Помечаем все оставшиеся в стеке '(' для удаления
	for _, idx := range stack {
		runes[idx] = 0
	}
	// runes:  [108 101 101 40 116 40 99 41 111 41 100 101 0]
	// string(runes): lee(t(c)o)de

	// Собираем результат, пропуская помеченные символы
	var result strings.Builder // билдер для эффективного создания результирующей строки
	for _, char := range runes {
		if char != 0 { // проверяем, что символ не помечен на удаление
			result.WriteRune(char) // записываем в билдер все символы кроме помеченных
		}
	}

	// Преобразуем билдер в строку и возвращаем результат
	return result.String()
}
