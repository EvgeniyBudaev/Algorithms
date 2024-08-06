/* https://leetcode.com/problems/two-sum/description/

Учитывая массив целых чисел nums и целочисленную target, верните индексы двух чисел так, чтобы их сумма составляла
target. Вы можете предположить, что каждый вход будет иметь ровно одно решение, и вы не можете использовать один и тот
же элемент дважды. Вы можете вернуть ответ в любом порядке.

Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]
*/

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    const map = new Map();
    for (let i = 0; i < nums.length; ++i) {
        const num = nums[i];
        const diff = target - num;
        if (map.has(diff)) return [map.get(diff), i];
        map.set(num, i);
    }
};

// var twoSum = function(nums, target) {
//     const mp = {};
//     for (let i = 0; i < nums.length; i++) {
//         const diff = target - nums[i];
//         if (diff in mp) return [i, mp[diff]];
//         mp[nums[i]] = i;
//     }
// };

//console.log(twoSum([2, 7, 11, 15], 9)); //  [0, 1]
console.log(twoSum([3, 2, 4], 6)); // [1, 2]
