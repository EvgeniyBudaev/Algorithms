package main

import (
	"fmt"
	"sort"
)

/* 1196 - How Many Apples Can You Put into the Basket

У вас есть несколько яблок, где arr[i] — вес i-го яблока.
У вас также есть корзина, способная выдержать до 5000 единиц веса.
Верните максимальное количество яблок, которое вы можете положить в корзину.

Input: arr = [100,200,150,1000]
Output: 4
Пояснение: Все 4 яблока можно нести в корзине, так как их сумма весов равна 1450.

Input: arr = [900,950,800,1000,700,800]
Output: 5
Пояснение: Сумма весов 6 яблок превышает 5000, поэтому мы выбираем любые 5 из них.
*/

func main() {
	arr := []int{900, 950, 800, 1000, 700, 800}
	fmt.Println(maxNumberOfApples(arr)) // 5
}

// maxNumberOfApples возвращает максимальное количество яблок, которое можно поместить в корзину без превышения веса.
// time: O(n), space: O(1)
func maxNumberOfApples(appleWeights []int) int {
	totalWeight := 0        // Общий вес яблок
	limit := 5000           // Максимальный вес корзины
	sort.Ints(appleWeights) // Сортируем яблоки по весу (от легких к тяжелым)

	for i, weight := range appleWeights { // Перебираем все яблоки
		totalWeight += weight // Добавляем вес текущего яблока к общему весу
		// Если превысили лимит в 5000, возвращаем текущее количество
		if totalWeight > limit {
			return i
		}
	}

	return len(appleWeights) // Если все яблоки поместились, возвращаем общее количество
}
