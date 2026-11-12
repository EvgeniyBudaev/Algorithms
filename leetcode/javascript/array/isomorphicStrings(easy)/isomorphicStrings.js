/* 205. Isomorphic Strings
https://leetcode.com/problems/isomorphic-strings/description/

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
 * @time O(n)
 * @space O(1)
 */
var isIsomorphic = function(s, t) {
    if (s.length !== t.length) return false;
    
    const mapS = new Map(); // Хранит соответствие s[i] -> t[i]
    const mapT = new Map(); // Хранит соответствие t[i] -> s[i]
    
    for (let i = 0; i < s.length; i++) {
        const charS = s[i];
        const charT = t[i];
        
        // Если для charS уже есть сопоставление, но оно не равно charT
        if (mapS.has(charS) && mapS.get(charS) !== charT) return false;
        
        // Если для charT уже есть сопоставление, но оно не равно charS
        // (это предотвращает ситуацию, когда два разных символа s маппятся на один символ t)
        if (mapT.has(charT) && mapT.get(charT) !== charS) return false;
        
        // Записываем новые сопоставления
        mapS.set(charS, charT);
        mapT.set(charT, charS);
    }
    
    return true;
};

console.log(isIsomorphic("egg", "add")); // true
console.log(isIsomorphic("foo", "bar")); // false