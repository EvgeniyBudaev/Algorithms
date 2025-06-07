package main

import (
	"fmt"
	"strings"
)

/* Convert string to camel case

Дополните метод/функцию так, чтобы он преобразовывал слова, разделенные тире/подчеркиванием, в camel case.
Первое слово в выходных данных должно быть написано заглавными буквами, только если исходное слово было написано
заглавными буквами (известно как Upper Camel Case, также часто называемый Pascal case). Следующие слова должны всегда
быть написаны заглавными буквами.

Examples
"the-stealth-warrior" gets converted to "theStealthWarrior"
"The_Stealth_Warrior" gets converted to "TheStealthWarrior"
"The_Stealth-Warrior" gets converted to "TheStealthWarrior"
*/

func main() {
	fmt.Println(ToCamelCase("the-stealth-warrior")) // "theStealthWarrior"
	fmt.Println(ToCamelCase("The_Stealth_Warrior")) // "theStealthWarrior"
	fmt.Println(ToCamelCase("The_Stealth-Warrior")) // "theStealthWarrior"
}

// ToCamelCase преобразует строку в camel case.
// time: O(n), space: O(n)
func ToCamelCase(s string) string {
	if s == "" {
		return ""
	}

	// Разделяем строку по разделителям '-' или '_'
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == '-' || r == '_'
	})

	// Если строка пустая после разделения (например, была "--"), возвращаем пустую строку
	if len(words) == 0 {
		return ""
	}

	for i := 1; i < len(words); i++ {
		words[i] = CapitalizeFirstLetter(words[i])
	}

	return strings.Join(words, "")
}

// CapitalizeFirstLetter возвращает строку с заглавной первой буквой.
// time: O(n), space: O(n)
func CapitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
