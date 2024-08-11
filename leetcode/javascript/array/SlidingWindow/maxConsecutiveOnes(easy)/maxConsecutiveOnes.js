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
    let left = 0, zeroCount = 0, ans = 0;

    for (let right = 0; right < nums.length; right++) {
        if (nums[right] === 0) zeroCount++;
        while (zeroCount > 0) {
            if (nums[left] === 0) zeroCount--;
            left++;
        }
        ans = Math.max(ans, right - left + 1);
    }

    return ans;
};

const findMaxConsecutiveOnes2 = function(nums) {
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
