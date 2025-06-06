package main

import "fmt"

/* 438. Find All Anagrams in a String
https://leetcode.com/problems/find-all-anagrams-in-a-string/description/

Учитывая две строки s и p, верните массив всех начальных индексов анаграмм p в s. Вы можете вернуть ответ в любом
порядке.
Анаграмма — это слово или фраза, образованная путем перестановки букв другого слова или фразы, обычно с использованием
всех исходных букв ровно один раз.

Input: s = "cbaebabacd", p = "abc"
Output: [0,6]
Объяснение:
Подстрока с начальным индексом = 0 — это «cba», которая является анаграммой слова «abc».
Подстрока с начальным индексом = 6 — это «bac», которая является анаграммой слова «abc».
*/

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc")) // [0,6]
}

// findAnagrams находит все анаграммы строки p в строке s и возвращает их начальные индексы.
// time: O(n), space: O(1)
func findAnagrams(s string, p string) []int {
	var result []int              // Список индексов анаграмм
	charMap := make(map[byte]int) // Словарь для подсчета символов в строке p

	// Заполняем map для строки p
	for i := 0; i < len(p); i++ {
		charMap[p[i]]++ // Увеличиваем значение в charMap для текущего символа
	}

	left, right, count := 0, 0, len(p) // left - левый указатель, right - правый указатель, count - количество символов в p

	for right < len(s) { // Перебираем все символы в строке s
		// Если символ s[right] есть в charMap, уменьшаем count
		if charMap[s[right]] > 0 {
			count--
		}
		charMap[s[right]]-- // Уменьшаем значение в charMap для текущего символа
		right++             // Перемещаем правый указатель

		// Если count == 0, значит, нашли анаграмму
		if count == 0 {
			result = append(result, left) // Добавляем индекс анаграммы в result
		}

		// Если окно достигло размера p, сдвигаем left
		if right-left == len(p) {
			// Если символ s[left] был в charMap, увеличиваем count
			if charMap[s[left]] >= 0 {
				count++ // Увеличиваем count, так как мы убрали символ из подстроки
			}
			// Восстанавливаем значение в charMap для символа s[left]
			charMap[s[left]]++
			left++ // Перемещаем левый указатель
		}
	}

	return result // Возвращает список индексов анаграмм
}
