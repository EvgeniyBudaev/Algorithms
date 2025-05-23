package main

import (
	"fmt"
)

/* 387. First Unique Character in a String
https://leetcode.com/problems/first-unique-character-in-a-string/description/

Дана строка s, найти в ней первый неповторяющийся символ и вернуть его индекс. Если он не существует, вернуть -1.

Input: s = "leetcode"
Output: 0
Пояснение: Символ «l» в индексе 0 — это первый символ, который не встречается ни в каком другом индексе.
*/

func main() {
	fmt.Println(firstUniqChar("leetcode")) // 0
}

// firstUniqChar возвращает индекс первого уникального символа в строке s.
// time: O(n), space: O(n)
func firstUniqChar(s string) int {
	charCount := make(map[rune]int) // Создаем мап для подсчета количества вхождений каждого символа

	// Первый проход: подсчитываем количество каждого символа
	for _, char := range s {
		charCount[char]++
	}

	// Второй проход: ищем первый символ с количеством 1
	for i, char := range s {
		if charCount[char] == 1 {
			return i
		}
	}

	return -1 // Если не нашли уникальных символов
}
