/* https://leetcode.com/problems/find-all-anagrams-in-a-string/description/

Учитывая две строки s и p, верните массив всех начальных индексов анаграмм p в s. Вы можете вернуть ответ в любом
порядке.
Анаграмма — это слово или фраза, образованная путем перестановки букв другого слова или фразы, обычно с использованием
всех исходных букв ровно один раз.

Input: s = "cbaebabacd", p = "abc"
Output: [0,6]
Объяснение:
Подстрока с начальным индексом = 0 — это «cba», которая является анаграммой слова «abc».
Подстрока с начальным индексом = 6 — это «bac», которая является анаграммой слова «abc».
*/

/**
 * @param {string} s
 * @param {string} p
 * @return {number[]}
 */
var findAnagrams = function(s, p) {
    const arr = [];
    const obj = {};

    for (let i of p) {
        obj[i] ? obj[i]++ : obj[i] = 1;
    }

    let left = 0, right = 0, count = p.length; // count: 3

    while (right < s.length) {
        if (obj[s[right]] > 0) count--;
        obj[s[right]]--;
        right++;
        if (count === 0) arr.push(left);
        if (right - left === p.length) {
            if (obj[s[left]] >= 0) count++;
            obj[s[left]]++;
            left++;
        }
    }

    return arr;
};

console.log(findAnagrams("cbaebabacd", "abc")); // [0,6]