package main

import (
	"fmt"
)

/* Beginner Series #2 Clock
https://www.codewars.com/kata/55f9bca8ecaa9eac7100004a/train/go

Часы показывают h часов, m минут и s секунд после полуночи.
Ваша задача — написать функцию, которая возвращает время с полуночи в миллисекундах.

Example:
h = 0
m = 1
s = 1

result = 61000
*/

func main() {
	fmt.Println(Past(0, 1, 1)) // 61000
	fmt.Println(Past(1, 1, 1)) // 3661000
}

// Past возвращает время с полуночи в миллисекундах.
// time: O(1), space: O(1)
func Past(h, m, s int) int {
	// 1 час = 3600 секунд = 3600000 миллисекунд
	// 1 минута = 60 секунд = 60000 миллисекунд
	// 1 секунда = 1000 миллисекунд
	return h*3600000 + m*60000 + s*1000
}
