/* https://leetcode.com/problems/find-all-duplicates-in-an-array/description/

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
 */
var findDuplicates = function(nums) {
    const set = new Set();
    const result = [];
    for (let num of nums ) {
        if (set.has(num)) result.push(num);
        else set.add(num);
    }
    return result;
};

console.log(findDuplicates([4,3,2,7,8,2,3,1])); // [2,3]