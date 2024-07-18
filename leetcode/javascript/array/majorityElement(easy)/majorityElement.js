/* https://leetcode.com/problems/majority-element/description/

Учитывая массив nums размера n, верните элемент большинства.
Мажоритарным элементом является элемент, который появляется более ⌊n / 2⌋ раз. Вы можете предположить, что элемент
большинства всегда существует в массиве.

Input: nums = [3,2,3]
Output: 3

Input: nums = [2,2,1,1,1,2,2]
Output: 2

Input: nums = [3,3,4]
Output: 3
*/

/**
 * @param {number[]} nums
 * @return {number}
 */
var majorityElement = function(nums) {
    let count = 0;
    let candidate = 0;

    for (let i = 0; i < nums.length; i++) {
        if (count === 0) {
            candidate = nums[i];
        }
        count += candidate === nums[i] ? 1 : -1;
    }

    return candidate;
};

console.log(majorityElement([3,2,3])); // 3
console.log(majorityElement([2,2,1,1,1,2,2])); // 2
console.log(majorityElement([3,3,4])); // 3