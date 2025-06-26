package main

import "fmt"

func main() {
	fmt.Println(fact(3))
}

// fact вычисляет факториал.
// time: O(n), space: O(n), где n - число
func fact(n int) int {
	if n == 0 { // терминальная ветка — то есть условие выхода из рекурсии
		return 1
	} else { // рекурсивная ветка
		return n * fact(n-1)
	}
}
