package main

import "fmt"

func main() {
	fmt.Println(fact(5)) // 120
}

// fact возвращает факториал числа.
func fact(x int) int {
	if x == 0 {
		return 1
	}

	return x * fact(x-1)
}
