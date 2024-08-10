/* https://leetcode.com/problems/max-consecutive-ones-iii/description/

Учитывая двоичный массив nums и целое число k, верните максимальное количество последовательных единиц в массиве,
если вы можете перевернуть не более k нулей.

Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
Output: 6
Explanation: [1,1,1,0,0,1,1,1,1,1,1]
Числа, выделенные жирным шрифтом, были перевернуты с 0 на 1. Самый длинный подмассив подчеркнут.

Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
Output: 10
Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
*/

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var longestOnes = function(nums, k) {
    let left = 0, zeroCount = 0, ans = 0;

    for (let right = 0; right < nums.length; right++) {
        if (nums[right] === 0) zeroCount++;
        while (zeroCount > k) {
            if (nums[left] === 0) zeroCount--;
            left++;
        }
        ans = Math.max(ans, right - left + 1);
    }

    return ans;
};

console.log(longestOnes([1,1,1,0,0,0,1,1,1,1,0], 2)); // 6