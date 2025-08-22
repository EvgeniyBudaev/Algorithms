package main

import (
	"fmt"
)

/* String ends with?
https://www.codewars.com/kata/51f2d1cafc9c0f745c00037d/train/go

Дополните решение так, чтобы оно возвращало значение true, если первый переданный аргумент (строка) заканчивается вторым
аргументом (тоже строкой).

Примеры:

solution("abc", "bc") // true
solution("abc", "d") // false
*/

func main() {
	fmt.Println(solution("abc", "bc"))     // true
	fmt.Println(solution("abc", "d"))      // false
	fmt.Println(solution("", ""))          // true
	fmt.Println(solution(" ", ""))         // true
	fmt.Println(solution("abc", "c"))      // true
	fmt.Println(solution("banana", "ana")) // true
}

// solution возвращает true, если первый аргумент (строка) заканчивается вторым аргументом (строкой).
// time: O(n), space: O(1), n - длина строки
func solution(str, ending string) bool {
	if len(ending) > len(str) {
		return false
	}
	return str[len(str)-len(ending):] == ending
}
