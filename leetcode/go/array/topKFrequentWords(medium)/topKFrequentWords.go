package main

import (
	"fmt"
	"sort"
	"strings"
)

/* 692. Top K Frequent Words
https://leetcode.com/problems/top-k-frequent-words/description/

Учитывая массив строковых слов и целое число k, верните k наиболее часто встречающихся строк.
Возвращает ответ, отсортированный по частоте от самого высокого до самого низкого. Отсортируйте слова с одинаковой
частотой по их лексикографическому порядку.

Input: words = ["i","love","leetcode","i","love","coding"], k = 2
Output: ["i","love"]
Пояснение: «i» и «love» — два наиболее часто встречающихся слова.
Обратите внимание, что «i» стоит перед словом «love» из-за более низкого алфавитного порядка.
*/

func main() {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	fmt.Println(topKFrequent(words, 2)) // ["i", "love"]
}

func topKFrequent(words []string, k int) []string {
	// Создаем map для подсчета частоты слов
	freqMap := make(map[string]int)
	for _, word := range words {
		freqMap[word]++
	}

	// Создаем slice для хранения пар (слово, частота)
	type wordFreq struct {
		word  string
		count int
	}
	var wordFreqs []wordFreq
	for word, count := range freqMap {
		wordFreqs = append(wordFreqs, wordFreq{word, count})
	}

	// Сортируем slice по частоте и лексикографическому порядку
	sort.Slice(wordFreqs, func(i, j int) bool {
		if wordFreqs[i].count != wordFreqs[j].count {
			return wordFreqs[i].count > wordFreqs[j].count
		}
		return strings.Compare(wordFreqs[i].word, wordFreqs[j].word) < 0
	})

	// Выбираем первые k элементов
	result := make([]string, 0, k)
	for i := 0; i < k && i < len(wordFreqs); i++ {
		result = append(result, wordFreqs[i].word)
	}

	return result
}
