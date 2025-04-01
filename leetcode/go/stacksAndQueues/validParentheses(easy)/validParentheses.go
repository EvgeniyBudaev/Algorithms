package main

import "fmt"

/* https://leetcode.com/problems/valid-parentheses/description/

Учитывая строку s, содержащую только символы '(', ')', '{', '}', '[' и ']', определите, является ли входная строка
допустимой.

Входная строка действительна, если:

Открытые скобки должны закрываться скобками того же типа.
Открытые скобки должны закрываться в правильном порядке.
Каждой закрывающей скобке соответствует открытая скобка того же типа.

Input: s = "()"
Output: true

Input: s = "()[]{}"
Output: true

Input: s = "(]"
Output: false
*/

func main() {
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
}

func isValid(s string) bool {
	stack := make([]rune, 0)
	brackets := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	for _, char := range s {
		// Если это закрывающая скобка
		if matching, isClose := brackets[char]; isClose {
			// Проверяем, соответствует ли последняя открывающая скобка
			if len(stack) == 0 || stack[len(stack)-1] != matching {
				return false
			}
			// Удаляем соответствующую открывающую скобку из стека
			stack = stack[:len(stack)-1]
		} else {
			// Добавляем открывающую скобку в стек
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
