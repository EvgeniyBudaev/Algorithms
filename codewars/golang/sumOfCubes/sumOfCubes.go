package main

import (
	"fmt"
)

/* Sum of Cubes
https://www.codewars.com/kata/59a8570b570190d313000037/train/go

Напишите функцию, которая принимает положительное целое число n, суммирует все кубы всех значений от 1 до n
(включительно) и возвращает эту сумму.

Предположим, что входное число n всегда будет положительным целым числом.
*/

func main() {
	fmt.Println(SumCubes(2)) // 9 (sum of the cubes of 1 and 2 is 1 + 8)
	fmt.Println(SumCubes(3)) // 36 (sum of the cubes of 1, 2, and 3 is 1 + 8 + 27)
}

// SumCubes возвращает сумму кубов всех чисел от 1 до n включительно.
// time: O(n), space: O(1), n - входное число
func SumCubes(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i * i
	}
	return sum
}
