package main

import "fmt"

/* Is this a triangle?
https://www.codewars.com/kata/56606694ec01347ce800001b/train/go

Реализуйте функцию, принимающую три целых числа: a, b, c. Функция должна возвращать значение true, если треугольник
можно построить со сторонами заданной длины, и false в любом другом случае.

(В этом случае для принятия функции площадь поверхности всех треугольников должна быть больше 0).
*/

func main() {
	fmt.Println(IsTriangle(1, 2, 2)) // true
	fmt.Println(IsTriangle(5, 1, 2)) // false
}

// IsTriangle принимает три целых числа a, b, c и возвращает значение true, если треугольник можно построить.
// time: O(1), space: O(1)
func IsTriangle(a, b, c int) bool {
	return a > 0 && b > 0 && c > 0 &&
		a+b > c && a+c > b && b+c > a
}
