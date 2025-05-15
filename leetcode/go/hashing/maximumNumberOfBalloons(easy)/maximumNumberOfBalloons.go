package main

import (
	"fmt"
	"math"
)

/* 1189. Maximum Number of Balloons
https://leetcode.com/problems/maximum-number-of-balloons/description/

Учитывая строковый текст, вы хотите использовать символы текста для формирования как можно большего количества
экземпляров слова «balloon».
Каждый символ в тексте можно использовать не более одного раза. Верните максимальное количество экземпляров,
которые можно сформировать.

Input: text = "nlaebolko"
Output: 1

Input: text = "loonbalxballpoon"
Output: 2

Input: text = "leetcode"
Output: 0
*/

func main() {
	fmt.Println(maxNumberOfBalloons("nlaebolko")) // 1
}

// maxNumberOfBalloons - подсчитывает, сколько раз можно составить слово "balloon" из букв строки.
// time: O(n), space: O(1)
func maxNumberOfBalloons(text string) int {
	// Создаем мапу для подсчета нужных букв
	counts := map[rune]float64{
		'b': 0,
		'a': 0,
		'l': 0,
		'o': 0,
		'n': 0,
	}

	// Подсчитываем только те буквы, которые входят в слово "balloon"
	for _, char := range text {
		if _, exists := counts[char]; exists { // Если буква есть в мапе, то увеличиваем ее счетчик
			counts[char]++
		}
	}

	// Для букв 'l' и 'o' нужно 2 экземпляра на каждое слово "balloon",
	// поэтому делим их количество на 2
	counts['l'] /= 2
	counts['o'] /= 2

	// Находим минимальное значение среди всех требуемых букв
	minValue := math.Inf(1) // Начинаем с положительной бесконечности
	for _, count := range counts {
		minValue = min(minValue, count) // Находим минимальное значение
	}

	// Возвращаем целую часть минимального значения (округляем вниз)
	return int(math.Floor(minValue))
}
