package main

import (
	"fmt"
)

/* Fibonacci
https://www.codewars.com/kata/57a1d5ef7cb1f3db590002af/train/go

Создайте функцию fib, которая возвращает n-й элемент последовательности Фибоначчи (классическая задача программирования).
*/

func main() {
	fmt.Println(Fib(4)) // 3
}

// Fib возвращает n-й элемент последовательности Фибоначчи.
// time: O(n), space: O(1), n - число
func Fib(n int) int {
	a, b := 0, 1
	for ; n > 0; n-- {
		a, b = b, a+b
	}
	return a
}

// Fib2 возвращает n-й элемент последовательности Фибоначчи.
// time: O(2^n), space: O(n), n - число
func Fib2(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)
}
