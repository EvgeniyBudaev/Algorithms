package main

import (
	"fmt"
)

/* You Can't Code Under Pressure #1
https://www.codewars.com/kata/53ee5429ba190077850011d4/train/go

Пишите код как можно быстрее! Вам нужно удвоить целое число и вернуть его.
*/

func main() {
	fmt.Println(DoubleInteger(1)) // 2
}

// DoubleInteger удваивает целое число.
// time: O(1), space: O(1)
func DoubleInteger(i int) int {
	return i * 2
}
