package main

import "fmt"

func main() {
	// hello, adit!
	// how are you, adit?
	// getting ready to say bye...
	// ok bye!
	greet("adit")
}

// greet2 - выводит приветствие.
func greet2(name string) {
	fmt.Println("how are you, " + name + "?")
}

// bye - завершает процесс.
func bye() {
	fmt.Println("ok bye!")
}

// greet - выводит приветствие и вызывает другие функции.
func greet(name string) {
	fmt.Println("hello, " + name + "!")
	greet2(name)
	fmt.Println("getting ready to say bye...")
	bye()
}
