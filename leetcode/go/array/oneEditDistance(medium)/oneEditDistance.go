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

func isOneEditDistance(s string, t string) bool {
	m, n := len(s), len(t)
	if m < n {
		return isOneEditDistance(t, s)
	}
	if m-n > 1 {
		return false
	}

	for i := range t {
		if s[i] != t[i] {
			if m == n {
				return s[i+1:] == t[i+1:]
			}
			return s[i+1:] == t[i:]
		}
	}

	return m == n+1
}
