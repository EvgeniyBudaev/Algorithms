package main

import (
	"fmt"
)

/* Returning Strings
https://www.codewars.com/kata/55a70521798b14d4750000a4/train/go

Создайте функцию, которая принимает параметр, представляющий имя, и возвращает сообщение:
"Hello, <name> how are you doing today?".

[Убедитесь, что вы вводите именно то, что я написал, иначе программа может работать неправильно]
*/

func main() {
	fmt.Println(Greet("Ryan")) // "Hello, Ryan how are you doing today?"
}

// Greet возвращает сообщение "Hello, <name> how are you doing today?".
// time: O(1), space: O(1)
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s how are you doing today?", name)
}
