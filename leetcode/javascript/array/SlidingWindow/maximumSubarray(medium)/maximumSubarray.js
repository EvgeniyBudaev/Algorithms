/* https://leetcode.com/problems/maximum-subarray/description/
Учитывая целочисленный массив чисел, найдите подмассив с наибольшей суммой и вернуть ее сумму.

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: Подмассив [4,-1,2,1] имеет наибольшую сумму 6.

Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.

Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
*/

/**
 * @param {number[]} nums
 * @return {number}
 */
var maxSubArray = function(nums) {
    let maxSum = nums[0];
    let currentSum = maxSum;

    for(let i = 1; i < nums.length; i++) {
        currentSum = Math.max(currentSum + nums[i], nums[i]);
        maxSum = Math.max(maxSum, currentSum);
    }

    return maxSum;
};

console.log(maxSubArray([-2,1,-3,4,-1,2,1,-5,4])); // 6