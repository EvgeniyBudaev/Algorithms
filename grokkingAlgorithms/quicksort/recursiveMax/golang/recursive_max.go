package main

import "fmt"

func main() {
	fmt.Println(recursiveMax([]int{1, 5, 10, 25, 16, 1})) // 25
}

// recursiveMax возвращает максимальное число из массива.
func recursiveMax(list []int) int {
	// В случае массива из двух элементов возвращаем большее.
	if len(list) == 2 {
		if list[0] > list[1] {
			return list[0]
		}
		return list[1]
	}

	// Находим максимальное число в подмассиве.
	subMax := recursiveMax(list[1:])
	// Если первый элемент больше максимального числа в подмассиве,
	// то возвращаем его, иначе возвращаем максимальное число из подмассива.
	if list[0] > subMax {
		return list[0]
	}
	// Возвращаем максимальное число из подмассива.
	return subMax
}
