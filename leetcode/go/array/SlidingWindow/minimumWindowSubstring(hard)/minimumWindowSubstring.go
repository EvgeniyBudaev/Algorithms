package main

import "fmt"

/* https://leetcode.com/problems/minimum-window-substring/description/
solution https://leetcode.com/problems/minimum-window-substring/solutions/4673781/easy-sliding-window-solution-explained/

Учитывая две строки s и t длиной m и n соответственно, верните минимальную подстроку окна s, такую, что каждый символ
в t (включая дубликаты) включен в окно. Если такой подстроки нет, верните пустую строку «».

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Объяснение: Подстрока минимального окна "BANC" включает "A", "B" и "C" из строки t.

Input: s = "a", t = "a"
Output: "a"
Объяснение: Вся строка s представляет собой минимальное окно.

Input: s = "a", t = "aa"
Output: ""
Объяснение: В окно должны быть включены обе буквы 'a' из t.
Поскольку самое большое окно s имеет только одну букву «а», верните пустую строку.
*/

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC")) // "BANC"
}

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	// Создаем карту для подсчета символов в t
	charMap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		charMap[t[i]]++
	}

	left, minLen := 0, len(s)+1 // Инициализируем минимальную длину как максимально возможную
	ansLeft, ansRight := 0, 0   // Индексы для минимального окна
	required := len(t)          // Количество символов, которые нужно найти

	for right := 0; right < len(s); right++ {
		current := s[right]
		if charMap[current] > 0 {
			required-- // Уменьшаем количество требуемых символов
		}
		charMap[current]--

		// Когда все символы найдены, пытаемся сузить окно
		for required == 0 {
			// Обновляем минимальное окно
			if right-left+1 < minLen {
				minLen = right - left + 1
				ansLeft, ansRight = left, right
			}

			// Убираем левый символ из окна
			charMap[s[left]]++
			if charMap[s[left]] > 0 {
				required++ // Если символ был частью t, увеличиваем required
			}
			left++
		}
	}

	// Если минимальное окно не найдено, возвращаем пустую строку
	if minLen > len(s) {
		return ""
	}

	return s[ansLeft : ansRight+1]
}
