package leetcode.java.dynamicProgramming.WordBreak;

import java.util.Arrays;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

/* 139. Word Break
https://leetcode.com/problems/word-break/description/

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

public class WordBreak {
    public static void main(String[] args) {
        String s = "leetcode";
        List<String> wordDict = Arrays.asList("leet", "code");
        System.out.println(wordBreak(s, wordDict)); // true ("leet" + "code")
    }

    // wordBreak проверяет, можно ли разбить строку на слова из словаря
    // time: O(n² × k), space: O(n + w)
    private static boolean wordBreak(String s, List<String> wordDict) {
        int n = s.length(); // Длина входной строки

        // Создаём HashSet для быстрого поиска слов — O(w) по памяти
        Set<String> wordSet = new HashSet<>(wordDict);

        // Инициализируем массив для динамического программирования
        boolean[] dp = new boolean[n + 1];
        dp[0] = true; // Пустая строка считается разбиваемой

        // Проходим по всем возможным длинам подстроки — O(n)
        for (int i = 1; i <= n; i++) {
            // Проходим по всем возможным длинам предыдущей подстроки — O(n)
            for (int j = 0; j < i; j++) {
                // Проверяем: можно ли разбить s[0:j] и есть ли s[j:i] в словаре
                if (dp[j] && wordSet.contains(s.substring(j, i))) {
                    dp[i] = true; // Подстрока разбиваема
                    break;        // Нашли разбиение, переходим к следующему i
                }
            }
        }

        return dp[n];
    }
}
