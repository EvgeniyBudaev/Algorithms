package main

import (
	"fmt"
)

/* 50. Pow(x, n)
https://leetcode.com/problems/powx-n/description/

Реализуйте pow(x, n), которая вычисляет x, возведенный в степень n (т. е. xn).

Input: x = 2.00000, n = 10
Output: 1024.00000
*/

func main() {
	fmt.Println(myPow(2, 10)) // 1024
}

// myPow - рекурсивный метод вычисления степени числа.
func myPow(x float64, n int) float64 {
	N := int64(n) // преобразование степени числа в int64
	if N < 0 {    // проверка на отрицательную степень числа.
		x = 1 / x // деление на единицу
		N = -N    // преобразование степени числа в положительное
	}
	return powHelper(x, N) // вызов рекурсивного метода для вычисления степени числа.
}

// powHelper - рекурсивный метод вычисления степени числа.
// time: O(log n), space: O(log n)
func powHelper(x float64, n int64) float64 {
	if n == 0 { // проверка на равенство степени числа нулю.
		return 1 // возвращаем единицу
	}
	half := powHelper(x, n/2) // рекурсивный вызов метода для вычисления степени числа.
	result := half * half     // возведение в квадрат
	if n%2 == 1 {             // проверка на нечетность степени числа.
		result *= x // умножение на число
	}
	return result // возвращаем результат
}

// myPow - рекурсивный метод вычисления степени числа.
// time: O(log n), space: O(log n)
//func myPow(x float64, n int) float64 {
//	return math.Pow(x, float64(n))
//}
