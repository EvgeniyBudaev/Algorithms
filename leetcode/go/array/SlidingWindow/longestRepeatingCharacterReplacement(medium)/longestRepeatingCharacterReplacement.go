package main

import "fmt"

/* https://leetcode.com/problems/longest-repeating-character-replacement/description/

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

func characterReplacement(s string, k int) int {
	charMap := make(map[byte]int) // Хранит количество каждого символа в текущем окне
	ans, maxFreq := 0, 0          // ans - максимальная длина, maxFreq - максимальная частота символа в окне
	left := 0                     // Указатель на начало окна

	for right := 0; right < len(s); right++ {
		currentChar := s[right]
		charMap[currentChar]++ // Увеличиваем количество текущего символа

		// Обновляем максимальную частоту символа в окне
		if charMap[currentChar] > maxFreq {
			maxFreq = charMap[currentChar]
		}

		window := right - left + 1
		// Если размер окна превышает maxFreq + k, сдвигаем left
		if window-maxFreq > k {
			charMap[s[left]]-- // Уменьшаем количество символа слева
			left++             // Сдвигаем left
		}

		// Обновляем максимальную длину окна
		ans = max(ans, window)
	}

	return ans
}
