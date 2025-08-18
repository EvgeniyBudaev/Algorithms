package main

import (
	"fmt"
)

/* Cat years, Dog years
https://www.codewars.com/kata/5a6663e9fd56cb5ab800008b/train/go

У меня есть кошка и собака. Я взял их одновременно, когда они были котятами и щенками. Это было humanYears лет назад.
Вычислите их текущий возраст в формате [humanYears,catYears,dogYears]

ПРИМЕЧАНИЯ:
humanYears >= 1
humanYears — только целые числа.
Cat Years
15 cat years for first year
+9 cat years for second year
+4 cat years for each after year.
Cat Years
15 dog years for first year
+9 dog years for second year
+5 dog years for each after year.
*/

func main() {
	fmt.Println(CalculateYears(1)) // [1,15,15]
}

// CalculateYears возвращает текущий возраст кошки и собаки.
// time: O(n), space: O(1), n - количество лет
func CalculateYears(years int) (result [3]int) {
	humanYears := years
	catYears := 0
	dogYears := 0

	for i := 1; i <= years; i++ {
		switch i {
		case 1:
			catYears += 15
			dogYears += 15
		case 2:
			catYears += 9
			dogYears += 9
		default:
			catYears += 4
			dogYears += 5
		}
	}

	result = [3]int{humanYears, catYears, dogYears}
	return
}
