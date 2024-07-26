/* https://leetcode.com/problems/reorganize-string/description/
solution https://leetcode.com/problems/reorganize-string/submissions/1333724578/

Дана строка s, переставьте символы s так, чтобы любые два соседних символа не были одинаковыми.
Верните любую возможную перестановку s или верните "", если это невозможно.

Input: s = "aab"
Output: "aba"

Input: s = "aaab"
Output: ""
*/

/**
 * @param {string} s
 * @return {string}
 */
var reorganizeString = function(s) {
    const charMap = {};
    const res = [];

    for (let char of s) {
        charMap[char] = (charMap[char] || 0) + 1;
    }

    const sortedMap = Object.entries(charMap).sort((a, b) => b[1] - a[1]);
    let position = 0;

    if (sortedMap[0][1] > Math.floor((s.length + 1) / 2)) return '';

    for (let entry of sortedMap) {
        const char = entry[0]
        const count = entry[1];
        for (let j = 0; j < count; j++) {
            res[position] = char;
            position +=2;
            if (position >= s.length) position = 1;
        }
    }

    return res.join('');
};

console.log(reorganizeString("aab")); // "aba"