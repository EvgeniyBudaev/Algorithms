package main

import "fmt"

/* https://github.com/EvgeniyBudaev/doocs_leetcode/tree/main/solution/0100-0199/0159.Longest%20Substring%20with%20At%20Most%20Two%20Distinct%20Characters

Учитывая строку s и целое число k, верните длину самой длинной подстроки s, содержащей не более k различных символов.

Input: s = "eceba", k = 2
Output: 3
Объяснение: Это подстрока «ece» длиной 3.

Input: s = "aa", k = 1
Output: 2
Объяснение: Это подстрока «aa» длиной 2.
*/

func main() {
	fmt.Println(lengthOfLongestSubstringKDistinct("eceba", 2)) // Вывод: 3
}

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	charMap := make(map[byte]int) // Хранит частоту символов в текущем окне
	maxLength := 0                // Максимальная длина подстроки
	left := 0                     // Указатель на начало окна

	for right := 0; right < len(s); right++ {
		currentChar := s[right]
		charMap[currentChar]++ // Увеличиваем частоту текущего символа
		// Если количество уникальных символов превышает k, сдвигаем left
		for len(charMap) > k {
			leftChar := s[left]
			charMap[leftChar]-- // Уменьшаем частоту символа слева
			if charMap[leftChar] == 0 {
				delete(charMap, leftChar) // Удаляем символ, если его частота стала 0
			}
			left++ // Сдвигаем left
		}
		// Обновляем максимальную длину подстроки
		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}
