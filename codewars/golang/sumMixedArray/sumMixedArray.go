package main

import (
	"fmt"
	"strconv"
)

/* Sum Mixed Array
https://www.codewars.com/kata/57eaeb9578748ff92a000009/train/go

Дан массив целых чисел в виде строк и чисел. Верните сумму значений массива, как если бы все они были числами.
Верните ответ в виде числа.
*/

func main() {
	fmt.Println(SumMix([]interface{}{9, 3, "7", "3"})) // 22
}

// SumMix принимает массив чисел в виде строк и чисел и возвращает сумму всех чисел.
// time: O(n), space: O(1), n - длина массива
func SumMix(arr []any) int {
	sum := 0

	for _, v := range arr {
		switch val := v.(type) {
		case int:
			sum += val
		case string:
			num, _ := strconv.Atoi(val)
			sum += num
		}
	}

	return sum
}
