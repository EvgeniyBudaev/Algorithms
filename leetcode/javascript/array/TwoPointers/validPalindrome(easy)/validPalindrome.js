/* https://leetcode.com/problems/valid-palindrome/description/

Фраза является палиндромом, если после преобразования всех прописных букв в строчные и удаления всех небуквенно-цифровых
символов она читается одинаково и вперед, и назад. Буквенно-цифровые символы включают буквы и цифры.
Учитывая строку s, верните true, если это палиндром, или false в противном случае.

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
 */

/**
 * Time O(N) | Space O(1)
 * @param {string} s
 * @return {boolean}
 */
var isPalindrome = function(s) {
    const isAlphaNumeric = (c) => (c.toLowerCase() >= 'a' && c.toLowerCase() <= 'z') || c >= '0' && c <= '9';

    let left = 0;
    let right = s.length - 1;
    let skipLeft, skipRight, endsEqual = false;

    while (left < right) {
        skipLeft = !isAlphaNumeric(s[left])
        if (skipLeft) { left++; continue; }

        skipRight = !isAlphaNumeric(s[right])
        if (skipRight) { right--; continue; }

        endsEqual = s[left].toLowerCase() === s[right].toLowerCase();
        if (!endsEqual) return false;

        left++;
        right--;
    }

    return true;
};

console.log(isPalindrome("A man, a plan, a canal: Panama")); // true