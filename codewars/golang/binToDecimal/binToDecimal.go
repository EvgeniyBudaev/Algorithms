package main

import (
	"fmt"
	"strconv"
)

/* Bin to Decimal
https://www.codewars.com/kata/57a5c31ce298a7e6b7000334/train/go

Завершите функцию, которая преобразует двоичное число (заданное в виде строки) в десятичное число.
*/

func main() {
	fmt.Println(BinToDec("0"))  // 0
	fmt.Println(BinToDec("1"))  // 1
	fmt.Println(BinToDec("10")) // 2
}

// BinToDec принимает двоичное число в виде строки и возвращает его десятичное представление.
// time: O(n), space: O(1), где n - количество символов в двоичном числе
func BinToDec(bin string) int {
	i, _ := strconv.ParseInt(bin, 2, 64)
	return int(i)
}
