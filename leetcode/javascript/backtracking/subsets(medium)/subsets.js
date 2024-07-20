/* https://leetcode.com/problems/subsets/description/

Учитывая целочисленный массив чисел уникальных элементов, верните все возможные подмножества (набор мощности).
Набор решений не должен содержать повторяющихся подмножеств. Верните решение в любом порядке.

Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Input: nums = [0]
Output: [[],[0]]
*/

/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var subsets = function(nums) {
    const res = [];

    const dfs = (i, sub) => {
        if (i === nums.length) return res.push(sub);
        dfs(i + 1, [...sub, nums[i]]);
        dfs(i + 1, sub);
    }

    dfs(0, []);
    return res;
};

console.log(subsets([1,2,3])); // [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]