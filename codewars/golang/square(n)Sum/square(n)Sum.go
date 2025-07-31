package main

import (
	"fmt"
)

/* Square(n) Sum
https://www.codewars.com/kata/515e271a311df0350d00000f/train/go

Дополните функцию квадратной суммы так, чтобы она возводила в квадрат каждое переданное ей число,
а затем суммировала результаты.

Например, для [1, 2, 2] она должна вернуть 9, потому что 1^2 + 2^2 + 2^2 = 9.
*/

func main() {
	numbers := []int{1, 2, 2}
	fmt.Println(SquareSum(numbers)) // 9
}

// SquareSum возводит в квадрат каждое переданное число и возвращает сумму квадратов.
// time: O(n), space: O(1), n - количество чисел
func SquareSum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number * number
	}
	return result
}
