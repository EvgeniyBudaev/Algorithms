/* Fixed window size

Учитывая целочисленный массив nums и целое число k, найдите сумму подмассива с наибольшей суммой, длина которой равна k.
*/

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
// time complexity: O(n), space complexity: O(1)
var findMaxSumSubarrayLengthK = function(nums, k) {
    let curr = 0; // Some data to track the window
    for (let i = 0; i < k; i++) {
        curr += nums[i];
    }

    let ans = curr;
    for (let i = k; i < nums.length; i++) {
        curr += nums[i] - nums[i - k]; // Add arr[i] to window and Remove arr[i - k] from window
        ans = Math.max(ans, curr); // Update ans
    }

    return ans;
}

console.log(findMaxSumSubarrayLengthK([3, -1, 4, 12, -8, 5, 6], 4)); // 18 <- [3, -1, 4, 12,]