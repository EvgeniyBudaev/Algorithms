package main

import (
	"fmt"
	"strconv"
)

/* 1323. Maximum 69 Number

Вам дано целое положительное число, состоящее только из цифр 6 и 9.
Верните максимальное число, которое вы можете получить, изменив не более одной цифры
(6 становится 9, а 9 становится 6).

Input: num = 9669
Output: 9969
Explanation:
Changing the first digit results in 6669.
Changing the second digit results in 9969.
Changing the third digit results in 9699.
Changing the fourth digit results in 9666.
The maximum number is 9969.

Input: num = 9996
Output: 9999
Explanation: Changing the last digit 6 to 9 results in the maximum number.

Input: num = 9999
Output: 9999
Explanation: It is better not to apply any change.
*/

func main() {
	fmt.Println(maximum69Number(9996)) // 9999
}

// maximum69Number возвращает максимальное число, полученное из числа с заменой одной цифры 6 на 9
// time: O(n), space: O(n), где n - длина строки
func maximum69Number(num int) int {
	s := strconv.Itoa(num) // Преобразуем число в строку, чтобы работать с отдельными цифрами
	digits := []rune(s)    // Преобразуем строку в слайс рун для изменения символов

	// Ищем первую 6 и заменяем ее на 9
	for i, char := range digits {
		if char == '6' {
			digits[i] = '9'
			break // Заменяем только первую 6
		}
	}

	// Преобразуем обратно в число
	result, _ := strconv.Atoi(string(digits))
	return result // Возвращаем максимальное число
}
