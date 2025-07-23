package main

import "fmt"

func main() {
	// defer1 -> defer2 -> panic
	defer func() {
		fmt.Println("defer2")
	}()

	F()
}

func F() {
	defer func() {
		fmt.Println("defer1")
	}()

	panic("panic")
}
