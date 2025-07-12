package main

import (
	"fmt"
)

/* Opposite number
https://www.codewars.com/kata/56dec885c54a926dcd001095/train/go

Очень просто: дано число (целое / десятичное / и то, и другое в зависимости от языка), найти его противоположность
(аддитивное обратное).

Examples:

1: -1
14: -14
-34: 34
*/

func main() {
	fmt.Println(Opposite(1))   // -1
	fmt.Println(Opposite(14))  // -14
	fmt.Println(Opposite(-34)) // 34
}

// Opposite возвращает противоположное число.
// time: O(1), space: O(1)
func Opposite(value int) int {
	return -value
}
