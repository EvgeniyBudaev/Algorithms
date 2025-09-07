package main

import (
	"fmt"
)

/* Wilson primes
https://www.codewars.com/kata/55dc4520094bbaf50e0000cb/train/go

Простые числа Уилсона удовлетворяют следующему условию. Пусть P представляет собой простое число. Тогда
(P−1)!+1 / P * P должно быть целое число, где P! - это факториал числа P.

Ваша задача — создать функцию, которая возвращает значение true, если заданное число является простым числом Уилсона,
и false в противном случае.
*/

func main() {
	fmt.Println(AmIWilson(0)) // false
	fmt.Println(AmIWilson(1)) // false
	fmt.Println(AmIWilson(5)) // true
}

// AmIWilson возвращает true, если n является простым числом Уилсона, и false в противном случае.
// time: O(sqrt(n)), space: O(1)
func AmIWilson(n int) bool {
	// Проверяем, является ли число простым
	if !isPrime(n) {
		return false
	}

	// Для простых чисел 2 и 3 проверяем отдельно
	if n == 2 || n == 3 {
		return false
	}

	// Вычисляем (n-1)! mod (n*n)
	factorialMod := 1
	for i := 2; i < n; i++ {
		factorialMod = (factorialMod * i) % (n * n)
	}

	// Проверяем условие: (n-1)! ≡ -1 mod (n²)
	// Это эквивалентно (n-1)! + 1 ≡ 0 mod (n²)
	return (factorialMod+1)%(n*n) == 0
}

// isPrime проверяет, является ли число простым
// time: O(sqrt(n)), space: O(1)
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
