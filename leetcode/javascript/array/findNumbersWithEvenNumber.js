/*  Find Numbers with Even Number of Digits */
/* Найдите числа с четным числом цифр */
/* 
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

    for (let i = 0; i <= nums.length; i++) {
        if (nums[i] && nums[i].toString().length % 2 === 0) {
            counter++;
        }
    }

    return counter;
};

const nums = [12, 345, 2, 6, 7896];
console.log(findNumbers(nums));
