package main

import "fmt"

/* 763. Partition Labels

Вам дана строка s. Мы хотим разделить строку на как можно большее количество частей, чтобы каждая буква появлялась не
более чем в одной части. Обратите внимание, что разбиение выполнено так, что после объединения всех частей по порядку
результирующая строка должна быть s. Верните список целых чисел, представляющих размер этих частей.

Input: s = "ababcbacadefegdehijhklij"
Output: [9,7,8]
Объяснение:
Раздел — «ababcbaca», «defegde», «hijhklij».
Это раздел, в котором каждая буква встречается не более чем в одной части.
Раздел типа «ababcbacadefegde», «hijhklij» неверен, поскольку он разбивает s на меньшее количество частей.

Input: s = "eccbbbbdec"
Output: [10]
*/

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij")) // [9,7,8]
}

func partitionLabels(s string) []int {
	lastIdx := make(map[rune]int)
	for i, char := range s {
		lastIdx[char] = i
	}

	curLastIdx := 0
	res := []int{}
	acc := 0

	for i, char := range s {
		if lastIdx[char] > curLastIdx {
			curLastIdx = lastIdx[char]
		}
		if i == curLastIdx {
			res = append(res, i+1-acc)
			acc = i + 1
		}
	}

	return res
}
