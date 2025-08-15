package main

import (
	"fmt"
)

/* All Inclusive?
https://www.codewars.com/kata/5700c9acc1555755be00027e/train/go

Входные данные:

строка strng
массив строк arr
Выходные данные функции contain_all_rots(strng, arr) (или containAllRots или contain-all-rots):

логическое значение true, если все вращения strng включены в arr
в противном случае false

Примеры:
contain_all_rots(
"bsjq", ["bsjq", "qbsj", "sjqb", "twZNsslC", "jqbs"]) -> true
*/

func main() {
	arr1 := []string{"bsjq", "qbsj", "sjqb", "twZNsslC", "jqbs"}
	fmt.Println(ContainAllRots("bsjq", arr1)) // true

	arr2 := []string{"TzYxlgfnhf", "yqVAuoLjMLy", "BhRXjYA", "YABhRXj", "hRXjYAB", "jYABhRX", "XjYABhR", "ABhRXjY"}
	fmt.Println(ContainAllRots("XjYABhR", arr2)) // false
}

// ContainAllRots возвращает true, если все вращения strng включены в arr, иначе false.
// time: O(n), space: O(n), n - длина строки
func ContainAllRots(strng string, arr []string) bool {
	rotations := make(map[string]struct{})
	for i := range strng {
		rotations[strng[i:]+strng[:i]] = struct{}{}
	}

	for _, v := range arr {
		delete(rotations, v)
	}

	return len(rotations) == 0
}
