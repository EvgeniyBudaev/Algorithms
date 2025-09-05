package main

import (
	"fmt"
	"strconv"
)

/* Convert a String to a Number!
https://www.codewars.com/kata/544675c6f971f7399a000e79/train/go

Нам нужна функция, которая может преобразовать строку в число. Какие способы этого добиться вы знаете?
Примечание: Не волнуйтесь, все входные данные будут строками, и каждая строка — это совершенно допустимое представление
целого числа.

Examples:
"1234" --> 1234
"605"  --> 605
"1405" --> 1405
"-7" --> -7
*/

func main() {
	fmt.Println(StringToNumber("1234")) // 1234
}

// StringToNumber возвращает число, представленное строкой str.
// time: O(1), space: O(1)
func StringToNumber(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}
