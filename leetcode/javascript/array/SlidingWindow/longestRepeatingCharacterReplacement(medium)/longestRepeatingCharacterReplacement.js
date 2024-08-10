/* https://leetcode.com/problems/longest-repeating-character-replacement/description/

Вам дана строка s и целое число k. Вы можете выбрать любой символ строки и заменить его на любой другой символ
английского верхнего регистра. Эту операцию можно выполнить не более k раз.

Возвращает длину самой длинной подстроки, содержащей ту же букву, которую вы можете получить после выполнения
вышеуказанных операций.

Input: s = "ABAB", k = 2
Output: 4
Пояснение: Замените две буквы «А» двумя буквами «B» или наоборот.

Input: s = "AABABBA", k = 1
Output: 4
Пояснение: Замените букву «A» посередине на «B» и получите «AABBBBA».
Подстрока «BBBB» имеет самую длинную повторяющуюся букву — 4.
Могут существовать и другие способы получить этот ответ.
*/

/**
 * @param {string} s
 * @param {number} k
 * @return {number}
 */
var characterReplacement = function(s, k) {
    const map = {};
    let ans = 0, maxFreq = 0;
    let left = 0, right = 0;

    while (right < s.length) {
        let current = s[right];
        map[current] = (map[current] || 0) + 1;
        maxFreq = Math.max(maxFreq, map[current]);
        if (right - left + 1 > maxFreq + k) {
            map[s[left]]--;
            left++;
        }
        ans = Math.max(ans, right - left + 1);
        right++;
    }

    return ans;
};

console.log(characterReplacement("ABAB", 2)); // 4