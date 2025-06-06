package main

import (
	"fmt"
	"sort"
	"strings"
)

/* Two to One

Возьмите 2 строки s1 и s2, содержащие только буквы от a до z. Верните новую отсортированную строку (в алфавитном порядке
возрастания), максимально длинную, содержащую различные буквы — каждая взята только один раз — из s1 или s2.

Examples:
a = "xyaabbbccccdefww"
b = "xxxxyyyyabklmopq"
longest(a, b) -> "abcdefklmopqwxy"
*/

func main() {
	fmt.Println(TwoToOne("xyaabbbccccdefww", "xxxxyyyyabklmopq"))                     // "abcdefklmopqwxy"
	fmt.Println(TwoToOne("abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyz")) // "abcdefghijklmnopqrstuvwxyz"
}

// TwoToOne возвращает новую отсортированную строку (в алфавитном порядке) из s1 или s2.
// time: (n), space: O(n)
func TwoToOne(s1 string, s2 string) string {
	chars := strings.Split(s1+s2, "")
	sort.Strings(chars)
	result := ""

	for _, s := range chars {
		if !strings.Contains(result, s) {
			result += s
		}
	}

	return result
}

// TwoToOne возвращает новую отсортированную строку (в алфавитном порядке) из s1 или s2.
// time: (n), space: O(n)
//func TwoToOne(s1 string, s2 string) string {
//	charCounter := make(map[rune]bool)
//
//	for _, char := range s1 {
//		charCounter[char] = true
//	}
//
//	for _, char := range s2 {
//		charCounter[char] = true
//	}
//
//	var result []rune
//	for char := range charCounter {
//		result = append(result, char)
//	}
//
//	sort.Slice(result, func(i, j int) bool {
//		return result[i] < result[j]
//	})
//
//	return string(result)
//}
