package main

import (
	"fmt"
	"sort"
	"strings"
)

/* 451. Sort Characters By Frequency
https://leetcode.com/problems/sort-characters-by-frequency/description/

Дана строка s, отсортируйте ее в порядке убывания на основе частоты символов. Частота символа — это количество раз,
которое он появляется в строке.
Верните отсортированную строку. Если есть несколько ответов, верните любой из них.

Input: s = "tree"
Output: "eert"
Объяснение: «e» встречается дважды, а «r» и «t» встречаются по одному разу.
Поэтому «e» должно встречаться перед «r» и «t». Поэтому «eetr» также является допустимым ответом.
*/

func main() {
	fmt.Println(frequencySort("tree")) // "eetr"
}

// frequencySort - Функция для сортировки символов по частоте в порядке убывания.
// time: O(n*log(n)), space: O(n)
func frequencySort(s string) string {
	// Подсчитываем частоту символов
	counter := make(map[rune]int)
	for _, char := range s {
		counter[char]++
	}

	// Создаем срез пар «символ-частота»
	type charFreq struct {
		char rune // Символ
		freq int  // Частота
	}
	pq := make([]charFreq, 0, len(counter)) // Создаем срез с начальной емкостью
	for char, freq := range counter {       // Проходим по всем символам и их частотам
		pq = append(pq, charFreq{char, freq}) // Добавляем пару «символ-частота» в срез
	}

	// Сортируем по частоте в порядке убывания
	sort.Slice(pq, func(i, j int) bool {
		return pq[i].freq > pq[j].freq
	})

	// Строим строку результата
	var result strings.Builder // Используем билдер для эффективного построения строки
	for _, cf := range pq {
		result.WriteString(strings.Repeat(string(cf.char), cf.freq))
	}

	return result.String() // Возвращаем построенную строку результат
}
