/* https://leetcode.com/problems/ransom-note/description/

Учитывая две строки ransomNote и magazine, верните true, если ransomNote можно создать с использованием букв из magazine,
и false в противном случае.
Каждое письмо в magazine можно использовать в ransomNote только один раз.

Input: ransomNote = "a", magazine = "b"
Output: false

Input: ransomNote = "aa", magazine = "ab"
Output: false

Input: ransomNote = "aa", magazine = "aab"
Output: true
*/

/**
 * @param {string} ransomNote
 * @param {string} magazine
 * @return {boolean}
 */
var canConstruct = function(ransomNote, magazine) {
     const map = {};

     for (const letter of magazine) {
         map[letter] = map[letter] + 1 || 1;
     }

     for (let letter of ransomNote) {
         if (!map[letter]) return false;
         map[letter]--;
     }

     return true;
};

console.log(canConstruct("aa", "ab")); // false