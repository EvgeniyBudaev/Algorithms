/*
Учитывая массив положительных целых чисел nums и целое число k, верните количество подмассивов, в которых произведение
всех элементов в подмассиве строго меньше k.

Например, учитывая входные числа = [10, 5, 2, 6], k = 100, ответ — 8. Подмассивы с продуктами меньше k:

[10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6]
 */

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var numSubarrayProductLessThanK = function(nums, k) {
    if (k <= 1) {
        return 0;
    }

    let ans = 0, left = 0, curr = 1;

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