/* Учитывая строку s и целое число k, верните длину самой длинной подстроки s, содержащей не более k различных символов.

Input: s = "eceba", k = 2
Output: 3
Объяснение: Это подстрока «ece» длиной 3.

Input: s = "aa", k = 1
Output: 2
Объяснение: Это подстрока «aa» длиной 2.
*/

function lengthOfLongestSubstringKDistinct(s, k) {
    const cnt = new Map();
    let l = 0;
    for (const c of s) {
        cnt.set(c, (cnt.get(c) ?? 0) + 1);
        if (cnt.size > k) {
            cnt.set(s[l], cnt.get(s[l]) - 1);
            if (cnt.get(s[l]) === 0) cnt.delete(s[l]);
            l++;
        }
    }
    return s.length - l;
}

console.log(lengthOfLongestSubstringKDistinct("eceba", 2)); // 3