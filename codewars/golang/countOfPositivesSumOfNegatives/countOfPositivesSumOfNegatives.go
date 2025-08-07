package main

import (
	"fmt"
)

/* Count of positives / sum of negatives
https://www.codewars.com/kata/576bb71bbbcf0951d5000044/train/go

Дан массив целых чисел.

Вернуть массив, где первый элемент — количество положительных чисел, а второй — сумма отрицательных чисел.
0 не является ни положительным, ни отрицательным.

Если входные данные — пустой массив или значение NULL, вернуть пустой массив.

Example:
[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15] -> [10, -65]
*/

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	fmt.Println(CountPositivesSumNegatives(numbers)) // [10, -65]
}

// CountPositivesSumNegatives возвращает массив, где первый элемент — количество положительных чисел, а второй — сумму отрицательных чисел.
// time: O(n), space: O(1), n - длина массива
func CountPositivesSumNegatives(numbers []int) []int {
	if len(numbers) == 0 {
		return []int{}
	}

	positiveNumbers := 0
	sumNegativeNumbers := 0
	for _, num := range numbers {
		if num > 0 {
			positiveNumbers++
		} else if num < 0 {
			sumNegativeNumbers += num
		}
	}

	return []int{positiveNumbers, sumNegativeNumbers}
}
