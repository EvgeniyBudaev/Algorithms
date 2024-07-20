/* https://leetcode.com/problems/permutations-ii/description/

Учитывая набор чисел nums, который может содержать дубликаты, возвращает все возможные уникальные перестановки в любом
порядке.

Input: nums = [1,1,2]
Output: [[1,1,2], [1,2,1], [2,1,1]]

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/

/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var permuteUnique = function(nums) {
    const result = [];

    function backtrack(start) {
        if (start === nums.length) {
            result.push([...nums]);
            return;
        }

        const seen = new Set();

        for (let i = start; i < nums.length; i++) {
            if (seen.has(nums[i])) continue;
            seen.add(nums[i]);

            [nums[start], nums[i]] = [nums[i], nums[start]];
            backtrack(start + 1);
            [nums[start], nums[i]] = [nums[i], nums[start]];
        }
    }

    backtrack(0);
    return result;
};

console.log(permuteUnique([1,1,2])); // [[1,1,2], [1,2,1], [2,1,1]]