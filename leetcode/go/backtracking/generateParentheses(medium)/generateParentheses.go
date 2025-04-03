package main

import "fmt"

/* https://leetcode.com/problems/generate-parentheses/description/

Учитывая n пар круглых скобок, напишите функцию, которая генерирует все комбинации правильных круглых скобок.

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

Input: n = 1
Output: ["()"]
*/

func main() {
	fmt.Println(generateParenthesis(3)) // ["((()))","(()())","(())()","()(())","()()()"]
}

// generateParenthesis генерирует все правильные скобочные последовательности длины 2n
func generateParenthesis(n int) []string {
	var result []string

	// Внутренняя рекурсивная функция для генерации последовательностей
	var backtrack func(string, int, int)
	backtrack = func(current string, open int, close int) {
		// Если достигли нужной длины, добавляем в результат
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
