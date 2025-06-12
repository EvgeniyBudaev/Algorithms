package main

import (
	"fmt"
)

/* Convert boolean values to strings 'Yes' or 'No'.
https://www.codewars.com/kata/53369039d7ab3ac506000467/train/go

Завершите метод, который принимает логическое значение и возвращает строку "Yes" для значения «true» или строку "No"
для значения «false».
*/

func main() {
	fmt.Println(BoolToWord(true))  // "Yes"
	fmt.Println(BoolToWord(false)) // "No"
}

func BoolToWord(word bool) string {
	if word {
		return "Yes"
	}
	return "No"
}
