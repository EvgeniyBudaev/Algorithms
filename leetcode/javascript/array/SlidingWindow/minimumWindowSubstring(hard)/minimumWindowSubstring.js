/* https://leetcode.com/problems/minimum-window-substring/description/
solution https://leetcode.com/problems/minimum-window-substring/solutions/4673781/easy-sliding-window-solution-explained/

Учитывая две строки s и t длиной m и n соответственно, верните минимальную подстроку окна s, такую, что каждый символ
в t (включая дубликаты) включен в окно. Если такой подстроки нет, верните пустую строку «».

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
    if (s.length < t.length) return "";
    let map = {}, len = t.length, ans = [], min = Infinity;

    for (let char of t) {
        map[char] = (map[char] || 0) + 1;
    }

    for (let right = 0, left = 0; right < s.length; right++) {
        const current = s[right];
        if (map[current] > 0) len--;
        map[current]--;
        while (!len) {
            if (right - left < min) {
                min = right - left;
                ans = [left, right];
            }
            if (map[s[left]] >= 0) len++;
            map[s[left]]++;
            left++;
        }
    }

    return s.slice(ans[0], ans[1] + 1);
};

console.log(minWindow("ADOBECODEBANC", "ABC")); // "BANC"