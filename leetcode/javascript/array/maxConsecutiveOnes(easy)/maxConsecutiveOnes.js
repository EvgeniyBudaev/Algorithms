/* https://leetcode.com/problems/max-consecutive-ones/description/

Для двоичного массива nums вернуть максимальное количество последовательных 1 в массиве.

Example 1
Input: nums = [1,1,0,1,1,1]
Output: 3

Example 2
Input: nums = [1,0,1,1,0,1]
Output: 2
*/

/**
 * @param {number[]} nums
 * @return {number}
 */
const findMaxConsecutiveOnes = function(nums) {
    let count = 0, max = 0;
    for (let i = 0; i < nums.length; i++) {
        if (nums[i] === 1) {
            count++;
            if (count > max) max = count;
        } else count = 0;
    }
    return max;
};

console.log("result: ", findMaxConsecutiveOnes([1, 1, 0, 1, 1, 1]));
