package main

import (
	"fmt"
)

/* Strong Number (Special Numbers Series #2)
https://www.codewars.com/kata/5a4d303f880385399b000001/train/go

Сильное число — это число, сумма факториала цифр которого равна самому числу.
Например, 145 — сильное число, потому что 1! + 4! + 5! = 1 + 24 + 120 = 145.

Задание
Данному положительному числу определить, является ли оно сильным, и вернуть ответ «СИЛЬНОЕ!!!!» или «Не сильное!!».
*/

func main() {
	fmt.Println(Strong(1))   // "STRONG!!!!"
	fmt.Println(Strong(145)) // "STRONG!!!!"
	fmt.Println(Strong(7))   // "Not Strong !!"
}

// Strong возвращает строку "STRONG!!!!" или "Not Strong !!" в зависимости от того, является ли число сильным.
// time: O(n), space: O(1), n - количество цифр в числе
func Strong(n int) string {
	original := n
	sum := 0

	// Вычисляем сумму факториалов цифр числа
	for n > 0 {
		digit := n % 10
		sum += Factorial(digit)
		n /= 10
	}

	// Проверяем, равна ли сумма факториалов исходному числу
	if sum == original {
		return "STRONG!!!!"
	}
	return "Not Strong !!"
}

// Factorial возвращает факториал числа n.
// time: O(n), space: O(n), n - число
func Factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * Factorial(n-1)
}
