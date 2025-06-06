package main

import "fmt"

/* 1952. Three Divisors
https://leetcode.com/problems/three-divisors/description/

Дано целое число n, вернуть true, если n имеет ровно три положительных делителя. В противном случае вернуть false.
Целое число m является делителем n, если существует целое число k, такое что n = k * m.

Input: n = 2
Output: false
Объяснение: у числа 2 всего два делителя: 1 и 2.

Input: n = 4
Output: true
Объяснение: 4 имеет три делителя: 1, 2 и 4.
*/

func main() {
	fmt.Println(isThree(2)) // false
	fmt.Println(isThree(4)) // true
}

// isThree принимает число и возвращает true если число имеет ровно три положительных делителя, иначе false.
// time: O(n), space: O(1)
func isThree(n int) bool {
	// Если длина результирующего массива равна 3, значит число имеет ровно три положительных делителя
	if len(div(n)) == 3 {
		return true
	}
	// Если длина результирующего массива больше 3, значит число имеет больше трех положительных делителей
	return false
}

// div принимает число и возвращает его делители.
// time: O(n), space: O(n)
func div(n int) []int {
	var result []int // Создаем пустой массив для хранения делителей числа

	for i := 1; i <= n; i++ { // Проходимся по всем числам от 1 до n
		if n%i == 0 { // Если число делится нацело на i, то оно является делителем числа
			result = append(result, i)

		}
	}

	return result // Возвращаем массив делителей числа
}
