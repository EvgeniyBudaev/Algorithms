package main

import (
	"fmt"
	"reflect"
)

/* https://leetcode.com/problems/flatten-deeply-nested-array/description/

Учитывая многомерный массив arr и глубину n, верните сглаженную версию этого массива.
Многомерный массив — это рекурсивная структура данных, содержащая целые числа или другие многомерные массивы.
Сглаженный массив — это версия этого массива, в которой некоторые или все подмассивы удалены и заменены фактическими
элементами этого подмассива. Эту операцию выравнивания следует выполнять только в том случае, если текущая глубина
вложенности меньше n. Глубина элементов в первом массиве считается равной 0.
Пожалуйста, решите эту проблему без встроенного метода Array.flat.

Input: arr = [1, 2, 3, [4, 5, 6], [7, 8, [9, 10, 11], 12], [13, 14, 15]], n = 0
Output: [1, 2, 3, [4, 5, 6], [7, 8, [9, 10, 11], 12], [13, 14, 15]]
Объяснение:
Передача глубины n=0 всегда приводит к исходному массиву. Это связано с тем, что наименьшая возможная глубина подмассива
(0) не меньше n=0. Таким образом, ни один подмассив не должен быть сглажен.

Input: arr = [1, 2, 3, [4, 5, 6], [7, 8, [9, 10, 11], 12], [13, 14, 15]]
n = 1
Output: [1, 2, 3, 4, 5, 6, 7, 8, [9, 10, 11], 12, 13, 14, 15]
Объяснение:
Все подмассивы, начинающиеся с 4, 7 и 13, выравниваются. Это связано с тем, что их глубина 0 меньше 1.
Однако [9, 10, 11] остается несглаженной, поскольку ее глубина равна 1.

Input: arr = [[1, 2, 3], [4, 5, 6], [7, 8, [9, 10, 11], 12], [13, 14, 15]]
n = 2
Output: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15]
Объяснение:
Максимальная глубина любого подмассива равна 1. Таким образом, все они сглажены.
*/

func main() {
	arr := []interface{}{1, 2, 3, []interface{}{4, 5, 6}, []interface{}{7, 8, []interface{}{9, 10, 11}, 12}, []interface{}{13, 14, 15}}
	n := 0
	fmt.Println(flat(arr, n)) // [1, 2, 3, [4, 5, 6], [7, 8, [9, 10, 11], 12], [13, 14, 15]]
}

func flat(arr []interface{}, n int) []interface{} {
	res := []interface{}{}

	var flatten func(arr []interface{}, depth int)
	flatten = func(arr []interface{}, depth int) {
		for _, item := range arr {
			if reflect.TypeOf(item) == reflect.TypeOf([]interface{}{}) && depth < n {
				// Если элемент является массивом и глубина меньше n, рекурсивно сглаживаем
				flatten(item.([]interface{}), depth+1)
			} else {
				// Иначе добавляем элемент в результат
				res = append(res, item)
			}
		}
	}

	flatten(arr, 0)
	return res
}
