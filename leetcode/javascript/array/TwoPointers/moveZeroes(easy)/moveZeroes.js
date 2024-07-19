/* https://leetcode.com/problems/move-zeroes/description/

Учитывая целочисленный массив nums, переместите все 0 в его конец, сохраняя относительный порядок ненулевых элементов.
Обратите внимание, что вы должны сделать это на месте, не копируя массив.

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

Input: nums = [0]
Output: [0]
*/

/**
 * @param {number[]} nums
 * @return {number[]} Do not return anything, modify nums in-place instead.
 */
var moveZeroes = function(nums) {
    let left = 0;

    for (let right = 0; right < nums.length; right++) {
        if (nums[right] !== 0) {
            [nums[right], nums[left]] = [nums[left], nums[right]];
            left++;
        }
    }

    return nums;
};

console.log(moveZeroes([0,1,0,3,12])); // [1,3,12,0,0]
