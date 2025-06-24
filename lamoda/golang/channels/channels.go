package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
		defer close(ch)
	}()

	for i := range ch {
		fmt.Println("i", i)
	}
}
