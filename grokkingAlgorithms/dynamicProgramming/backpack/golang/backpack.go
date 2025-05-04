package main

import "fmt"

// Item представляет структуру товара
type Item struct {
	Name   string
	Weight int
	Price  int
}

// Задача о рюкзаке
func main() {
	items := []Item{
		{"tapeRecorder", 4, 3000},
		{"laptop", 3, 2000},
		{"guitar", 1, 1500},
	}
	capacity := 4 // вместимость рюкзака

	result := GetGoodsWithMaxPrice(items, capacity)
	fmt.Printf("Максимальная стоимость: %d\n", result.MaxPrice)
	fmt.Printf("Выбранные предметы: %v\n", result.SelectedItems)
}

// KnapsackResult содержит результат работы алгоритма
type KnapsackResult struct {
	MaxPrice      int
	SelectedItems []string
}

// GetGoodsWithMaxPrice реализует решение задачи о рюкзаке
func GetGoodsWithMaxPrice(items []Item, capacity int) KnapsackResult {
	// Инициализация таблицы динамического программирования
	dp := make([][]int, len(items)+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	// Заполнение таблицы
	for i := 1; i <= len(items); i++ {
		for j := 1; j <= capacity; j++ {
			item := items[i-1] // получаем товар из списка товаров
			// если текущий товар помещается в рюкзак, то выбираем максимальную стоимость между текущим товаром и предыдущими
			if item.Weight <= j {
				// если стоимость текущего товара + стоимость предыдущих товаров, которые помещаются в оставшуюся вместимость,
				// больше стоимости предыдущих товаров, которые помещаются в рюкзак, то выбираем текущий товар
				if item.Price+dp[i-1][j-item.Weight] > dp[i-1][j] {
					dp[i][j] = item.Price + dp[i-1][j-item.Weight]
					// если стоимость предыдущих товаров больше, то выбираем их
				} else {
					dp[i][j] = dp[i-1][j]
				}
				// если текущий товар не помещается в рюкзак, то выбираем предыдущие товары
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	// Восстановление выбранных предметов
	var selectedItems []string
	i := len(items) // количество товаров
	j := capacity   // вместимость рюкзака

	// Ищем выбранные предметы, начиная с конца
	for i > 0 && j > 0 {
		// если стоимость текущего товара + стоимость предыдущих товаров, которые помещаются в оставшуюся вместимость,
		// больше стоимости предыдущих товаров, которые помещаются в рюкзак, то выбираем текущий товар
		if dp[i][j] != dp[i-1][j] {
			selectedItems = append(selectedItems, items[i-1].Name)
			j -= items[i-1].Weight
		}
		i--
	}

	// Переворачиваем список, так как мы шли с конца
	for k, l := 0, len(selectedItems)-1; k < l; k, l = k+1, l-1 {
		selectedItems[k], selectedItems[l] = selectedItems[l], selectedItems[k]
	}

	return KnapsackResult{
		MaxPrice:      dp[len(items)][capacity],
		SelectedItems: selectedItems,
	}
}
