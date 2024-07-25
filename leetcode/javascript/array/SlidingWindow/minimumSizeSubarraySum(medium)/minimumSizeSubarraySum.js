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
    let start = 0, end = 0;
    let minValue = Infinity;
    let subarraySum = nums[0];
    while (start <= end && end < nums.length) {
        if(subarraySum >= target){
            minValue = Math.min(minValue, end - start + 1);
            subarraySum -= nums[start];
            start++;
        } else {
            end++;
            subarraySum += nums[end];
        }
    }
    return minValue === Infinity ? 0 : minValue;
};

const target = 7;
const nums = [2,3,1,2,4,3];
console.log(minSubArrayLen(target, nums)); // 2