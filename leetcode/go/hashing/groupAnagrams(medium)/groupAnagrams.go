package main

import (
	"fmt"
	"sort"
)

/* 49. Group Anagrams
https://leetcode.com/problems/group-anagrams/description/

Учитывая массив строк strs, сгруппируйте анаграммы вместе. Вы можете вернуть ответ в любом порядке.
Анаграмма — это слово или фраза, образованная путем перестановки букв другого слова или фразы, обычно с использованием
всех исходных букв ровно один раз.

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Input: strs = [""]
Output: [[""]]

Input: strs = ["a"]
Output: [["a"]]
*/

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs)) // [["bat"],["nat","tan"],["ate","eat","tea"]]
}

// groupAnagrams группирует анаграммы в исходном срезе строк
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	anagramMap := make(map[string][]string)

	for _, str := range strs {
		// Преобразовываем строку в фрагмент рун для правильной сортировки
		s := []rune(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s) // aet, aet, ant, aet, ant, abt

		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}
	// anagramMap: [[eat tea ate] [tan nat] [bat]]

	// Преобразовываем значения карты в срезы срезов
	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		// [eat tea ate]
		// [tan nat]
		// [bat]
		result = append(result, group)
	}

	return result
}
