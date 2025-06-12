package main

import (
	"fmt"
)

/* Quarter of the year
https://www.codewars.com/kata/5ce9c1000bab0b001134f5af/train/go

Если месяц задан как целое число от 1 до 12, верните, к какому кварталу года он принадлежит, как целое число.

Например: месяц 2 (февраль) является частью первого квартала; месяц 6 (июнь) является частью второго квартала;
а месяц 11 (ноябрь) является частью четвертого квартала.

Ограничение:
1 <= месяц <= 12
*/

func main() {
	fmt.Println(QuarterOf(3))  // 1
	fmt.Println(QuarterOf(8))  // 3
	fmt.Println(QuarterOf(11)) // 4
}

// QuarterOf определяет, в каком квартале года находится месяц.
// time: O(1), space: O(1)
func QuarterOf(month int) int {
	return (month + 2) / 3
}
