package main

import "fmt"

/* Count by X
https://www.codewars.com/kata/5513795bd3fafb56c200049e/train/go

Создайте функцию с двумя аргументами, которая будет возвращать массив из первых n значений, кратных x.
Предположите, что и заданное число, и количество подсчётов будут положительными числами больше 0.
Верните результаты в виде массива или списка (в зависимости от языка).

Examples:
x = 1, n = 10 --> [1,2,3,4,5,6,7,8,9,10]
x = 2, n = 5  --> [2,4,6,8,10]
*/

func main() {
	fmt.Println(CountBy(1, 5)) // []int{1, 2, 3, 4, 5}
	fmt.Println(CountBy(2, 5)) // []int{2, 4, 6, 8, 10}
}

// CountBy возвращает массив из первых n значений, кратных x.
// time: O(n), space: O(n), n - количество подсчётов
func CountBy(x, n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = (i + 1) * x
	}
	return result
}
