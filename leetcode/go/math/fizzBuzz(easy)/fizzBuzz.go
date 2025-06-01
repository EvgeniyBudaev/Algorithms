package main

import (
	"fmt"
	"strconv"
)

/* 412. Fizz Buzz
https://leetcode.com/problems/fizz-buzz/description/

Дано целое число n, вернуть ответ в виде массива строк (индексированный на 1), где:

answer[i] == "FizzBuzz", если i делится на 3 и 5.
answer[i] == "Fizz", если i делится на 3.
answer[i] == "Buzz", если i делится на 5.
answer[i] == i (в виде строки), если ни одно из вышеперечисленных условий не выполняется.

Input: n = 3
Output: ["1","2","Fizz"]
*/

func main() {
	fmt.Println(fizzBuzz(3)) // ["1","2","Fizz"]
}

// fizzBuzz принимает целое число n и возвращает массив строк в соответствии с условием задачи.
// tine: O(n), space: O(n)
func fizzBuzz(n int) []string {
	var output []string // Инициализация результирующего массива строк

	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			output = append(output, "FizzBuzz")
		case i%3 == 0 && i%5 != 0:
			output = append(output, "Fizz")
		case i%5 == 0 && i%3 != 0:
			output = append(output, "Buzz")
		default:
			output = append(output, strconv.Itoa(i))
		}
	}

	return output
}
