package main

import (
	"fmt"
	"math/big"
)

/* Last digit of a large number
https://www.codewars.com/kata/5511b2f550906349a70004e1/train/go

Определите функцию, которая принимает два неотрицательных целых числа a и b и возвращает последнюю десятичную цифру a^b.
*/

func main() {
	fmt.Println(LastDigit("9", "7")) // 9
	//fmt.Println(LastDigit("1606938044258990275541962092341162602522202993782792835301376", "2037035976334486086268445688409378161051468393665936250636140449354381299763336706183397376")) // 6
}

// LastDigit возвращает последнюю десятичную цифру a^b.
// time: O(1), space: O(1)
func LastDigit(n1, n2 string) int {
	a, b := big.NewInt(0), big.NewInt(0) // Преобразование строки в число
	a.SetString(n1, 10)
	b.SetString(n2, 10)

	// Exp возводит число в степень, результат записывается в a. Вычисляет a^b mod 10.
	// *Без вычисления по модулю, используйте метод Exp с nil в качестве третьего аргумента.
	a.Exp(a, b, big.NewInt(10))

	return int(a.Int64()) // Преобразование числа в int
}
