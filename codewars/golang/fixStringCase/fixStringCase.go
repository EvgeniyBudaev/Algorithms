package main

import (
	"fmt"
	"strings"
	"unicode"
)

/* Fix string case
https://www.codewars.com/kata/5b180e9fedaa564a7000009a/train/go

В этом задании вам будет дана строка, которая может содержать как строчные, так и заглавные буквы. Ваша задача —
преобразовать эту строку либо в строку только со строчными, либо в строку только со строчными буквами, основываясь на
как можно меньшем количестве изменений;
если строка содержит одинаковое количество строчных и заглавных букв, преобразовать её в строку со строчными.

solve("coDe") = "code". Lowercase characters > uppercase. Change only the "D" to lowercase.
solve("CODe") = "CODE". Uppercase characters > lowecase. Change only the "e" to uppercase.
solve("coDE") = "code". Upper == lowercase. Change all to lowercase.
*/

func main() {
	fmt.Println(solve("coDe")) // "code"
	fmt.Println(solve("CODe")) // "CODE"
	fmt.Println(solve("coDE")) // "code"
}

// solve преобразует строку в строку только со строчными буквами, основываясь на как можно меньшем количестве изменений.
// time: O(n), space: O(1), где n - длина строки
func solve(str string) string {
	lowerCount, upperCount := 0, 0

	for _, char := range str {
		if unicode.IsLower(char) {
			lowerCount++
		} else {
			upperCount++
		}
	}

	if upperCount > lowerCount {
		return strings.ToUpper(str)
	}

	return strings.ToLower(str)
}
