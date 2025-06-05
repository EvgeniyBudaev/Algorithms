package main

import (
	"fmt"
	"sort"
)

// Are they the "same"?

func main() {
	a1 := []int{121, 144, 19, 161, 19, 144, 19, 11}
	b1 := []int{121, 14641, 20736, 361, 25921, 361, 20736, 361}
	fmt.Println(Comp(a1, b1)) // true

	a2 := []int{121, 144, 19, 161, 19, 144, 19, 11}
	b2 := []int{132, 14641, 20736, 361, 25921, 361, 20736, 361}
	fmt.Println(Comp(a2, b2)) // false
}

// Comp принимает два массива чисел и проверяет, являются ли они квадратами друг друга.
// time: O(n), space: O(1)
func Comp(array1 []int, array2 []int) bool {
	if array1 == nil || array2 == nil || (len(array1) != len(array2)) {
		return false
	}

	for i, num := range array1 {
		array1[i] = num * num
	}

	sort.Ints(array1)
	sort.Ints(array2)

	for i := range array1 {
		if array1[i] != array2[i] {
			return false
		}
	}

	return true
}
