package main

import (
	"fmt"
	"strings"
	"unicode"
)

/* Mumbling

Examples:
accum("abcd") -> "A-Bb-Ccc-Dddd"
accum("RqaEzty") -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
accum("cwAt") -> "C-Ww-Aaa-Tttt"
*/

func main() {
	fmt.Println(Accum("abcd")) // "A-Bb-Ccc-Dddd"
}

// Accum принимает строку и возвращает преобразованную строку в соответствии с заданными правилами.
// time: O(n), space: O(n)
func Accum(s string) string {
	var result []string // Используем слайс для хранения преобразованных подстрок

	for i, char := range s {
		firstChar := string(unicode.ToUpper(char))               // Преобразование первого символа в верхний регистр
		rest := strings.Repeat(string(unicode.ToLower(char)), i) // Преобразование оставшихся символов в нижний регистр и повторение их i раз
		result = append(result, firstChar+rest)                  // Добавление преобразованной подстроки к результату
	}

	return strings.Join(result, "-") // Объединение преобразованных подстрок с помощью дефисов
}
