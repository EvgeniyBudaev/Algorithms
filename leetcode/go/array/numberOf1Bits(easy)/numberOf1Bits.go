package main

import (
	"fmt"
	"strconv"
	"strings"
)

/* 191. Number of 1 Bits
https://leetcode.com/problems/number-of-1-bits/description/

Для заданного положительного целого числа n напишите функцию, которая возвращает количество установленных битов в его
двоичном представлении (также известном как вес Хэмминга).

Input: n = 11
Output: 3
Пояснение:
Входная двоичная строка 1011 имеет всего три установленных бита.
*/

func main() {
	fmt.Println(hammingWeight(11)) // 3
}

// hammingWeight принимает положительное целое число и возвращает количество установленных битов в его двоичном
// представлении.
// time: O(log(n)), space: O(log(n))
func hammingWeight(n int) int {
	// Преобразовать целое число в двоичное представление
	binaryRepresentation := strconv.FormatInt(int64(n), 2)

	// Извлечь двоичное представление в виде строки и подсчитать биты «1»
	return strings.Count(binaryRepresentation, "1")
}
