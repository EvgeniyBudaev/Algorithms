package main

import "fmt"

func main() {
	countdown(5) // => 5 4 3 2 1 0
}

// countdown возвращает i и вызывает себя снова, пока i больше или равно нулю.
func countdown(i int) {
	fmt.Println(i)
	if i <= 0 {
		return
	}

	countdown(i - 1)
}
