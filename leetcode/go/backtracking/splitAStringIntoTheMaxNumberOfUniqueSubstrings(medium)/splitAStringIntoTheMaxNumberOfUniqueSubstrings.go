package main

import "fmt"

/* https://leetcode.com/problems/split-a-string-into-the-max-number-of-unique-substrings/description/

Учитывая строку s, верните максимальное количество уникальных подстрок, на которые можно разбить данную строку.
Вы можете разделить строку на любой список непустых подстрок, где объединение подстрок образует исходную строку.
Однако вам необходимо разделить подстроки так, чтобы все они были уникальными.
Подстрока — это непрерывная последовательность символов внутри строки.

Input: s = "ababccc"
Output: 5
Пояснение: Один из способов максимального разделения — это ['a', 'b', 'ab', 'c', 'cc'].
Разделение типа ['a', 'b', 'a', 'b', 'c', 'cc'] недопустимо, поскольку у вас есть 'a' и 'b' несколько раз.

Input: s = "aba"
Output: 2
Пояснение: Один из способов максимального разделения — ['a', 'ba'].

Input: s = "aa"
Output: 1
Объяснение: Дальнейшее разделение строки невозможно.
*/

func main() {
	fmt.Println(maxUniqueSplit("ababccc")) // 5
}

// maxUniqueSplit возвращает максимальное количество уникальных подстрок,
// на которые можно разбить строку s
func maxUniqueSplit(s string) int {
	return getMax(0, s, make(map[string]bool))
}

// Вспомогательная рекурсивная функция для поиска максимального разбиения
func getMax(start int, s string, set map[string]bool) int {
	// Базовый случай: если дошли до конца строки, возвращаем количество уникальных подстрок
	if start == len(s) {
		return len(set)
	}

	maxCount := 0

	// Перебираем все возможные подстроки, начиная с текущей позиции
	for i := start + 1; i <= len(s); i++ {
		sub := s[start:i]

		// Если подстрока еще не встречалась
		if !set[sub] {
			set[sub] = true // Добавляем подстроку в множество

			// Рекурсивно ищем максимальное разбиение для оставшейся части строки
			currentCount := getMax(i, s, set)
			if currentCount > maxCount {
				maxCount = currentCount
			}

			delete(set, sub) // Удаляем подстроку (backtracking)
		}
	}

	return maxCount
}
