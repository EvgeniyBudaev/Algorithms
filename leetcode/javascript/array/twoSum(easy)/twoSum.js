/* 1. Two Sum
https://leetcode.com/problems/two-sum/description/

Учитывая массив целых чисел nums и целочисленную target, верните индексы двух чисел так, чтобы их сумма составляла
target. Вы можете предположить, что каждый вход будет иметь ровно одно решение, и вы не можете использовать один и тот
же элемент дважды. Вы можете вернуть ответ в любом порядке.

Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]
*/

/**
 * Находит индексы двух чисел, сумма которых равна target.
 * @param {number[]} nums - Массив целых чисел
 * @param {number} target - Целевая сумма
 * @return {number[]} - Массив из двух индексов
 * @time O(n) - один проход по массиву, операции с Map занимают O(1) в среднем
 * @space O(n) - в худшем случае Map хранит n элементов
 */
var twoSum = function(nums, target) {
    const map = new Map();

    for (let i = 0; i < nums.length; i++) {
        const num = nums[i];
        const diff = target - num;
        if (map.has(diff)) return [map.get(diff), i];
        map.set(num, i);
    }
};

console.log(twoSum([3, 2, 4], 6)); // [1, 2]
