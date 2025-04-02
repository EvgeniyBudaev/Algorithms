package main

import (
	"fmt"
	"sort"
)

/* https://leetcode.com/problems/top-k-frequent-elements/description/

Учитывая целочисленный массив nums и целое число k, верните k наиболее часто встречающихся элементов. Вы можете вернуть
ответ в любом порядке.

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

Input: nums = [1], k = 1
Output: [1]
*/

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)) // [1 2]
}

// topKFrequent возвращает k наиболее часто встречающихся чисел
func topKFrequent(nums []int, k int) []int {
	// Создаем map для подсчета частоты каждого числа
	frequency := make(map[int]int)

	// Заполняем map: ключ - число, значение - частота встречаемости
	for _, num := range nums {
		frequency[num]++
	}

	// Создаем структуру для хранения чисел и их частот
	type numFreq struct {
		num  int
		freq int
	}

	// Преобразуем map в слайс структур для сортировки
	var numFreqs []numFreq
	for num, freq := range frequency {
		numFreqs = append(numFreqs, numFreq{num, freq})
	}

	// Сортируем по убыванию частоты
	sort.Slice(numFreqs, func(i, j int) bool {
		return numFreqs[i].freq > numFreqs[j].freq
	})

	// Выбираем k самых частых чисел
	result := make([]int, 0, k)
	for i := 0; i < k && i < len(numFreqs); i++ {
		result = append(result, numFreqs[i].num)
	}

	return result
}
