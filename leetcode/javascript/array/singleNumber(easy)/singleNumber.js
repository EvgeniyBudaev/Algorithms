/* https://leetcode.com/problems/single-number/description/

Учитывая непустой массив целых чисел nums, каждый элемент появляется дважды, кроме одного. Найдите этого единственного.
Вы должны реализовать решение с линейной сложностью времени выполнения и использовать только постоянное дополнительное
пространство.

Input: nums = [2,2,1]
Output: 1

Input: nums = [4,1,2,1,2]
Output: 4

Input: nums = [1]
Output: 1
*/

/**
 * @param {number[]} nums
 * @return {number}
 */
var singleNumber = function(nums) {
    nums.sort((a, b) => a - b); // [1, 2, 2]

    for(let i  = 0; i < nums.length - 1; i += 2) {
        if (nums[i] !== nums[i + 1]) {
            return nums[i];
        }
    }

    return nums[nums.length - 1];
};

console.log(singleNumber([2,2,1])); // 1