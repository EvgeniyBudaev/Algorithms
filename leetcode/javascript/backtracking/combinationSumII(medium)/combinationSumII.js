/* https://leetcode.com/problems/combination-sum-ii/description/

Учитывая набор номеров кандидатов (кандидатов) и целевое число (цель), найдите все уникальные комбинации среди
кандидатов, в которых сумма чисел кандидатов равна целевому значению.
Каждое число кандидатов можно использовать в комбинации только один раз.
Примечание. Набор решений не должен содержать повторяющихся комбинаций.

Input: candidates = [10,1,2,7,6,1,5], target = 8
Output: [[1,1,6], [1,2,5], [1,7], [2,6]]
*/

/**
 * @param {number[]} candidates
 * @param {number} target
 * @return {number[][]}
 */
var combinationSum2 = function(candidates, target) {
    candidates.sort((a, b) => a - b);
    const res = [];
    function dfs(target, start, comb) {
        if (target < 0) return;
        if (target === 0) {
            res.push(comb.slice());
            return;
        }
        for (let i = start; i < candidates.length; i++) {
            if (i > start && candidates[i] === candidates[i-1]) continue;
            if (candidates[i] > target) break;
            dfs(target - candidates[i], i + 1, comb.concat(candidates[i]));
        }
    }
    dfs(target, 0, []);
    return res;
};

const candidates = [10,1,2,7,6,1,5];
const target = 8;
console.log(combinationSum2(candidates, target)); // [[1,1,6], [1,2,5], [1,7], [2,6]]