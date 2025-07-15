package main

import (
	"fmt"
	"sort"
)

/* Sort array by string length
https://www.codewars.com/kata/57ea5b0b75ae11d1e800006c/train/go

Напишите функцию, которая принимает массив строк в качестве аргумента и возвращает отсортированный массив, содержащий
те же строки, упорядоченные от самой короткой к самой длинной.
*/

func main() {
	arr := []string{"Telescopes", "Glasses", "Eyes", "Monocles"}
	fmt.Println(SortByLength(arr)) // ["Eyes", "Glasses", "Monocles", "Telescopes"]
}

// SortByLength возвращает отсортированный массив, содержащий те же строки,
// упорядоченные от самой короткой к самой длинной.
// time: O(n^2), space: O(1), n - длина массива
func SortByLength(arr []string) []string {
	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i]) < len(arr[j])
	})
	return arr
}
