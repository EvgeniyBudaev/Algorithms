package main

import (
	"fmt"
)

/* Counting sheep...
https://www.codewars.com/kata/54edbc7200b811e956000556/train/go

Рассмотрим массив/список овец, где некоторые овцы могут отсутствовать на своих местах. Нам нужна функция, которая
подсчитывает количество овец, присутствующих в массиве (true означает присутствие).

Examples:
[true,  true,  true,  false,
  true,  true,  true,  true ,
  true,  false, true,  false,
  true,  false, false, true ,
  true,  true,  true,  true ,
  false, false, true,  true]

Правильный ответ — 17.

Подсказка: не забудьте проверить наличие плохих значений, таких как null/undefined
*/

func main() {
	numbers := []bool{true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true}
	fmt.Println(CountSheeps(numbers)) // 17
}

func CountSheeps(numbers []bool) int {
	count := 0

	for _, num := range numbers {
		if num {
			count++
		}
	}

	return count
}
