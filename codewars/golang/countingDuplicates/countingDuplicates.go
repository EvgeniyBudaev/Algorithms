package main

import (
	"fmt"
	"strings"
)

/* Counting Duplicates
https://www.codewars.com/kata/54bf1c2cd5b56cc47f0007a1/train/go

Подсчитайте количество дубликатов.
Напишите функцию, которая будет возвращать количество различных нечувствительных к регистру буквенных символов и
цифровых цифр, которые встречаются более одного раза во входной строке. Можно предположить, что входная строка содержит
только буквы алфавита (как заглавные, так и строчные) и цифровые цифры.

Example
"abcde" -> 0 # no characters repeats more than once
"aabbcde" -> 2 # 'a' and 'b'
"aabBcde" -> 2 # 'a' occurs twice and 'b' twice (`b` and `B`)
"indivisibility" -> 1 # 'i' occurs six times
"Indivisibilities" -> 2 # 'i' occurs seven times and 's' occurs twice
"aA11" -> 2 # 'a' and '1'
"ABBA" -> 2 # 'A' and 'B' each occur twice
*/

func main() {
	fmt.Println(duplicate_count("abcde"))          // 0 # no characters repeats more than once
	fmt.Println(duplicate_count("aabbcde"))        // 2 # 'a' and 'b'
	fmt.Println(duplicate_count("aabBcde"))        // 2 # 'a' occurs twice and 'b' twice (`b` and `B`)
	fmt.Println(duplicate_count("indivisibility")) // 1 # 'i' occurs six times
}

// duplicate_count возвращает количество дубликатов.
// time: O(n), space: O(n)
func duplicate_count(s1 string) int {
	counts := make(map[rune]int)
	for _, char := range strings.ToLower(s1) {
		counts[char]++
	}

	duplicates := 0
	for _, v := range counts {
		if v > 1 {
			duplicates++
		}
	}

	return duplicates
}
