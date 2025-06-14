package main

import (
	"fmt"
)

/* All Balanced Parentheses
https://www.codewars.com/kata/5426d7a2c2c7784365000783/train/go

Напишите функцию, которая создает список строк, представляющих все способы, которыми вы можете сбалансировать n пар скобок.

Examples:
balancedParens 0 -> [""]
balancedParens 1 -> ["()"]
balancedParens 2 -> ["()()","(())"]
balancedParens 3 -> ["()()()","(())()","()(())","(()())","((()))"]
*/

func main() {
	fmt.Println(BalancedParens(1)) // ["()"]
	fmt.Println(BalancedParens(2)) // ["()()","(())"]
}

// BalancedParens генерирует все возможные варианты сбалансированных скобок.
// time: O(2^n), space: O(n), n - количество пар скобок
func BalancedParens(n int) []string {
	var result []string                  // Срез для хранения результатов
	var backtrack func(string, int, int) // Внутренняя рекурсивная функция для генерации последовательностей

	backtrack = func(current string, open int, close int) {
		// Если достигли нужного количества пар скобок
		if len(current) == 2*n {
			result = append(result, current)
			return
		}

		// Добавляем открывающую скобку, если еще не достигли максимума
		if open < n {
			backtrack(current+"(", open+1, close)
		}

		// Добавляем закрывающую скобку, если она соответствует открывающей
		if close < open {
			backtrack(current+")", open, close+1)
		}
	}

	backtrack("", 0, 0) // Начинаем с пустой строки и нулевых счетчиков
	return result
}
