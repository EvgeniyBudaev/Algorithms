package main

import "fmt"

// MyError структура для ошибок.
type MyError struct {
	data string
}

// Error имплементирует интерфейс error.
func (e *MyError) Error() string {
	return e.data
}

func main() {
	err := foo(4)
	if err != nil {
		fmt.Println("oops")
	} else {
		fmt.Println("ok")
	}
}

// foo возвращает ошибку.
// Bad -> oops
//func foo(i int) error {
//	var err *MyError
//	if i > 5 {
//		err = &MyError{data: "i > 5"}
//	}
//	return err // nil
//}

// Good -> ok
func foo(i int) error {
	if i > 5 {
		return &MyError{data: "i > 5"}
	}
	return nil
}
