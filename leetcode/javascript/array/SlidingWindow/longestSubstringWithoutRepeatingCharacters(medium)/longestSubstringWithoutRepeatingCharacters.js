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

    let left = 0, right = 0, ans = 0;
    const set = new Set();

    while (right < s.length) {
        if (set.has(s[right])) {
            set.delete(s[left]);
            left++;
        } else {
            set.add(s[right]);
            ans = Math.max(ans, right - left + 1);
            right++;
        }
    }

    return ans;
};

var lengthOfLongestSubstringWithMap = function(s) {
    if (s.length <= 1) return s.length;

    let left = 0, max = 0;
    let map = new Map();

    for (let right = 0; right < s.length; right++) {
        let current = s[right];
        if (map.has(current)) {
            left = Math.max(map.get(current) + 1, left);
        }
        max = Math.max(max, right - left + 1);
        map.set(current, right);
    }

    return max;
};

console.log(lengthOfLongestSubstringWithMap("abcabcbb")); // 3