/* https://leetcode.com/problems/minimum-height-trees/description/

Проверьте, можно ли однозначно восстановить исходную последовательность org из последовательностей в seqs.
Последовательность org представляет собой перестановку целых чисел от 1 до n, где 1 ≤ n ≤ 104.
Реконструкция означает построение кратчайшей общей суперпоследовательности последовательностей в seqs (т. е. кратчайшей
последовательности, чтобы все последовательности в seqs были ее подпоследовательностью).  Определите, существует ли
только одна последовательность, которую можно восстановить из seqs, и это последовательность org.

Input: org: [1,2,3], seqs: [[1,2],[1,3]]
Output: false
Explanation:
[1,2,3] — не единственная последовательность, которую можно восстановить, поскольку [1,3,2] также является допустимой
последовательностью, которую можно восстановить.
*/

function sequenceReconstruction(nums, sequences) {
    const n = nums.length;
    const g = Array.from({ length: n }, () => []);
    const indeg = Array(n).fill(0);
    for (const seq of sequences) {
        for (let i = 1; i < seq.length; ++i) {
            const [a, b] = [seq[i - 1] - 1, seq[i] - 1];
            g[a].push(b);
            ++indeg[b];
        }
    }
    const q = indeg.map((v, i) => (v === 0 ? i : -1)).filter(v => v !== -1);
    while (q.length === 1) {
        const i = q.pop();
        for (const j of g[i]) {
            if (--indeg[j] === 0) q.push(j);
        }
    }
    return q.length === 0;
}

const org = [1,2,3];
const seqs = [[1,2],[1,3]];
console.log(sequenceReconstruction(org, seqs)); // false