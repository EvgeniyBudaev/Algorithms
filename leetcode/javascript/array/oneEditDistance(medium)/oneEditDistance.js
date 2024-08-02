/* https://github.com/EvgeniyBudaev/doocs_leetcode/tree/main/solution/0100-0199/0161.One%20Edit%20Distance

Учитывая две строки s и, вернитесь, t если их расстояние редактирования равно, иначе верните .1 true false
Возможны три ситуации, в которых расстояние редактирования между строкой s и строкой равно 1:t

s Вставьте в него ровно один символ , чтобы получить t
s Удалите из него ровно один символ, чтобы получить t
Замена ровно одного символа другим символом в s дает t

Input: s = "ab", t = "acb"
Output: true
Объяснение: в строку s можно вставить 'c',  чтобы получить t.
*/

function isOneEditDistance(s, t) {
    const [m, n] = [s.length, t.length];
    if (m < n) return isOneEditDistance(t, s);
    if (m - n > 1) return false;
    for (let i = 0; i < n; ++i) {
        if (s[i] !== t[i]) return s.slice(i + 1) === t.slice(i + (m === n ? 1 : 0));
    }
    return m === n + 1;
}

console.log(isOneEditDistance("ab", "acb")); // true