/* https://leetcode.com/problems/find-numbers-with-even-number-of-digits/description/

Найдите числа с четным числом цифр.
Дан массив целых чисел, верните, сколько из них содержат четное количество цифр.

Example 1:
Input: nums = [12,345,2,6,7896]
Output: 2

Example 2:
Input: nums = [555,901,482,1771]
Output: 1 
*/

/**
 * @param {number[]} nums
 * @return {number}
 */
const findNumbers = function(nums) {
    let counter = 0;
    nums.forEach(num => {
        if (num.toString().length % 2 === 0) counter++;
    })
    return counter;
};

const nums = [12, 345, 2, 6, 7896]; // 2
console.log(findNumbers(nums));
