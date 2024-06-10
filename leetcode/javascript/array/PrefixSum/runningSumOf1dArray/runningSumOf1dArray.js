// https://www.youtube.com/watch?v=at9xNhYlXJ4&ab_channel=DevGym Объяснение Prefix Sum
/*
Дан массив чисел. Мы определяем текущую сумму массива как RunningSum[i] = sum(nums[0]…nums[i]).
Верните текущую сумму чисел.
 */

/*
Input: nums = [1,2,3,4]
Output: [1,3,6,10]
Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].
 */

/*
Input: nums = [1,1,1,1,1]
Output: [1,2,3,4,5]
Explanation: Running sum is obtained as follows: [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1].
 */

/*
Input: nums = [3,1,2,10,1]
Output: [3,4,6,16,17]
 */

/**
 * @param {number[]} nums
 * @return {number[]}
 */
var runningSum = function(nums) {
    for (let i = 1; i < nums.length; i++) {
        nums[i] = nums[i] + nums[i - 1];
    }

    return nums;
};

console.log(runningSum([1,2,3,4])); // [1,3,6,10]