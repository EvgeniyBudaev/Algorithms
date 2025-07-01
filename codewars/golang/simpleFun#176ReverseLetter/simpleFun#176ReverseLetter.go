package main

import (
	"fmt"
	"unicode"
)

/* Simple Fun #176: Reverse Letter
https://www.codewars.com/kata/58b8c94b7df3f116eb00005b/train/go

Задача:
Дана строка str, переверните ее и исключите все неалфавитные символы.

Пример
Для str = "krishan" вывод должен быть "nahsirk".

Для str = "ultr53o?n" вывод должен быть "nortlu".

Ввод/вывод:
[ввод] строка str
Строка состоит из строчных латинских букв, цифр и символов.

[вывод] строка
*/

func main() {
	fmt.Println(ReverseLetters("krishan"))   // "nahsirk"
	fmt.Println(ReverseLetters("ultr53o?n")) // "nortlu"
}

// ReverseLetters возвращает строку, которая получена из исходной строки s путем переворота и
// исключения всех неалфавитных символов.
// time: O(n), space: O(n), где n - длина строки
func ReverseLetters(s string) string {
	var letters []rune // Создаем слайс рун для хранения только букв

	// Фильтруем строку, оставляя только буквы
	for _, char := range s {
		if unicode.IsLetter(char) {
			letters = append(letters, char)
		}
	}

	// Переворачиваем слайс рун
	left, right := 0, len(letters)-1
	for left < right {
		letters[left], letters[right] = letters[right], letters[left]
		left++
		right--
	}

	return string(letters) // Преобразуем слайс рун обратно в строку
}
