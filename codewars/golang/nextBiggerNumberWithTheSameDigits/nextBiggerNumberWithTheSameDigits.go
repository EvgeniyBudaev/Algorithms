package main

import (
	"fmt"
	"sort"
)

/* Next bigger number with the same digits
https://www.codewars.com/kata/55983863da40caa2c900004e/train/go

Создайте функцию, которая принимает положительное целое число и возвращает следующее большее число, которое можно
получить путем перестановки его цифр. Например:

12 ==> 21
513 ==> 531
2017 ==> 2071

Если цифры невозможно переставить так, чтобы получилось большее число, вернуть -1 (или nil в Swift, None в Rust):
9 ==> -1
111 ==> -1
531 ==> -1
*/

func main() {
	fmt.Println(NextBigger(2017)) // 2071
}

// NextBigger находит следующее большее число, которое можно получить путем перестановки цифр в числе.
// time O(n*log(n)), space O(n)
func NextBigger(n int) int {
	digits := make([]int, 0, 8) // Максимальное число 9876543210

	for n > 10 { // Разбиваем число на цифры
		digits = append(digits, n%10) // Получаем цифру
		sort.Ints(digits)             // Сортируем цифры
		n /= 10                       // Уменьшаем число на порядок

		for i, digit := range digits { // Ищем первую цифру, большую текущей
			if digit > n%10 { // Если нашли
				digits[i] = n % 10 // Заменяем ее
				sort.Ints(digits)  // Сортируем цифры

				n += digit - n%10          // Добавляем найденную цифру к текущему числу
				for _, d := range digits { // Добавляем цифры к текущему числу
					n = n*10 + d // Добавляем цифру к текущему числу
				}

				return n // Возвращаем число
			}
		}
	}

	return -1 // Если не нашли, возвращаем -1
}
