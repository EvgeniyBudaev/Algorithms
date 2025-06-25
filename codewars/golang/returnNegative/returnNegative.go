package main

import (
	"fmt"
)

/* Return Negative
https://www.codewars.com/kata/55685cd7ad70877c23000102/train/go

В этом простом задании вам дано число и вы должны сделать его отрицательным. Но, может быть, число уже отрицательное?

Examples:
MakeNegative(1)    // return -1
MakeNegative(-5)   // return -5
MakeNegative(0)    // return 0
*/

func main() {
	fmt.Println(MakeNegative(1))  // -1
	fmt.Println(MakeNegative(-5)) // -5
}

// MakeNegative возвращает отрицательное число.
// time: O(1), space: O(1)
func MakeNegative(x int) int {
	if x > 0 {
		return -x
	}
	return x
}
