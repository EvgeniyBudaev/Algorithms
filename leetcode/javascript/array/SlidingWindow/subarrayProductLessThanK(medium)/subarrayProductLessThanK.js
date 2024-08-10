/* https://leetcode.com/problems/subarray-product-less-than-k/description/

Учитывая массив положительных целых чисел nums и целое число k, верните количество подмассивов, в которых произведение
всех элементов в подмассиве строго меньше k.

Input: nums = [10,5,2,6], k = 100
Output: 8
Пояснение: 8 подмассивов с произведением меньше 100:
[10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6]
Обратите внимание, что [10, 5, 2] не включены, поскольку произведение 100 не строго меньше k.

Input: nums = [1,2,3], k = 0
Output: 0
*/

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var numSubarrayProductLessThanK = function(nums, k) {
    if (k <= 1) return 0;
    let left = 0, curr = 1, ans = 0;

    for (let right = 0; right < nums.length; right++) {
        curr *= nums[right];
        while (curr >= k) {
            curr /= nums[left];
            left++;
        }
        ans += right - left + 1;
    }

    return ans;
};

console.log(numSubarrayProductLessThanK([10, 5, 2, 6], 100)); // 8