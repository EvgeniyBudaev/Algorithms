/* https://leetcode.com/problems/minimum-window-substring/description/
solution https://leetcode.com/problems/minimum-window-substring/solutions/4673781/easy-sliding-window-solution-explained/

Учитывая две строки s и t длиной m и n соответственно, верните минимальное окно подстрока из s так, что каждый символ из
t (включая дубликаты) включен в окно. Если такой подстроки нет, верните пустую строку «».
Тестовые примеры будут сгенерированы таким образом, чтобы ответ был уникальным.

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Объяснение: Подстрока минимального окна "BANC" включает "A", "B" и "C" из строки t.

Input: s = "a", t = "a"
Output: "a"
Объяснение: Вся строка s представляет собой минимальное окно.

Input: s = "a", t = "aa"
Output: ""
Объяснение: В окно должны быть включены обе буквы 'a' из t.
Поскольку самое большое окно s имеет только одну букву «а», верните пустую строку.
*/

/**
 * Time O(n), Space O(n)
 * @param {string} s
 * @param {string} t
 * @return {string}
 */
var minWindow = function(s, t) {
    let cnt = {}, len = t.length, res = [], min = Infinity;
    for (let char of t) {
        cnt[char] = (cnt[char] || 0) + 1;
    }

    for (let r = 0, l = 0; r < s.length; r++) {
        let current = s[r];
        if (cnt[current] > 0) len--;
        cnt[current]--;
        while (!len) {
            if (r - l < min) {
                min = r - l;
                res = [l, r];
            }
            if (cnt[s[l]] >= 0) len++;
            cnt[s[l]]++;
            l++;
        }
    }

    return s.slice(res[0], res[1] + 1);
};

console.log(minWindow("ADOBECODEBANC", "ABC")); // "BANC"