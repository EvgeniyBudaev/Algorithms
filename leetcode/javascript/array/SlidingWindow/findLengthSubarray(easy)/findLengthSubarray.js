/*
Дан массив целых положительных чисел nums и целое число k. Найдите длину самого длинного подмассива, сумма которого
меньше или равна k.

Input: nums = [3, 1, 2, 7, 4, 2, 1, 1, 5], k = 8
Output: 4
*/

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var findLengthSubarray = function(nums, k) {
    let left = 0, sum = 0, ans = 0;

    for (let right = 0; right < nums.length; right++) {
        sum += nums[right];
        while (sum > k) {
            sum -= nums[left];
            left++;
        }
        ans = Math.max(ans, right - left + 1);
    }

    return ans;
}

console.log(findLengthSubarray([3, 1, 2, 7, 4, 2, 1, 1, 5], 8)); // 4