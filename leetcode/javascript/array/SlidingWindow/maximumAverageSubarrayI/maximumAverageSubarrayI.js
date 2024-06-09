// Fixed window size

/*
Вам дан целочисленный массив nums, состоящий из n элементов, и целое число k.

Найдите непрерывный подмассив длиной k, имеющий максимальное среднее значение, и верните это значение.
Принимается любой ответ с ошибкой расчета менее 10-5.
 */

/*
Input: nums = [1,12,-5,-6,50,3], k = 4
Output: 12.75000
Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75
 */

/*
Input: nums = [5], k = 1
Output: 5.00000
 */

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var findMaxAverage = function(nums, k) {
    let currSum = 0;

    for (let i = 0; i < k; i++) {
        currSum += nums[i];
    }

    let maxAvg = currSum / k;

    for (let i = k; i < nums.length; i++) {
        currSum +=  nums[i] - nums[i - k];
        const avg = currSum / k;
        maxAvg = Math.max(maxAvg, avg);
    }

    return maxAvg;
};

console.log(findMaxAverage([1,12,-5,-6,50,3], 4)); // 12.75