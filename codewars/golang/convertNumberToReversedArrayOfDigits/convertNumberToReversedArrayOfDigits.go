package main

import "fmt"

/* Convert number to reversed array of digits
https://www.codewars.com/kata/5583090cbe83f4fd8c000051/train/go

Дано случайное неотрицательное число, необходимо вернуть цифры этого числа в массиве в обратном порядке.

Example (Input => Output):
35231 => [1,3,2,5,3]
0     => [0]
*/

func main() {
	fmt.Println(Digitize(35231)) // [1,3,2,5,3]
	fmt.Println(Digitize(0))     // [0]
}

// Digitize конвертирует число в массив цифр в обратном порядке.
// time: O(n), где n - количество цифр в числе, space: O(n).
func Digitize(n int) []int {
	if n == 0 {
		return []int{0}
	}

	var result []int

	for n > 0 {
		result = append(result, n%10)
		n = n / 10
	}

	return result
}
