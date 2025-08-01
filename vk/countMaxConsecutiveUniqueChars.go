package main

/*
Нужно написать функцию, которая принимает на вход строку,
а на выходе возвращает для каждого уникального символа максимальное число его беспрерывных повторений.

Input: aafbaaaaffc
Output: a:4 b:1 f:2 c:1
*/

import "fmt"

func main() {
	fmt.Println(countMaxConsecutiveUniqueChars("aafbaaaaffc"))  // a:4 b:1 f:2 c:1
	fmt.Println(countMaxConsecutiveUniqueChars("aaaafbaaffc"))  // a:4 b:1 f:2 c:1
	fmt.Println(countMaxConsecutiveUniqueChars("saaaafbaaffc")) // a:4 b:1 f:2 c:1
}

// countMaxConsecutiveUniqueChars возвращает для каждого уникального символа максимальное число его беспрерывных повторений.
// time: O(n), space: O(n)
func countMaxConsecutiveUniqueChars(str string) map[byte]int {
	if len(str) == 0 {
		return nil
	}

	result := make(map[byte]int) // Мапа уникальных символов и количества их повторений
	currentChar := str[0]        // Текущий символ
	currentCount := 1            // Счетчик повторений текущего символа

	for i := 1; i < len(str); i++ { // Проходим по всей строке начиная со второго символа
		if str[i] == currentChar { // Если текущий символ равен предыдущему, то увеличиваем счетчик повторений
			currentCount++
			if currentCount > result[currentChar] {
				result[currentChar] = currentCount
			}
		} else { // Начинаем новую серию символов
			currentChar = str[i]
			currentCount = 1
			// Если символ встречается впервые, добавляем его
			if _, exists := result[currentChar]; !exists {
				result[currentChar] = currentCount
			}
		}
	}

	return result // Возвращаем мапу уникальных символов и количества их повторений
}

// countMaxConsecutiveUniqueChars возвращает для каждого уникального символа максимальное число его беспрерывных повторений.
// time: O(n), space: O(n)
//func countMaxConsecutiveUniqueChars(str string) map[byte]int {
//	charMap := make(map[byte]int) //Мапа уникальных символов и количества их повторений
//
//	count := 1 // Счетчик повторений
//	for i := len(str) - 1; i > 0; i-- {
//		if str[i] == str[i-1] { // Если текущий символ равен предыдущему, то увеличиваем счетчик повторений
//			count++
//			// Если текущий символ не равен предыдущему, то проверяем, не больше ли счетчика предыдущего символа,
//			// чем текущий символ, и если да, то обновляем значение в мапе
//		} else {
//			if charMap[str[i]] < count {
//				charMap[str[i]] = count
//			}
//			count = 1 // Обнуляем счетчик повторений
//		}
//	}
//
//	// Если последний символ не равен предыдущему, то проверяем, не больше ли счетчика последнего символа,
//	// чем текущий символ, и если да, то обновляем значение в мапе
//	if charMap[str[0]] < count {
//		charMap[str[0]] = count
//	}
//
//	return charMap // Возвращаем мапу уникальных символов и количества их повторений
//}
