package main

import "fmt"

func main() {
	var myTypePointer *MyType
	fmt.Println(myTypePointer == nil) // true
	fmt.Println(myTypePointer.Foo())  // "Hello"
}

type MyType struct {
	count int
}

func (m *MyType) Foo() string {
	return "Hello"
}
