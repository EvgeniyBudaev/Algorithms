package main

import "fmt"

/* https://leetcode.com/problems/find-all-anagrams-in-a-string/description/

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

func findAnagrams(s string, p string) []int {
	arr := []int{}
	charMap := make(map[byte]int)

	// Заполняем map для строки p
	for i := 0; i < len(p); i++ {
		charMap[p[i]]++
	}

	left, right, count := 0, 0, len(p)

	for right < len(s) {
		// Если символ s[right] есть в charMap, уменьшаем count
		if charMap[s[right]] > 0 {
			count--
		}
		// Уменьшаем значение в charMap для текущего символа
		charMap[s[right]]--
		right++
		fmt.Println("count:", count)
		// Если count == 0, значит, нашли анаграмму
		if count == 0 {
			arr = append(arr, left)
		}

		// Если окно достигло размера p, сдвигаем left
		if right-left == len(p) {
			// Если символ s[left] был в charMap, увеличиваем count
			if charMap[s[left]] >= 0 {
				count++
			}
			// Восстанавливаем значение в charMap для символа s[left]
			charMap[s[left]]++
			left++
		}
	}

	return arr
}
