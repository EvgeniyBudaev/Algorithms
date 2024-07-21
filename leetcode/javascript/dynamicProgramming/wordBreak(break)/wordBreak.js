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

/**
 * @param {string} s
 * @param {string[]} wordDict
 * @return {boolean}
 */
var wordBreak = function(s, wordDict) {
    const n = s.length;

    // Создайте набор слов из массива, чтобы стоимость поиска в наборе стала O (1).
    const wordSet = new Set(wordDict);

    // Инициализировать таблицу поиска
    const dp = new Array(n + 1).fill(false);

    // Установите для первого элемента значение true, поскольку он представляет пустую строку.
    dp[0] = true;

    for (let i = 1; i <= n; i++) {
        for (let j = 0; j < i; j++) {
            // Проверка наличия подстроки от j до i в wordSet и истинности dp[j]
            if (dp[j] && wordSet.has(s.substring(j, i))) {
                dp[i] = true;
                // Если подстрока найдена, нет необходимости проверять дальнейшие меньшие подстроки.
                break;
            }
        }
    }

    // Вернуть последний элемент массива dp
    return dp[n];
};

const s = "leetcode";
const wordDict = ["leet","code"];
console.log(wordBreak(s, wordDict));