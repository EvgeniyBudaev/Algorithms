/* https://leetcode.com/problems/permutation-in-string/description/

Учитывая две строки s1 и s2, верните true, если s2 содержит перестановку s1, или false в противном случае.
Другими словами, верните true, если одна из перестановок s1 является подстрокой s2.

Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Пояснение: s2 содержит одну перестановку s1 («ba»).

Input: s1 = "ab", s2 = "eidboaoo"
Output: false
*/

/**
 * @param {string} s1
 * @param {string} s2
 * @return {boolean}
 */
var checkInclusion = function(s1, s2) {
    if (s1.length > s2.length) return false;
    let neededChar = {};
    for (let i = 0; i < s1.length; i++) {
        neededChar[s1[i]] = (neededChar[s1[i]] || 0) + 1;
    }
    let left = 0, right = 0, requiredLength = s1.length;
    while (right < s2.length) {
        let current = s2[right];
        if (neededChar[current] > 0) requiredLength--;
        neededChar[current]--;
        right++;
        if (requiredLength === 0) return true;
        if (right - left === s1.length) {
            if (neededChar[s2[left]] >= 0) requiredLength++;
            neededChar[s2[left]]++;
            left++;
        }
    }
    return false;
};

console.log(checkInclusion("ab", "eidbaooo")); // true