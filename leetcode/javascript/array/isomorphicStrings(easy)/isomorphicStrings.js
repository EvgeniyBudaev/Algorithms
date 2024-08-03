/* https://leetcode.com/problems/isomorphic-strings/description/

Даны две строки s и t. Определите, изоморфны ли они.
Две строки s и t изоморфны, если символы в s можно заменить, чтобы получить t.
Все вхождения символа должны быть заменены другим символом с сохранением порядка символов. Никакие два символа не могут
сопоставляться одному и тому же символу, но символ может сопоставляться сам с собой.
Например, рассмотрим строки ACAB и XCXY.
Они изоморфны, поскольку мы можем отобразить 'A' —> 'X', 'B' —> 'Y'и 'C' —> 'C'.

Input: s = "egg", t = "add"
Output: true

Input: s = "foo", t = "bar"
Output: false

Input: s = "paper", t = "title"
Output: true
*/

/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isIsomorphic = function(s, t) {
    if (s.length !== t.length) return false;
    for (let i = 0; i < s.length; i++) {
        if (s.indexOf(s[i], i + 1) !== t.indexOf(t[i], i + 1)) return false;
    }
    return true;
};

console.log(isIsomorphic("egg", "add")); // true
console.log(isIsomorphic("foo", "bar")); // false