package main

import "fmt"

/* 161. One Edit Distance
https://github.com/EvgeniyBudaev/doocs_leetcode/tree/main/solution/0100-0199/0161.One%20Edit%20Distance

Учитывая две строки s и t, определите, находятся ли они на расстоянии одного редактирования друг от друга.

s Вставьте в него ровно один символ , чтобы получить t
s Удалите из него ровно один символ, чтобы получить t
Замена ровно одного символа другим символом в s дает t

Input: s = "ab", t = "acb"
Output: true
Объяснение: в строку s можно вставить 'c',  чтобы получить t.

Input: s = "cab", t = "ad"
Output: false
*/

func main() {
	fmt.Println(isOneEditDistance("ab", "acb")) // true
}

// isOneEditDistance определяет, находятся ли строки на расстоянии одного редактирования друг от друга.
// time: O(n), space: O(1)
func isOneEditDistance(s string, t string) bool {
	m, n := len(s), len(t)
	if m < n {
		// Меняем местами строки, чтобы s всегда была длиннее строки t.
		return isOneEditDistance(t, s)
	}
	// Если разница в длинах больше 1, то строки не могут быть на расстоянии одного редактирования друг от друга.
	if m-n > 1 {
		return false
	}

	// s: "acb", t: "ab"
	for i := range t {
		// Если символы строки s и t не совпадают, то проверяем, что строка s имеет длину n+1.
		if s[i] != t[i] {
			// Если строки имеют одинаковую длину, то проверяем, что строки равны, начиная с i+1 символа.
			if m == n {
				return s[i+1:] == t[i+1:]
			}
			// Если строки имеют разную длину, то проверяем, что строка s имеет длину n+1.
			// s[i+1:]: "b" t[i:]: "b
			return s[i+1:] == t[i:]
		}
	}

	// Если все символы строки s и t совпадают, то проверяем, что строка s имеет длину n+1.
	return m == n+1
}
