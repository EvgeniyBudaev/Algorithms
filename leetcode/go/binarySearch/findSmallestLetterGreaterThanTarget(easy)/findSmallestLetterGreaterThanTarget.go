package main

import "fmt"

/* https://leetcode.com/problems/find-smallest-letter-greater-than-target/description/

Вам дан массив символов-букв, отсортированный в неубывающем порядке, и целевой символ. В буквах есть как минимум два
разных символа.
Верните наименьший буквенный символ, который лексикографически больше целевого. Если такого символа не существует,
вернуть первый символ буквами.

Input: letters = ["c","f","j"], target = "a"
Output: "c"
Пояснение: Самый маленький символ, который лексикографически больше буквы «a», — это «c».
*/

func main() {
	letters := []byte{'c', 'f', 'j'}
	fmt.Println(nextGreatestLetter(letters, 'a')) // 99
}

func nextGreatestLetter(letters []byte, target byte) byte {
	length := len(letters)

	// Обработка крайних случаев
	if target < letters[0] || target >= letters[length-1] {
		return letters[0]
	}

	// Бинарный поиск
	low, high := 0, length-1
	for low <= high {
		mid := (low + high) / 2
		guess := letters[mid]

		if guess > target && target >= letters[mid-1] {
			return guess
		}
		if guess > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return letters[0] // На случай, если ничего не найдено (по логике не должно выполняться)
}
