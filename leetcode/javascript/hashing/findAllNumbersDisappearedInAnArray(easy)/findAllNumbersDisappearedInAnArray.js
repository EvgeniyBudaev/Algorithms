/* https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/

Учитывая массив nums из n целых чисел, где nums[i] находится в диапазоне [1, n], верните массив всех целых чисел в
диапазоне [1, n], которые не встречаются в nums.

Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]

Input: nums = [1,1]
Output: [2]
*/

/**
 * @param {number[]} nums
 * @return {number[]}
 */
var findDisappearedNumbers = function(nums) {
    const hashmap = {};
    const output = [];

    for (let i = 0; i < nums.length; i++) {
        hashmap[nums[i]] = 1;
    }

    for (let i = 1; i <= nums.length; i++) {
        if (!hashmap[i]) output.push(i);
    }

    return output;
};

console.log(findDisappearedNumbers([4,3,2,7,8,2,3,1])); // [5,6]