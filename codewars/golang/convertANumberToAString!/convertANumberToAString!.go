package main

import (
	"fmt"
	"strconv"
)

/* Convert a Number to a String!
https://www.codewars.com/kata/5265326f5fda8eb1160004c8/train/go

Нам нужна функция, которая может преобразовать число (целое) в строку.

Examples (input --> output):
123  --> "123"
999  --> "999"
-100 --> "-100"
*/

func main() {
	fmt.Println(NumberToString(123)) // "123"
}

// NumberToString преобразует число в строку.
// time: O(1), space: O(1)
func NumberToString(n int) string {
	return strconv.Itoa(n)
}
