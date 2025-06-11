package main

import (
	"fmt"
	"strings"
)

/* Abbreviate a Two Word Name
https://www.codewars.com/kata/57eadb7ecd143f4c9c0000a3/train/go

Напишите функцию для преобразования имени в инициалы. Эта ката строго принимает два слова с одним пробелом между ними.

Вывод должен состоять из двух заглавных букв с точкой между ними.

Это должно выглядеть так:

Sam Harris => S.H
patrick feeney => P.F
*/

func main() {
	fmt.Println(AbbrevName("Sam Harris")) // "S.H"
}

// AbbrevName возвращает инициалы.
// time: O(n), space: O(1)
func AbbrevName(name string) string {
	if len(name) == 0 {
		return ""
	}

	parts := strings.Fields(name)
	firstInitial := strings.ToUpper(string(parts[0][0]))
	secondInitial := strings.ToUpper(string(parts[1][0]))

	return firstInitial + "." + secondInitial
}
