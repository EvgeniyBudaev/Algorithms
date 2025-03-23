package main

import (
	"fmt"
)

/*
Учитывая два отсортированных массива целых чисел arr1 и arr2,
верните новый массив, который объединяет их оба и также отсортирован.
*/

func main() {
	arr1 := []int{1, 4, 7, 20}
	arr2 := []int{3, 5, 6}
	r := combine(arr1, arr2)
	fmt.Println(r) // [1 3 4 5 6 7 20]
}

func combine(arr1, arr2 []int) []int {
	m, n := len(arr1), len(arr2)
	p1 := m - 1
	p2 := n - 1
	arr1 = append(arr1, make([]int, n)...) // Расширяем arr1, чтобы он мог вместить элементы из arr2
	for i := m + n - 1; i >= 0; i-- {
		if p2 < 0 {
			break
		}
		if arr1[p1] > arr2[p2] {
			arr1[i] = arr1[p1]
			p1--
		} else {
			arr1[i] = arr2[p2]
			p2--
		}
	}
	return arr1
}

//func combine(arr1, arr2 []int) []int {
//	r := append(arr1, arr2...)
//	sort.Ints(r)
//	return r
//}
