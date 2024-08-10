/* https://leetcode.com/problems/maximum-average-subarray-i/description/
Fixed window size

Вам дан целочисленный массив nums, состоящий из n элементов, и целое число k.
Найдите непрерывный подмассив длиной k, имеющий максимальное среднее значение, и верните это значение.

Input: nums = [1,12,-5,-6,50,3], k = 4
Output: 12.75000
Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75

Input: nums = [5], k = 1
Output: 5.00000
*/

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var findMaxAverage = function(nums, k) {
    let sum = 0;

    for (let i = 0; i < k; i++) {
        sum += nums[i];
    }

    let maxAvg = sum / k;

    for (let i = k; i < nums.length; i++) {
        sum +=  nums[i] - nums[i - k];
        const avg = sum / k;
        maxAvg = Math.max(maxAvg, avg);
    }

    return maxAvg;
};

console.log(findMaxAverage([1,12,-5,-6,50,3], 4)); // 12.75