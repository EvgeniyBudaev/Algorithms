package main

import (
	"fmt"
)

/* Longest vowel chain
https://www.codewars.com/kata/59c5f4e9d751df43cf000035/train/go

Подстроки гласных в слове codewarriors: o, e, a, io. Самая длинная из них имеет длину 2. Если задана строка строчных
букв, состоящая только из букв (гласных и согласных) и не содержащая пробелов, верните длину самой длинной подстроки
гласных. Гласные — это любые буквы из списка aeiou.
*/

func main() {
	fmt.Println(Solve("codewarriors"))          // 2
	fmt.Println(Solve("suoidea"))               // 3
	fmt.Println(Solve("ultrarevolutionariees")) // 3
	fmt.Println(Solve("strengthlessnesses"))    // 1
	fmt.Println(Solve("cuboideonavicuare"))     // 2
	fmt.Println(Solve("chrononhotonthuooaos"))  // 5
	fmt.Println(Solve("iiihoovaeaaaoougjyaw"))  // 8
}

// Solve возвращает длину самой длинной подстроки гласных в строке.
// time: O(n), space: O(1), где n - длина строки
func Solve(s string) int {
	vowels := map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}

	maxLength := 0
	currentLength := 0

	for _, char := range s {
		if _, ok := vowels[char]; ok {
			currentLength++
		} else {
			if currentLength > maxLength {
				maxLength = currentLength
			}
			currentLength = 0
		}
	}

	return maxLength
}
