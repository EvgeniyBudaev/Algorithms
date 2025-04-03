package main

import "fmt"

/* https://leetcode.com/problems/palindrome-partitioning/description/

Дана строка s, раздел s такой, что каждый подстрока раздела представляет собой палиндром.
Вернуть все возможные палиндромные разбиения s.

Input: s = "aab"
Output: [["a","a","b"],["aa","b"]]

Input: s = "a"
Output: [["a"]]
*/

func main() {
	fmt.Println(partition("aab")) // [["a","a","b"],["aa","b"]]
}

// partition возвращает все возможные разбиения строки на палиндромы
func partition(s string) [][]string {
	var result [][]string

	// Внутренняя функция для проверки палиндрома
	isPalindrome := func(str string, left, right int) bool {
		for left < right {
			if str[left] != str[right] {
				return false
			}
			left++
			right--
		}
		return true
	}

	// Рекурсивная функция для генерации разбиений
	var backtrack func(int, []string)
	backtrack = func(start int, path []string) {
		// Если дошли до конца строки, добавляем текущее разбиение в результат
		if start == len(s) {
			// Создаем копию пути, чтобы избежать изменений
			temp := make([]string, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		// Перебираем все возможные подстроки, начиная с текущей позиции
		for end := start; end < len(s); end++ {
			// Если подстрока - палиндром, продолжаем рекурсивно
			if isPalindrome(s, start, end) {
				// Добавляем палиндром в текущий путь
				path = append(path, s[start:end+1])
				// Рекурсивно обрабатываем оставшуюся часть строки
				backtrack(end+1, path)
				// Удаляем последний элемент (backtracking)
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(0, []string{})
	return result
}
