package main

import "fmt"

func main() {
	a := []int{1, 2, 2, 2, 3}
	b := []int{2}
	fmt.Println(ArrayDiff(a, b)) // [1,3]
}

// ArrayDiff возвращает элементы первого массива, которые не встречаются во втором.
// time: O(n), space: O(n)
func ArrayDiff(a, b []int) []int {
	count := 0 // Счетчик для количества элементов в результате

	// Создаем множество (map) для быстрой проверки наличия элемента в `b`
	bSet := make(map[int]bool)
	for _, num := range b {
		bSet[num] = true
	}

	for _, num := range a { // Проходим по всем элементам массива `a`
		if !bSet[num] { // Если элемент не найден в `b`, добавляем его в результат
			a[count] = num // Добавляем элемент в `a` на позицию `count`
			count++        // Увеличиваем счетчик
		}
	}

	a = a[:count] // Обрезаем срез `a` до фактического размера
	return a      // Возвращаем результат
}
