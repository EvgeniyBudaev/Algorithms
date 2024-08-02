/* https://leetcode.com/problems/subarray-sum-equals-k/description/

Учитывая массив целых чисел nums и целое число k, верните общее количество подмассивов, сумма которых равна k.
Подмассив — это непрерывная непустая последовательность элементов массива.

Input: nums = [1,1,1], k = 2
Output: 2

Input: nums = [1,2,3], k = 3
Output: 2
*/

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var subarraySum = function(nums, k) {
    const map = new Map();
    map.set(0, 1);
    let sum = 0;
    let count = 0;

    nums.forEach((num) => {
        sum += num;
        if(map.has(sum - k)) count += map.get(sum - k);
        map.set(sum, map.has(sum) ? map.get(sum) + 1 : 1);
    });

    return count;
};

console.log(subarraySum([1,1,1], 2)); // 2