package main

import "fmt"

/* https://leetcode.com/problems/word-break/description/

Учитывая строку s и словарь строк wordDict, верните true, если s можно сегментировать в разделенную пробелами
последовательность из одного или нескольких словарных слов.
Обратите внимание, что одно и то же слово в словаре может использоваться при сегментации несколько раз.

Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
Объяснение: Возвращайте true, потому что "leetcode" можно сегментировать как "leet code".

Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Объяснение: Возвращайте true, поскольку "applepenapple" можно сегментировать как "apple pen apple".
Обратите внимание, что вам разрешено повторно использовать словарное слово.

Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: false
*/

func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak(s, wordDict)) // true ("leet" + "code")
}

// wordBreak проверяет, можно ли разбить строку на слова из словаря
func wordBreak(s string, wordDict []string) bool {
	n := len(s)

	// Создаем map для быстрого поиска слов
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	// Инициализируем массив для динамического программирования
	dp := make([]bool, n+1)
	dp[0] = true // Пустая строка считается разбиваемой

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			// Проверяем, можно ли разбить подстроку s[0:j] и содержится ли s[j:i] в словаре
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break // Нашли разбиение, переходим к следующему i
			}
		}
	}

	return dp[n]
}
