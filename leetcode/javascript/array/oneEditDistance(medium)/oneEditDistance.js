/* 161. One Edit Distance
https://github.com/EvgeniyBudaev/doocs_leetcode/tree/main/solution/0100-0199/0161.One%20Edit%20Distance

Учитывая две строки s и t, определите, находятся ли они на расстоянии одного редактирования друг от друга.

s Вставьте в него ровно один символ , чтобы получить t
s Удалите из него ровно один символ, чтобы получить t
Замена ровно одного символа другим символом в s дает t

Input: s = "ab", t = "acb"
Output: true
Объяснение: в строку s можно вставить 'c',  чтобы получить t.

Input: s = "cab", t = "ad"
Output: false
*/

function isOneEditDistance(s, t) {
    const [m, n] = [s.length, t.length];
    if (m < n) return isOneEditDistance(t, s); // "acb", "ab"
    if (m - n > 1) return false;
    for (let i = 0; i < n; i++) {
        if (s[i] !== t[i]) return s.slice(i + 1) === t.slice(i + (m === n ? 1 : 0));
    }
    return m === n + 1;
}

console.log(isOneEditDistance("ab", "acb")); // true