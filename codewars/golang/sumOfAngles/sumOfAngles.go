package main

import (
	"fmt"
)

/* Sum of angles
https://www.codewars.com/kata/5a03b3f6a1c9040084001765/train/go

Найдите общую сумму внутренних углов (в градусах) в n-стороннем простом многоугольнике. N будет больше 2.
*/

func main() {
	fmt.Println(Angle(3)) // 180
}

// Angle возвращает общую сумму внутренних углов в n-стороннем простом многоугольнике.
// time: O(1), space: O(1)
func Angle(n int) int {
	return (n - 2) * 180
}
