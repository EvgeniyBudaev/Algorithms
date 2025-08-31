package main

import (
	"fmt"
)

/* Length and two values.
https://www.codewars.com/kata/62a611067274990047f431a8/train/go

Даны целое число n и два других значения. Создайте массив размера n, заполненный этими двумя значениями поочередно.

Examples:
5, true, false     -->  [true, false, true, false, true]
10, "blue", "red"  -->  ["blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red"]
0, "one", "two"    -->  []
*/

func main() {
	fmt.Println(Alternate(5, "true", "false")) // [true, false, true, false, true]
	fmt.Println(Alternate(10, "blue", "red"))  // ["blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red"]
	fmt.Println(Alternate(0, "one", "two"))    // []
}

// Alternate создает массив размера n, заполненный этими двумя значениями поочередно.
// time: O(n), space: O(n), где n - длина массива
func Alternate(n int, firstValue string, secondValue string) []string {
	result := make([]string, n)

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			result[i] = firstValue
		} else {
			result[i] = secondValue
		}
	}

	return result
}
