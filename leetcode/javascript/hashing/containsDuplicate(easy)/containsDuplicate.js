/* https://leetcode.com/problems/contains-duplicate/description/
Javascript - set vs. object https://leetcode.com/problems/contains-duplicate/solutions/515531/javascript-set-vs-object/

Учитывая числа целочисленного массива, верните true, если какое-либо значение встречается в массиве хотя бы дважды, и
верните false, если каждый элемент различен.

Input: nums = [1,2,3,1]
Output: true

Input: nums = [1,2,3,4]
Output: false

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true
*/

/**
 * @param {number[]} nums
 * @return {boolean}
 */
var containsDuplicate = function(nums) {
  const obj = {};

  for (let i = 0; i < nums.length; i++) {
    if (obj[nums[i]]) return true;
    obj[nums[i]] = 1;
  }

  return false;
};

// var containsDuplicate = function(nums) {
//   const seen = new Map();
//
//   for (let num = 0; num < nums.length; num++) {
//     if (seen.has(nums[num])) return true;
//     seen.set(nums[num]);
//   }
//
//   return false;
// };

// var containsDuplicate = function(nums) {
//   let foundList = new Set();
//
//   for( let i = 0; i < nums.length; i++) {
//     if (foundList.has(nums[i])) return true;
//     foundList.add(nums[i]);
//   }
//
//   return false;
// };

// Time complexity: O(n)
// Space complexity: O(n)

// var containsDuplicate = function(nums) {
//   return new Set(nums).size !== nums.length;
// };

console.log(containsDuplicate([1,2,3,1])); // true
