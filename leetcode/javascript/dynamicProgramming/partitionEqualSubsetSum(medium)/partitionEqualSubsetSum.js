/* https://leetcode.com/problems/partition-equal-subset-sum/description/


Учитывая числа целочисленного массива, верните true, если вы можете разделить массив на два подмножества так, чтобы
сумма элементов в обоих подмножествах была равна или false в противном случае.

Input: nums = [1,5,11,5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11].

Input: nums = [1,2,3,5]
Output: false
Explanation: The array cannot be partitioned into equal sum subsets.
*/

/**
 * @param {number[]} nums
 * @return {boolean}
 */
var canPartition = function(nums) {
    const totalSum = nums.reduce((sum, num) => sum + num, 0);

    if (totalSum % 2 !== 0) return false;

    const target = totalSum / 2;
    const dp = Array(target + 1).fill(false);
    dp[0] = true;

    for (const num of nums) {
        for (let j = target; j >= num; j--) {
            if (dp[j - num]) {
                dp[j] = true;
            }
        }
    }

    return dp[target];
};

const nums = [1,5,11,5];
console.log(canPartition(nums)); // true