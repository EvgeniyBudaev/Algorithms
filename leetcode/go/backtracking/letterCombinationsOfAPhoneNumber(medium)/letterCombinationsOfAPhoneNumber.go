package main

import "fmt"

/*
https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/

Учитывая строку, содержащую цифры от 2 до 9 включительно, верните все возможные комбинации букв, которые может
представлять число. Верните ответ в любом порядке.
Соответствие цифр буквам (как на кнопках телефона) приведено ниже. Обратите внимание, что 1 не соответствует
никаким буквам.
*/

/*
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
*/

func main() {
	fmt.Println(letterCombinations("23")) // ["ad","ae","af","bd","be","bf","cd","ce","cf"]
}

// letterCombinations генерирует все возможные комбинации букв, соответствующих цифрам
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	// Создаем маппинг цифр к соответствующим буквам
	digitToLetters := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var result []string

	// Внутренняя рекурсивная функция для генерации комбинаций
	var backtrack func(int, string)
	backtrack = func(index int, combination string) {
		// Если прошли все цифры, добавляем комбинацию в результат
		if index == len(digits) {
			result = append(result, combination)
			return
		}

		// Получаем буквы для текущей цифры
		currentDigit := digits[index]
		letters := digitToLetters[currentDigit]

		// Добавляем каждую возможную букву и рекурсивно продолжаем
		for i := 0; i < len(letters); i++ {
			backtrack(index+1, combination+string(letters[i]))
		}
	}

	backtrack(0, "")
	return result
}
