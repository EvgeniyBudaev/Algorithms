package main

import (
	"fmt"
)

/* 28. Find the Index of the First Occurrence in a String
https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/

Для двух строк «needle» и «haystack» вернуть индекс первого вхождения «needle» в «haystack» или -1,
если «needle» не является частью «haystack».

Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Объяснение: «sad» встречается в индексах 0 и 6. Первое вхождение находится в индексе 0, поэтому мы возвращаем 0.
*/

func main() {
	haystack := "sadbutsad"
	needle := "sad"
	fmt.Println(strStr(haystack, needle)) // 0
}

// strStr возвращает индекс первого вхождения needle в haystack.
// time: O(n*m), где n - длина haystack, m - длина needle, space: O(1)
func strStr(haystack string, needle string) int {
	// Если needle пустой, то возвращаем 0
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ { // Проходим по строке
		// Проверяем, что в строке есть needle
		if haystack[i:i+len(needle)] == needle {
			return i // Возвращаем индекс
		}
	}

	return -1 // Если needle не найдено, возвращаем -1
}
