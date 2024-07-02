// https://leetcode.com/problems/max-consecutive-ones/description/
/* Max Consecutive Ones */
/* Максимальное количество последовательных единиц */
/*
Для двоичного массива nums вернуть максимальное количество последовательных 1 в массиве.

Example 1
Input: nums = [1,1,0,1,1,1]
Output: 3

Example 2
Input: nums = [1,0,1,1,0,1]
Output: 2


Ограничения:
1 <= nums.length <= 105
nums [i] равно 0 или 1.
*/

/**
 * @param {number[]} nums
 * @return {number}
 */
const findMaxConsecutiveOnes = function(nums) {
    let counter = 0, quantityMaxOne = 0;

    for (let i = 0; i < nums.length; i++) {
        if (nums[i] === 1) {
            counter++;
            if (counter > quantityMaxOne) {
                quantityMaxOne = counter;
            }
        } else {
            counter = 0;
        }
    }

    return quantityMaxOne;
};

const nums = [1, 1, 0, 1, 1, 1];
console.log("result: ", findMaxConsecutiveOnes(nums));
