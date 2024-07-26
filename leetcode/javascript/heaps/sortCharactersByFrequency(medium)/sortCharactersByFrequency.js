/* https://leetcode.com/problems/sort-characters-by-frequency/description/

Дана строка s, отсортируйте ее в порядке убывания частоты встречаемости символов. Частота символа — это количество раз,
которое он появляется в строке.
Верните отсортированную строку. Если ответов несколько, верните любой из них.

Input: s = "tree"
Output: "eert"
Пояснение: «e» появляется дважды, а «r» и «t» появляются по одному разу.
Таким образом, «e» должно стоять перед «r» и «t». Поэтому «eetr» также является допустимым ответом.

Input: s = "cccaaa"
Output: "aaaccc"
Пояснение: И «c», и «a» встречаются три раза, поэтому и «cccaaa», и «aaaccc» являются правильными ответами.
Обратите внимание, что слово «какаса» неверно, поскольку одни и те же символы должны быть вместе.

Input: s = "Aabb"
Output: "bbAa"
Пояснение: «bbaA» также является правильным ответом, но «Aabb» неверен.
Обратите внимание, что «A» и «a» рассматриваются как два разных символа.
*/

/**
 * @param {string} s
 * @return {string}
 */
var frequencySort = function(s) {
    const counter = new Map();

    for (const char of s) {
        counter.set(char, (counter.get(char) || 0) + 1);
    }

    const pq = Array.from(counter.entries());
    pq.sort((a, b) => b[1] - a[1]);
    let result = '';

    for (const [char, freq] of pq) {
        result += char.repeat(freq);
    }

    return result;
};

console.log(frequencySort("tree")); // "eert"