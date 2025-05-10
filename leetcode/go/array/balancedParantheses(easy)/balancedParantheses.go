package main

import "fmt"

/* Balanced Parantheses

Дана строка A, состоящая только из '(' и ')'.
Вам нужно выяснить, сбалансированы ли скобки в A или нет, если сбалансированы, то вернуть 1, иначе вернуть 0.

Input: A = "(()())"
Output: 1

Input: A = "(()"
Output: 0
*/

func main() {
	str := "(()())"
	fmt.Println(solve(str)) //1
}

// solve - проверяет, сбалансированы ли скобки в строке.
// time: O(n), space: O(1)
func solve(str string) int {
	balance := 0 // Счетчик открывающих скобок

	for _, char := range str {
		if char == '(' {
			balance++
		} else {
			balance--
		}

		// Если баланс стал отрицательным - сразу возвращаем 0
		if balance < 0 {
			return 0
		}
	}

	// Возвращаем 1 если баланс = 0, иначе 0
	if balance == 0 {
		return 1
	}

	// Если баланс не равен нулю, значит скобки не сбалансированы
	return 0
}
