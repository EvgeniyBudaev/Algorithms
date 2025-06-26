package main

import "fmt"

func main() {
	fmt.Println(Fib(3)) // 2
}

// Рекурсивный вариант.
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

// Итеративный вариант.
// Fib вычисляет n-ое число Фибоначчи.
// time: O(n), space: O(1), где n - порядковый номер числа
//func Fib(n int) int {
//	a, b := 0, 1
//	for n > 0 {
//		a, b = b, a+b
//		n--
//	}
//	return a
//}
