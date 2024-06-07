// Fixed window size

/*
Учитывая целочисленный массив nums и целое число k, найдите сумму подмассива с наибольшей суммой, длина которой равна k.
 */

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var findBestSubarray = function(nums, k) {
    let curr = 0;
    for (let i = 0; i < k; i++) {
        curr += nums[i];
    }

    let ans = curr;
    for (let i = k; i < nums.length; i++) {
        curr += nums[i] - nums[i - k];
        ans = Math.max(ans, curr);
    }

    return ans;
}

console.log(findBestSubarray([3, -1, 4, 12, -8, 5, 6], 4)); // 18 <- [3, -1, 4, 12,]