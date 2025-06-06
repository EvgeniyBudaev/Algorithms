package main

import (
	"fmt"
	"strings"
)

/* Duplicate Encoder

Цель этого упражнения — преобразовать строку в новую строку, где каждый символ в новой строке — это "(",
если этот символ встречается в исходной строке только один раз, или ")", если этот символ встречается в исходной строке
более одного раза. Игнорируйте регистр букв при определении того, является ли символ дубликатом.

Examples:
"din"      =>  "((("
"recede"   =>  "()()()"
"Success"  =>  ")())())"
"(( @"     =>  "))(("
*/

func main() {
	fmt.Println(DuplicateEncode("din"))     // "((("
	fmt.Println(DuplicateEncode("recede"))  // "()()()"
	fmt.Println(DuplicateEncode("Success")) // ")())())"
	fmt.Println(DuplicateEncode("(( @"))    // "))(("
}

// DuplicateEncode преобразует строку в новую строку с использованием символов "(" и ")".
// time: O(n), space: O(n)
func DuplicateEncode(word string) string {
	lowerWord := strings.ToLower(word) // Преобразование строки к нижнему регистру
	charCounter := make(map[rune]int)  // Используем карту для подсчета
	result := ""                       // Переменная для хранения

	for _, char := range lowerWord { // Подсчет количества вхождений каждого символа в строке
		charCounter[char]++
	}

	for _, char := range lowerWord { // Создание результата на основе количества вхождений каждого символа
		if charCounter[char] == 1 {
			result += "("
		} else {
			result += ")"
		}
	}

	return result // Возвращаем результат
}
