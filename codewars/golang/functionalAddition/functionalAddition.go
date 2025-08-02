package main

import (
	"fmt"
)

/* Functional Addition
https://www.codewars.com/kata/538835ae443aae6e03000547/train/go

Создайте функцию add(n)/Add(n), которая возвращает функцию, всегда прибавляющую n к любому числу.

var addOne = Add(1)
addOne(3) // 4
*/

func main() {
	fmt.Println(Add(1)(3)) // 4
}

// Add возвращает функцию, которая прибавляет n к любому числу.
// time: O(1), space: O(1)
func Add(n int) func(int) int {
	return func(i int) int {
		return n + i
	}
}
