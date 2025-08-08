package main

import (
	"fmt"
)

/* Make a function that does arithmetic!
https://www.codewars.com/kata/583f158ea20cfcbeb400000a/train/go

Даны два числа и арифметический оператор (его название, как строки), вернуть результат применения этого оператора к двум числам.
Оба числа a и b будут положительными целыми числами, при этом a всегда будет первым числом в операции, а b — вторым.
Четыре оператора: «сложение», «вычитание», «деление», «умножение».

5, 2, "add"      --> 7
5, 2, "subtract" --> 3
5, 2, "multiply" --> 10
5, 2, "divide"   --> 2.5
*/

func main() {
	fmt.Println(Arithmetic(5, 2, "add"))      // 7
	fmt.Println(Arithmetic(5, 2, "subtract")) // 3
	fmt.Println(Arithmetic(5, 2, "multiply")) // 10
	fmt.Println(Arithmetic(5, 2, "divide"))   // 2.5
}

// Arithmetic возвращает результат применения арифметического оператора к двум числам.
// time: O(1), space: O(1)
func Arithmetic(a int, b int, operator string) int {
	operators := map[string]func(int, int) float64{
		"add": func(a int, b int) float64 {
			return float64(a + b)
		},
		"subtract": func(a int, b int) float64 {
			return float64(a - b)
		},
		"multiply": func(a int, b int) float64 {
			return float64(a * b)
		},
		"divide": func(a int, b int) float64 {
			return float64(a / b)
		},
	}
	op := operators[operator]
	return int(op(a, b))
}
