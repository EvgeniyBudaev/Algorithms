package main

import "fmt"

/* Sum of positive
https://www.codewars.com/kata/5715eaedb436cf5606000381/train/go

Получаете массив чисел, возвращаете сумму всех положительных единиц.

Example:
[1, -4, 7, 12] => 1 + 7 + 12 = 20
*/

func main() {
	numbers := []int{1, -4, 7, 12}
	fmt.Println(PositiveSum(numbers)) // 20
}

// PositiveSum возвращает сумму положительных чисел в массиве.
// time O(n), space O(1)
func PositiveSum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		if number > 0 {
			sum += number
		}
	}

	return sum
}
