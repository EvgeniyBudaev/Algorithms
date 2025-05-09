package main

import (
	"fmt"
	"math"
)

/* 7. Reverse Integer
https://leetcode.com/problems/reverse-integer/description/

Дано знаковое 32-битное целое число x, вернуть x с перевернутыми цифрами. Если переворот x приводит к выходу значения
за пределы знакового 32-битного целого числа [-231, 231 - 1], то вернуть 0.

Предположим, что среда не позволяет вам хранить 64-битные целые числа (со знаком или без знака).

Input: x = 123
Output: 321
*/

func main() {
	fmt.Println(reverse(123)) // 321
}

// reverse - Переворачивает число.
// Time: O(n), Space: O(1)
func reverse(x int) int {
	sign := 1 // По умолчанию положительное число

	// Проверяем отрицательное ли число
	if x < 0 {
		sign = -1 // Если отрицательное, меняем знак
		x = -x    // Игнорируем минус при перевороте
	}

	res := 0     // Результат
	for x != 0 { // Пока число не 0
		res = res*10 + x%10 // Добавляем последнюю цифру в результат
		x /= 10             // Убираем последнюю цифру
	}
	res *= sign // Возвращаем знак

	// Проверяем 32-битные границы
	if res < math.MinInt32 || res > math.MaxInt32 {
		return 0
	}

	// Возвращаем результат
	return res
}

// reverse - Переворачивает число.
// Time: O(n), Space: O(n)
//func reverse(x int) int {
//	// Преобразуем число в строку
//	s := strconv.Itoa(x)
//
//	// Разделяем строку на руны (для корректной работы с Unicode)
//	runes := []rune(s)
//
//	// Проверяем отрицательное ли число
//	isNegative := false
//	if runes[0] == '-' {
//		isNegative = true
//		runes = runes[1:] // Игнорируем минус при перевороте
//	}
//
//	// Переворачиваем руны
//	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
//		runes[i], runes[j] = runes[j], runes[i]
//	}
//
//	// Если было отрицательное, добавляем минус обратно
//	if isNegative {
//		runes = append([]rune{'-'}, runes...)
//	}
//
//	// Преобразуем обратно в число
//	reversedStr := string(runes)
//	reversed, err := strconv.Atoi(reversedStr)
//	if err != nil {
//		return 0
//	}
//
//	// Проверяем 32-битные границы
//	if reversed > math.MaxInt32 || reversed < math.MinInt32 {
//		return 0
//	}
//
//	return reversed
//}
