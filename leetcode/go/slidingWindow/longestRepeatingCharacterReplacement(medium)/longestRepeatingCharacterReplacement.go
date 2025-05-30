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
	charCount := make(map[byte]int)  // Мапа для подсчета количества каждого символа в подстроке
	left, result, maxFreq := 0, 0, 0 // Левый индекс подстроки, результат и максимальная частота в подстроке

	for right := 0; right < len(s); right++ { // Проходим по строке
		current := s[right]                        // Текущий символ
		charCount[current]++                       // Увеличиваем счетчик текущего символа
		maxFreq = max(maxFreq, charCount[current]) // Обновляем максимальную частоту
		window := right - left + 1                 // Текущая длина подстроки

		if window-maxFreq > k { // Если текущая подстрока не подходит по условию
			charCount[s[left]]-- // Уменьшаем счетчик левого символа
			left++               // Сдвигаем левый индекс
		}

		currentWindow := right - left + 1   // Текущая длина подстроки
		result = max(result, currentWindow) // Обновляем результат
	}

	return result // Возвращаем результат
}
