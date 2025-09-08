package main

import (
	"fmt"
)

/* Round up to the next multiple of 5
https://www.codewars.com/kata/55d1d6d5955ec6365400006d/train/go

Если на вход подано целое число, можно ли округлить его до ближайшего (то есть «больше или равно») кратного 5 числа?

Examples:
input:    output:
0    ->   0
2    ->   5
3    ->   5
12   ->   15
21   ->   25
30   ->   30
-2   ->   0
-5   ->   -5
etc.
*/

func main() {
	fmt.Println(RoundToNext5(0))     // 0
	fmt.Println(RoundToNext5(3))     // 5
	fmt.Println(RoundToNext5(12))    // 15
	fmt.Println(RoundToNext5(-1))    // 0
	fmt.Println(RoundToNext5(-2))    // 0
	fmt.Println(RoundToNext5(-5))    // -5
	fmt.Println(RoundToNext5(-6373)) // -6370
}

// RoundToNext5 округляет число до ближайшего кратного 5.
// time: O(1), space: O(1)
func RoundToNext5(n int) int {
	if n%5 == 0 {
		return n
	}
	if n > 0 {
		return n + (5 - n%5)
	}
	return n - (n % 5)
}
