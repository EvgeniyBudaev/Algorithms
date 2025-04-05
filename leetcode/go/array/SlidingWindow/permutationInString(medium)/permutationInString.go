package main

import "fmt"

/* 567. Permutation in String
https://leetcode.com/problems/permutation-in-string/description/

Учитывая две строки s1 и s2, верните true, если s2 содержит перестановку s1, или false в противном случае.
Другими словами, верните true, если одна из перестановок s1 является подстрокой s2.

Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Пояснение: s2 содержит одну перестановку s1 («ba»).

Input: s1 = "ab", s2 = "eidboaoo"
Output: false
*/

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo")) // true
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	// Создаем карту для подсчета символов в s1
	neededChar := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		neededChar[s1[i]]++
	}

	left, right, requiredLength := 0, 0, len(s1)

	for right < len(s2) {
		currentChar := s2[right]
		if neededChar[currentChar] > 0 {
			requiredLength--
		}
		neededChar[currentChar]--
		right++

		// Если все символы s1 найдены в текущем окне
		if requiredLength == 0 {
			return true
		}

		// Если окно достигло длины s1, сдвигаем левую границу
		if right-left == len(s1) {
			if neededChar[s2[left]] >= 0 {
				requiredLength++
			}
			neededChar[s2[left]]++
			left++
		}
	}

	return false
}
