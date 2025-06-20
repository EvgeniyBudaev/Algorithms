package main

import "fmt"

/* Explosive Sum
https://www.codewars.com/kata/52ec24228a515e620b0005ef/train/go

Сколько способов можно составить сумму числа?
Из Википедии: https://en.wikipedia.org/wiki/Partition_(number_theory)

В теории чисел и комбинаторике разбиение положительного целого числа n, также называемое целочисленным разбиением, — это
способ записи n в виде суммы положительных целых чисел. Две суммы, отличающиеся только порядком слагаемых, считаются
одним и тем же разбиением. Если порядок имеет значение, сумма становится композицией. Например, 4 можно разбить пятью
различными способами:
4
3 + 1
2 + 2
2 + 1 + 1
1 + 1 + 1 + 1

Examples:
Basic:
ExpSum(1) // 1
ExpSum(2) // 2 -> 1+1 , 2
ExpSum(3) // 3 -> 1+1+1, 1+2, 3
ExpSum(4) // 5 -> 1+1+1+1, 1+1+2, 1+3, 2+2, 4
ExpSum(5) // 7 -> 1+1+1+1+1, 1+1+1+2, 1+1+3, 1+2+2, 1+4, 5, 2+3

ExpSum(10) // 42
Explosive:
ExpSum(50) // 204226
ExpSum(80) // 15796476
ExpSum(100) // 190569292
*/

func main() {
	fmt.Println(ExpSum(3)) // 3 -> 1+1+1, 1+2, 3
}

// ExpSum возвращает количество способов разбить число на сумму натуральных чисел.
// time: O(n^2), space: O(n), где n - входное число.
func ExpSum(n uint64) uint64 {
	if n == 0 { // Если число равно 0, то есть только один способ разбить его на сумму натуральных чисел: 0.
		return 1
	}

	dp := make([]uint64, n+1) // Создаем слайс для хранения количества способов для каждого числа от 0 до n
	dp[0] = 1                 // Базовый случай: одно разбиение для 0

	for i := uint64(1); i <= n; i++ { // Для каждого числа от 1 до n
		for j := i; j <= n; j++ { // Для каждого числа от i до n
			dp[j] += dp[j-i] // Добавляем количество способов для числа j, если мы вычитаем число i из него
		}
	}

	return dp[n] // Возвращаем количество способов разбить число n на сумму натуральных чисел
}
