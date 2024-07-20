/* https://leetcode.com/problems/combination-sum/description/

Учитывая массив различных целых чисел-кандидатов и целевую целочисленную цель, верните список всех уникальных комбинаций
кандидатов, в которых сумма выбранных чисел равна целевому значению. Вы можете возвращать комбинации в любом порядке.

Одно и то же число может быть выбрано из кандидатов неограниченное количество раз. Две комбинации уникальны, если
частота хотя бы одно из выбранных чисел отличается.

Тестовые примеры генерируются таким образом, чтобы количество уникальных комбинаций, дающих целевое значение, составляло
менее 150 комбинаций для данного входного сигнала.

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Объяснение:
2 и 3 являются кандидатами, а 2 + 2 + 3 = 7. Обратите внимание, что 2 можно использовать несколько раз.
7 — кандидат, а 7 = 7.
Это единственные две комбинации.
*/

/**
 * @param {number[]} candidates
 * @param {number} target
 * @return {number[][]}
 */
var combinationSum = function(candidates, target) {
    const res = [];

    function makeCombination(idx, comb, total) {
        if (total === target) {
            res.push([...comb]);
            return;
        }

        if (total > target || idx >= candidates.length) return;

        comb.push(candidates[idx]);
        makeCombination(idx, comb, total + candidates[idx]);
        comb.pop();
        makeCombination(idx + 1, comb, total);
    }

    makeCombination(0, [], 0);
    return res;
};

const candidates = [2,3,6,7];
const target = 7;
console.log(combinationSum(candidates, target));