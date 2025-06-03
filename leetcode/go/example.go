package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 10
		defer close(ch)
	}()
	v, closed := <-ch
	fmt.Println(v, closed)
	v, closed = <-ch
	fmt.Println(v, closed)
}
