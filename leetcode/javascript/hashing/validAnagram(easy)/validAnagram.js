/* https://leetcode.com/problems/valid-anagram/description/

Учитывая две строки s и t, верните true, если t является анаграммой s, и false в противном случае.
Анаграмма — это слово или фраза, образованная путем перестановки букв другого слова или фразы, обычно с использованием
всех исходных букв ровно один раз.

Input: s = "anagram", t = "nagaram"
Output: true

Input: s = "rat", t = "car"
Output: false
*/

/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isAnagram = function(s, t) {
    if (s.length !== t.length) return false;
    let map = new Map();

    for (let char of s) {
        if (map.has(char)) map.set(char, map.get(char) + 1);
        else map.set(char, 1);
    }

    for (let char of t) {
        if (!map.has(char) || map.get(char) <= 0) return false;
        map.set(char, map.get(char) - 1);
    }

    return true;
};

console.log(isAnagram("anagram", "nagaram")); // true