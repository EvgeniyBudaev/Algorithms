package main

import (
	"fmt"
	"strconv"
)

/* https://leetcode.com/problems/evaluate-reverse-polish-notation/description/

Вам дан массив строковых токенов, который представляет арифметическое выражение в обратной польской нотации.
Оцените выражение. Возвращает целое число, представляющее значение выражения.

Обратите внимание, что:

Допустимые операторы: «+», «-», «*» и «/».
Каждый операнд может быть целым числом или другим выражением.
Деление между двумя целыми числами всегда сокращается до нуля.
Никакого деления на ноль не будет.
Входные данные представляют собой допустимое арифметическое выражение в обратной польской записи.
Ответ и все промежуточные вычисления можно представить в виде 32-битного целого числа.

Input: tokens = ["2","1","+","3","*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9

Input: tokens = ["4","13","5","/","+"]
Output: 6
Explanation: (4 + (13 / 5)) = 6
*/

func main() {
	tokens := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(tokens)) // 9
}

func evalRPN(tokens []string) int {
	stack := []int{}
	operators := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}

	for _, token := range tokens {
		if op, exists := operators[token]; exists {
			// Получаем два последних числа из стека
			rightNum := stack[len(stack)-1]
			leftNum := stack[len(stack)-2]
			stack = stack[:len(stack)-2] // Удаляем их из стека

			// Выполняем операцию и добавляем результат обратно в стек
			result := op(leftNum, rightNum)
			stack = append(stack, result)
		} else {
			// Если это число, преобразуем в int и добавляем в стек
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}
