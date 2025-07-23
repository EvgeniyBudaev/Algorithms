package main

import "fmt"

func main() {
	fmt.Println(isEqual(1, 2))     // false
	fmt.Println(isEqual(1.3, 1.3)) // true
}

func isEqual[T comparable](lhs, rhs T) bool {
	return lhs == rhs
}

type Number interface {
	int64 | float64
}

func SunNumber[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
