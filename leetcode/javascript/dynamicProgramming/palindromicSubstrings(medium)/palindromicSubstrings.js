/* https://leetcode.com/problems/palindromic-substrings/description/

Дана строка s, верните количество палиндромных подстрок в ней.
Строка является палиндромом, если она читается как в прямом, так и в обратном направлении.
Подстрока — это непрерывная последовательность символов внутри строки.

Input: s = "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".

Input: s = "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
*/

/**
 * @param {string} s
 * @return {number}
 */
var countSubstrings = function(s) {
    let totalCount = 0;
    for (let i = 0; i < s.length; i++) {
        // Count palindromes with odd length
        totalCount += extendPalindrome(s, i, i);
        // Count palindromes with even length
        totalCount += extendPalindrome(s, i, i + 1);
    }
    return totalCount;
};

const extendPalindrome = (s, left, right) => {
    let count = 0;
    while (left >= 0 && right < s.length && s[left] === s[right]) {
        count++;
        left--;
        right++;
    }
    return count;
};
console.log(countSubstrings("abc")); // 3