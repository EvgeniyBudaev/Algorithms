package main

import (
	"fmt"
)

/* Ordered Count of Characters
https://www.codewars.com/kata/57a6633153ba33189e000074/train/go

Подсчитайте количество вхождений каждого символа и верните результат в виде списка кортежей в порядке их появления.
Для пустого вывода верните пустой список.

Examples:
OrderedCount("abracadabra") == []Tuple{Tuple{'a', 5}, Tuple{'b', 2}, Tuple{'r', 2}, Tuple{'c', 1}, Tuple{'d', 1}}

type Tuple struct {
    Char  rune
    Count int
}
*/

func main() {
	fmt.Println(OrderedCount("abracadabra")) // []Tuple{Tuple{'a', 5}, Tuple{'b', 2}, Tuple{'r', 2}, Tuple{'c', 1}, Tuple{'d', 1}}
	fmt.Println(OrderedCount("Code Wars"))   // []Tuple{Tuple{'C', 1}, Tuple{'o', 1}, Tuple{'d', 1}, Tuple{'e', 1}, Tuple{' ', 1}, Tuple{'W', 1}, Tuple{'a', 1}, Tuple{'r', 1}, Tuple{'s', 1}}
	fmt.Println(OrderedCount(""))            // []Tuple{}
}

type Tuple struct {
	Char  rune
	Count int
}

// OrderedCount возвращает количество вхождений каждого символа в виде списка кортежей в порядке их появления.
// time: O(n), space: O(n), n - длина входной строки
func OrderedCount(text string) []Tuple {
	if text == "" {
		return []Tuple{}
	}

	charToIndex := make(map[rune]int)
	var result []Tuple

	for _, char := range text {
		if index, exists := charToIndex[char]; exists {
			result[index].Count++
		} else {
			charToIndex[char] = len(result)
			result = append(result, Tuple{char, 1})
		}
	}

	return result
}
