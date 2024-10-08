/* https://leetcode.com/problems/top-k-frequent-words/description/

Учитывая массив строковых слов и целое число k, верните k наиболее часто встречающихся строк.
Возвращает ответ, отсортированный по частоте от самого высокого до самого низкого. Отсортируйте слова с одинаковой
частотой по их лексикографическому порядку.

Input: words = ["i","love","leetcode","i","love","coding"], k = 2
Output: ["i","love"]
Пояснение: «я» и «люблю» — два наиболее часто встречающихся слова.
Обратите внимание, что «i» стоит перед словом «love» из-за более низкого алфавитного порядка.
*/

/**
 * @param {string[]} words
 * @param {number} k
 * @return {string[]}
 */
var topKFrequent = function(words, k) {
    const map = new Map();

    for (let i = 0; i < words.length; i++) {
        const word = words[i];
        map.set(word, (map.get(word) ?? 0) + 1);
    }
    // [...map.entries()]: [['i', 2], ['love', 2], ['leetcode', 1], ['coding', 1]]
    const arr = [...map.entries()] // arr: [['i', 2], ['love', 2], ['coding', 1], ['leetcode', 1]]
        .sort((a, b) => b[1] !== a[1] ? b[1] - a[1] : a[0]
            .localeCompare(b[0]));

    return arr.slice(0, k).map(item => item[0]);
};

const words = ["i","love","leetcode","i","love","coding"];
console.log(topKFrequent(words, 2)); // ["i","love"]