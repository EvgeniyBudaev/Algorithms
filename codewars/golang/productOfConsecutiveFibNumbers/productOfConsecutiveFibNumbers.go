package main

import "fmt"

/* Product of consecutive Fib numbers
https://www.codewars.com/kata/5541f58a944b85ce6d00006a/train/go

Числа Фибоначчи — это числа в следующей целочисленной последовательности (Fn):
0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, ...

такой, что:
F(0)=0
F(1)=1
F(n)=F(n−1)+F(n−2)

Дано число, скажем, prod (для произведения), мы ищем два числа Фибоначчи F(n) и F(n+1), проверяя:
F(n)∗F(n+1)=prod
*/

func main() {
	fmt.Println(ProductFib(4895)) // [55 89 1]
}

// ProductFib возвращает массив из двух чисел Фибоначчи, которые умножены друг на друга равны prod.
// time: O(n), space: O(1), где n - количество чисел Фибоначчи
func ProductFib(prod uint64) [3]uint64 {
	// Начинаем с первых двух чисел Фибоначчи
	a, b := uint64(0), uint64(1)

	for a*b < prod {
		// Переходим к следующей паре чисел Фибоначчи
		a, b = b, a+b
	}

	// Проверяем, равно ли произведение prod
	if a*b == prod {
		return [3]uint64{a, b, 1} // 1 означает true (нашли точное совпадение)
	}
	return [3]uint64{a, b, 0} // 0 означает false (не нашли точное совпадение)
}
