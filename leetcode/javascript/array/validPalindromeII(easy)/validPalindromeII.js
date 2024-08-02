/* https://leetcode.com/problems/valid-palindrome-ii/description/

Учитывая строку s, верните true, если s может быть палиндромом после удаления из нее не более одного символа.

Input: s = "aba"
Output: true

Input: s = "abca"
Output: true
Объяснение: Вы можете удалить символ 'c'.

Input: s = "abc"
Output: false
*/

/**
 * @param {string} s
 * @return {boolean}
 */
var validPalindrome = function(s) {
    let low = 0, high = s.length - 1;
    while (low < high) {
        if (s[low] !== s[high]) {
            return isPalindrome(s, low + 1, high) || isPalindrome(s, low, high - 1);
        }
        low++;
        high--;
    }
    return true;
};

function isPalindrome(str, low, high) {
    while (low < high) {
        if (str[low] !== str[high]) return false;
        low++;
        high--;
    }
    return true;
}

console.log(validPalindrome("aba")); // true