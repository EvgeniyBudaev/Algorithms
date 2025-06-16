package main

import (
	"fmt"
	"unicode"
)

/* Not very secure
https://www.codewars.com/kata/526dbd6c8c0eb53254000110/train/go

В этом примере вам нужно проверить, является ли введенная пользователем строка буквенно-цифровой.
Данная строка не является nil/null/NULL/None, поэтому вам не нужно это проверять.

Строка имеет следующие условия, чтобы быть буквенно-цифровой:

По крайней мере один символ ("" недопустим)
Допустимые символы — заглавные/строчные латинские буквы и цифры от 0 до 9
Без пробелов/подчеркивания
*/

func main() {
	fmt.Println(alphanumeric(".*?"))          // false
	fmt.Println(alphanumeric("a"))            // true
	fmt.Println(alphanumeric("Mazinkaiser"))  // true
	fmt.Println(alphanumeric("hello world_")) // false
	fmt.Println(alphanumeric("PassW0rd"))     // true
	fmt.Println(alphanumeric("     "))        // false
	fmt.Println(alphanumeric(""))             // false
	fmt.Println(alphanumeric("\\n\\t\\n"))    // false
}

// alphanumeric проверяет является ли введенная пользователем строка буквенно-цифровой.
// time: O(n), space: O(1)
func alphanumeric(str string) bool {
	if len(str) == 0 {
		return false
	}

	for _, v := range str {
		if !unicode.IsLetter(v) && !unicode.IsDigit(v) {
			return false
		}
	}

	return true
}
