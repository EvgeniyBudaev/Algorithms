package main

import "fmt"

/* 567. Permutation in String
https://leetcode.com/problems/permutation-in-string/description/

Учитывая две строки s1 и s2, верните true, если s2 содержит перестановку s1, или false в противном случае.
Другими словами, верните true, если одна из перестановок s1 является подстрокой s2.

Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Пояснение: s2 содержит одну перестановку s1 («ba»).

Input: s1 = "ab", s2 = "eidboaoo"
Output: false
*/

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo")) // true
}

// checkInclusion - проверяет, содержит ли строка s2 перестановку строки s1.
// time: O(n), space: O(1)
func checkInclusion(s1 string, s2 string) bool {
	// Если длина s1 больше длины s2, то s1 не может быть перестановкой s2
	if len(s1) > len(s2) {
		return false
	}

	// Создаем карту для подсчета символов в s1
	neededChar := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		neededChar[s1[i]]++
	}

	left, right, requiredLength := 0, 0, len(s1) // Инициализируем левую и правую границы окна, и счетчик необходимых символов

	for right < len(s2) { // Проходим по s2
		currentChar := s2[right] // Текущий символ в s2
		// Если текущий символ найден в s1, уменьшаем счетчик необходимых символов
		if neededChar[currentChar] > 0 {
			requiredLength-- // Уменьшаем счетчик необходимых символов
		}
		neededChar[currentChar]-- // Уменьшаем счетчик текущего символа
		right++                   // Сдвигаем правую границу окна

		// Если все символы s1 найдены в текущем окне
		if requiredLength == 0 {
			return true
		}

		// Если окно достигло длины s1, сдвигаем левую границу
		if right-left == len(s1) {
			// Если символ на левой границе был найден в s1, увеличиваем счетчик необходимых символов
			if neededChar[s2[left]] >= 0 {
				requiredLength++ // Увеличиваем счетчик необходимых символов
			}
			neededChar[s2[left]]++ // Восстанавливаем счетчик текущего символа
			left++                 // Сдвигаем левую границу окна
		}
	}

	return false // Если перестановка не найдена
}
