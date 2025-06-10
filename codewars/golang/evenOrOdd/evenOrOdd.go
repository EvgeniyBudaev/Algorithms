package main

import "fmt"

/* Even or Odd
https://www.codewars.com/kata/53da3dbb4a5168369a0000fe/train/go

Создайте функцию, которая принимает целое число в качестве аргумента и возвращает «Even» для четных чисел или «Odd» для нечетных чисел..
*/

func main() {
	fmt.Println(EvenOrOdd(1)) // Odd
	fmt.Println(EvenOrOdd(2)) // Even
}

// EvenOrOdd возвращает строку "Even" для четных чисел или "Odd" для нечетных чисел.
// time: O(1), space: O(1)
func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	}
	return "Odd"
}
