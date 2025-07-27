package main

import (
	"fmt"
)

/* Vowel Count
https://www.codewars.com/kata/54ff3102c1bad923760001f3/train/go

Возвращает количество гласных в заданной строке.
Мы будем считать a, e, i, o, u гласными для этой ката (но не y).
Входная строка будет состоять только из строчных букв и/или пробелов.
*/

func main() {
	fmt.Println(GetCount("abracadabra")) // 5
}

// GetCount возвращает количество гласных в строке.
// time: O(n), space: O(1), где n - длина строки
func GetCount(str string) (count int) {
	for _, char := range str {
		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}
