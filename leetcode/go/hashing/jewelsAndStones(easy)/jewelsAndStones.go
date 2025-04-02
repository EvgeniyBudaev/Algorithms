package main

import "fmt"

/* https://leetcode.com/problems/jewels-and-stones/description/

Вам даны строки jewels, представляющие типы камней, которые являются jewels, и stones,
представляющие те камни, которые у вас есть. Каждый символ в stones — это тип камня, который у вас есть.
Вы хотите знать, сколько stones у вас тоже являются jewels.
Буквы чувствительны к регистру, поэтому «а» считается другим типом камня, чем «А».

Input: jewels = "aA", stones = "aAAbbbb"
Output: 3

Input: jewels = "z", stones = "ZZ"
Output: 0
*/

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb")) // 3
}

func numJewelsInStones(jewels string, stones string) int {
	count := 0
	jewelSet := make(map[rune]bool)

	// Создаем набор камней для поиска O(1)
	for _, j := range jewels {
		jewelSet[j] = true
	}

	// Считаем сколько камней в наборе
	for _, s := range stones {
		if jewelSet[s] {
			count++
		}
	}

	return count
}
