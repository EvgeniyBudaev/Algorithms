package main

import (
	"fmt"
)

/* Third Angle of a Triangle
https://www.codewars.com/kata/5a023c426975981341000014/train/go

Даны два внутренних угла треугольника (в градусах).
Напишите функцию, возвращающую третий угол.
Примечание: будут проверяться только положительные целые числа.
*/

func main() {
	fmt.Println(OtherAngle(30, 60)) // 90
}

// OtherAngle возвращает третий угол треугольника.
// time: O(1), space: O(1)
func OtherAngle(a int, b int) int {
	return 180 - (a + b)
}
