package main

import (
	"fmt"
)

/* 172. Factorial Trailing Zeroes
https://leetcode.com/problems/factorial-trailing-zeroes/description/

Дано целое число n, вернуть количество конечных нулей в n!.
Обратите внимание, что n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1.

Input: n = 3
Output: 0
Пояснение: 3! = 6, без конечного нуля.
*/

func main() {
	fmt.Println(trailingZeroes(3)) // 0
}

// trailingZeroes - количество конечных нулей в n!.
// time: O(log5(n)), space: O(1)
func trailingZeroes(n int) int {
	count := 0  // счетчик конечных нулей
	for n > 0 { // пока n больше нуля
		n /= 5     // делим n на 5
		count += n // добавляем количество делений на 5 в счетчик
	}
	return count // возвращаем количество конечных нулей
}
