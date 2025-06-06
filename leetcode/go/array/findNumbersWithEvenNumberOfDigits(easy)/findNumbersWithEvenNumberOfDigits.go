package main

import (
	"fmt"
	"strconv"
)

/* 1295. Find Numbers with Even Number of Digits
https://leetcode.com/problems/find-numbers-with-even-number-of-digits/description/

Найдите числа с четным числом цифр.
Дан массив целых чисел, верните, сколько из них содержат четное количество цифр.

Input: nums = [12,345,2,6,7896]
Output: 2

Input: nums = [555,901,482,1771]
Output: 1
*/

func main() {
	fmt.Println(findNumbers([]int{12, 345, 2, 6, 7896})) // 2
}

// findNumbers возвращает количество чисел с четным числом цифр.
// time: O(n), space: O(1)
func findNumbers(nums []int) int {
	counter := 0 // Счетчик чисел с четным числом цифр

	for _, num := range nums {
		if len(strconv.Itoa(num))%2 == 0 { // Проверяем, является ли число четным
			counter++ // Увеличиваем счетчик
		}
	}

	return counter // Возвращаем счетчик чисел с четным числом цифр
}
