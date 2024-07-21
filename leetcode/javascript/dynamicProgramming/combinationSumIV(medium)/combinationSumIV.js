/* https://leetcode.com/problems/combination-sum-iv/description/

Учитывая массив различных целых чисел nums и целевое целочисленное значение, верните количество возможных комбинаций,
которые в сумме составляют целевое значение.
Тестовые случаи генерируются таким образом, чтобы ответ мог уместиться в 32-битное целое число.

Input: nums = [1,2,3], target = 4
Output: 7
Explanation:
Возможные способы комбинирования:
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
Обратите внимание, что разные последовательности считаются разными комбинациями.

Input: nums = [9], target = 3
Output: 0
*/

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var combinationSum4 = function(nums, target) {
    const dp = Array(target + 1).fill(0);
    dp[0] = 1;

    for (let i = 1; i <= target; i++) {
        for (const num of nums) {
            if (i - num >= 0) {
                dp[i] += dp[i - num];
            }
        }
    }

    return dp[target];
};

const nums = [1,2,3];
const target = 4;
console.log(combinationSum4(nums, target));