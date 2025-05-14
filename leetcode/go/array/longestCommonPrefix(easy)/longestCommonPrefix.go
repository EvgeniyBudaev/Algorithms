package main

import (
	"fmt"
	"strings"
)

/* 14. Longest Common Prefix
https://leetcode.com/problems/longest-common-prefix/description/

Напишите функцию для поиска самой длинной общей строки префикса среди массива строк.
Если общего префикса нет, верните пустую строку "".

Input: strs = ["flower","flow","flight"]
Output: "fl"
*/

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs)) // "fl"
}

// longestCommonPrefix - возвращает самый длинный общий префикс среди массива строк.
// time: O(S), где S - суммарная длина всех строк в массиве strs.
// space: O(1)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 { // Если массив пуст, возвращаем пустую строку.
		return ""
	}

	prefix := strs[0]                // Задаем префикс как первую строку. "flower"
	for i := 1; i < len(strs); i++ { // "flow", "flight"
		// Проверяем, что длина префикса не превышает длину строки.
		for strings.Index(strs[i], prefix) != 0 { // strings.Index - возвращает индекс первого вхождения подстроки в строке.
			// Удаляем последний символ из prefix, пока не найдем общий префикс.
			prefix = prefix[:len(prefix)-1] // Здесь используется срез для получения подстроки. "flowe", "flow", "flo", "fl"
			if prefix == "" {               // Если префикс стал пустым, возвращаем его.
				return ""
			}
		}
	}

	return prefix // Возвращаем найденный префикс.
}
