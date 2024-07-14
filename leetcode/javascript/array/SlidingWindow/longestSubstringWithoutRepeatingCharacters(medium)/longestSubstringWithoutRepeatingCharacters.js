/* https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

Дана строка s. Найдите длину самой длинной подстрока без повторения символов.

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

/**
 * Time O(N) | Space O(N)
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function(s) {
    if (s.length <= 1) return s.length;

    let left = 0, right = 0, maxLength = 0;
    const charSet = new Set();

    while (right < s.length) {
        if (charSet.has(s[right])) {
            charSet.delete(s[left]);
            left++;
        } else {
            charSet.add(s[right]);
            right++;
            maxLength = Math.max(maxLength, right - left);
        }
    }

    return maxLength;
};

var lengthOfLongestSubstringWithMap = function(s) {
    if (s.length <= 1) return s.length;

    let left = 0, max = 0;
    let charMap = new Map();

    for (let i = 0; i < s.length; i++) {
        let current = s[i];
        if (charMap.has(current)) {
            left = Math.max(charMap.get(current) + 1, left);
        }
        max = Math.max(max, (i - left) + 1);
        charMap.set(current, i);
    }

    return max;
};

console.log(lengthOfLongestSubstringWithMap("abcabcbb")); // 3