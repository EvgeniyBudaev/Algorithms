package main

import "fmt"

/* 424. Longest Repeating Character Replacement
https://leetcode.com/problems/longest-repeating-character-replacement/description/

Вам дана строка s и целое число k. Вы можете выбрать любой символ строки и заменить его на любой другой символ
английского верхнего регистра. Эту операцию можно выполнить не более k раз.

Верните длину самой длинной подстроки, содержащей ту же букву, которую вы можете получить после выполнения
вышеуказанных операций.

Input: s = "ABAB", k = 2
Output: 4
Пояснение: Замените две буквы «А» двумя буквами «B» или наоборот.

Input: s = "AABABBA", k = 1
Output: 4
Пояснение: Замените букву «A» посередине на «B» и получите «AABBBBA».
Подстрока «BBBB» имеет самую длинную повторяющуюся букву — 4.
Могут существовать и другие способы получить этот ответ.
*/

func main() {
	fmt.Println(characterReplacement("ABAB", 2)) // 4
}

// characterReplacement находит длину самой длинной подстроки, в которой можно сделать не более k замен символов.
// time: O(n), space: O(1)
func characterReplacement(s string, k int) int {
	charMap := make(map[byte]int) // Хранит количество каждого символа в текущем окне
	result, maxFreq := 0, 0       // result - максимальная длина, maxFreq - максимальная частота символа в окне
	left := 0                     // Указатель на начало окна

	for right := 0; right < len(s); right++ {
		currentChar := s[right] // Текущий символ
		charMap[currentChar]++  // Увеличиваем количество текущего символа

		if charMap[currentChar] > maxFreq { // Обновляем максимальную частоту символа в окне
			maxFreq = charMap[currentChar]
		}

		window := right - left + 1 // Размер текущего окна

		if window-maxFreq > k { // Если размер окна превышает maxFreq + k, сдвигаем left
			charMap[s[left]]-- // Уменьшаем количество символа слева
			if charMap[s[left]] == 0 {
				delete(charMap, s[left]) // Удаляем символ, если его количество становится 0
			}
			left++ // Сдвигаем left
		}

		result = max(result, window) // Обновляем максимальную длину окна
	}

	return result // Возвращаем максимальную длину
}
