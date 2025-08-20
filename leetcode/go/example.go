package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4}
	b := a[2:3] // b = [3]
	fmt.Println(b, len(b), cap(b))
}
