package main

import (
	"fmt"
	"strings"
	"unicode"
)

/* V A P O R C O D E
https://www.codewars.com/kata/5966eeb31b229e44eb00007a/train/go

Напишите функцию, которая преобразует любое предложение в предложение с V A P O R W A V E.
Предложение с V A P O R W A V E преобразует все буквы в заглавные и добавляет два пробела между каждой буквой
(или специальным символом), создавая эффект V A P O R W A V E.

Обратите внимание, что в этом случае пробелы следует игнорировать.

Examples:
"Lets go to the movies"       -->  "L  E  T  S  G  O  T  O  T  H  E  M  O  V  I  E  S"
*/

func main() {
	fmt.Println(Vaporcode("Lets go to the movies")) // "L  E  T  S  G  O  T  O  T  H  E  M  O  V  I  E  S"
}

// Vaporcode возвращает строку, где каждая буква преобразована в заглавную букву и добавлена два пробела между собой.
// time: O(n), space: O(n), n - длина строки
func Vaporcode(s string) string {
	var letters []string
	for _, v := range s {
		if v != ' ' {
			letters = append(letters, string(unicode.ToUpper(v)))
		}
	}
	return strings.Join(letters, "  ")
}
