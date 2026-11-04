/* 442. Find All Duplicates in an Array
https://leetcode.com/problems/find-all-duplicates-in-an-array/description/

Учитывая целочисленный массив nums длины n, где все целые числа nums находятся в диапазоне [1, n] и каждое целое число
встречается один или два раза, верните массив всех целых чисел, которые встречаются дважды.
Вы должны написать алгоритм, который работает за время O(n) и использует только постоянное дополнительное пространство.

Input: nums = [4,3,2,7,8,2,3,1]
Output: [2,3]

Input: nums = [1,1,2]
Output: [1]

Input: nums = [1]
Output: []
*/

/**
 * @param {number[]} nums
 * @return {number[]}
 * @time O(n)
 * @space O(1)
 */
var findDuplicates = function(nums) {
    const result = [];

    for (let i = 0; i < nums.length; i++) {
        const num = Math.abs(nums[i]);
        const index = num - 1;

        if (nums[index] < 0) {
            result.push(num);
        } else {
            nums[index] = -nums[index];
        }
    }

    return result;
};

console.log(findDuplicates([4,3,2,7,8,2,3,1])); // [2,3]