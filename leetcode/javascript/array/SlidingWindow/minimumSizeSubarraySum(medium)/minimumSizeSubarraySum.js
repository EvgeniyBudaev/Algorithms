/* https://leetcode.com/problems/minimum-size-subarray-sum/description/
solution https://leetcode.com/problems/minimum-size-subarray-sum/solutions/2657137/sliding-window-dynamic-approach-o-n-o-n-k-javascript/

Учитывая массив положительных целых чисел nums и цель положительного целого числа, верните минимальную длину
подмассив сумма которых больше или равна целевой. Если такого подмассива нет, вместо него верните 0.

Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Пояснение: Подмассив [4,3] имеет минимальную длину при условии ограничения задачи

Input: target = 4, nums = [1,4,4]
Output: 1

Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0
*/

/**
 * @param {number} target
 * @param {number[]} nums
 * @return {number}
 */
var minSubArrayLen = function(target, nums) {
    let left = 0, right = 0;
    let minLength = Infinity;
    let subarraySum = nums[0];

    while (right < nums.length) {
        if(subarraySum >= target) {
            minLength = Math.min(minLength, right - left + 1);
            subarraySum -= nums[left];
            left++;
        } else {
            right++;
            subarraySum += nums[right];
        }
    }

    return minLength === Infinity ? 0 : minLength;
};

console.log(minSubArrayLen(7, [2,3,1,2,4,3])); // 2