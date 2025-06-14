package main

import (
	"fmt"
)

/* Sum of odd numbers
https://www.codewars.com/kata/55fd2d567d94ac3bc9000064/train/go

Дан треугольник из последовательных нечетных чисел:
             1
          3     5
       7     9    11
   13    15    17    19
21    23    25    27    29

Вычислите сумму чисел в n-й строке этого треугольника (начиная с индекса 1), например: (Вход --> Выход)
1 -->  1
2 --> 3 + 5 = 8
*/

func main() {
	fmt.Println(RowSumOddNumbers(2)) // 8
}

// RowSumOddNumbers возвращает сумму чисел в n-й строке этого треугольника.
// time: O(n), space: O(1)
func RowSumOddNumbers(n int) int {
	return n * n * n
}
