package main

import "fmt"

func main() {
	fmt.Println(Fib(3)) // 2
}

// Fib вычисляет n-ое число Фибоначчи.
// time: O(2^n), space: O(n), где n - порядковый номер числа
func Fib(n int) int {
	switch {
	case n <= 1: // терминальная ветка
		return n
	default: // рекурсивная ветка
		return Fib(n-1) + Fib(n-2)
	}
}
