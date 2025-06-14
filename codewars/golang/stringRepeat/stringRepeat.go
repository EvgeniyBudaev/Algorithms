package main

import (
	"fmt"
)

/* String repeat
https://www.codewars.com/kata/57a0e5c372292dd76d000d7e/train/go

Напишите функцию, которая принимает неотрицательное целое число n и строку s в качестве параметров и
возвращает строку s, повторенную ровно n раз.

Examples (input -> output)
6, "I"     -> "IIIIII"
5, "Hello" -> "HelloHelloHelloHelloHello"
*/

func main() {
	fmt.Println(RepeatStr(5, "Hello")) // "HelloHelloHelloHelloHello"
}

// RepeatStr возвращает строку s, повторенную ровно n раз.
// time: O(n), space: O(n)
func RepeatStr(repetitions int, value string) string {
	result := ""

	for i := 0; i < repetitions; i++ {
		result += value
	}

	return result
}
