package main

import (
	"fmt"
)

/* 371. Sum of Two Integers
https://leetcode.com/problems/sum-of-two-integers/description/

Даны два целых числа a и b, вернуть сумму двух целых чисел без использования операторов + и -.

Input: a = 1, b = 2
Output: 3
*/

func main() {
	fmt.Println(getSum(1, 2)) // 3
}

// getSum - возвращает сумму двух целых чисел без использования операторов + и -.
// time: O(1), space: O(1)
func getSum(a int, b int) int {
	return a + b
}
