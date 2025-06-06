package main

import (
	"fmt"
	"strconv"
)

/* 9. Palindrome Number
https://leetcode.com/problems/palindrome-number/description/

Учитывая целое число x, верните true, если x является палиндром и false в противном случае.

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
*/

func main() {
	fmt.Println(isPalindrome(121))  // true
	fmt.Println(isPalindrome(-121)) // false
}

// isPalindrome проверяет, является ли целое число x палиндромом.
// time: O(n), space: O(1)
func isPalindrome(x int) bool {
	str := strconv.Itoa(x) // Преобразуем число в строку
	left, right := 0, len(str)-1

	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}

	return true
}

// isPalindrome - проверяет, является ли целое число x палиндромом.
// time: O(n), space: O(1)
//func isPalindrome(x int) bool {
//	num := x
//	var palindrome int // Переменная для хранения "перевернутого" числа
//	for i := 0; num > 0; i++ {
//		mod := num % 10                  // Получаем последнюю цифру числа num (остаток от деления на 10)
//		num = num / 10                   // Убираем последнюю цифру из числа num, деля его на 10
//		palindrome = palindrome*10 + mod // Добавляем полученную цифру mod к "перевернутому" числу palindrome, сдвигая его на один разряд влево
//	}
//	return palindrome == x
//}
