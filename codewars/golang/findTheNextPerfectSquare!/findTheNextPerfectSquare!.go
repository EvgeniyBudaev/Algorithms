package main

import (
	"fmt"
	"math"
)

/* Find the next perfect square!
https://www.codewars.com/kata/56269eb78ad2e4ced1000013/train/go

Возможно, вы знаете несколько довольно больших полных квадратов. Но как насчёт следующего?

Завершите метод findNextSquare, который находит следующий полный квадрат после переданного в качестве параметра.
Напомним, что полный квадрат — это целое число n, такое что sqrt(n) также является целым числом.

Если аргумент сам по себе не является полным квадратом, верните либо -1, либо пустое значение, например, None или null,
в зависимости от используемого языка. Можно считать, что аргумент неотрицательный.
*/

func main() {
	fmt.Println(FindNextSquare(121)) // 144
	fmt.Println(FindNextSquare(155)) // -1
}

// FindNextSquare находит следующий полный квадрат.
// time: O(1), space: O(1)
func FindNextSquare(sq int64) int64 {
	// Находим квадратный корень из входного числа
	root := math.Sqrt(float64(sq))

	// Проверяем, является ли корень целым числом
	if root != math.Floor(root) {
		return -1
	}

	// Увеличиваем корень на 1 и возводим в квадрат
	nextRoot := int64(root) + 1
	return nextRoot * nextRoot
}
