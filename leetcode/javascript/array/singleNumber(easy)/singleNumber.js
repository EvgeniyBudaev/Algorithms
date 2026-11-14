/* 136. Single Number
https://leetcode.com/problems/single-number/description/

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
 * Находит единственный элемент, который не имеет пары.
 * @param {number[]} nums - Массив чисел, где все элементы парные, кроме одного
 * @return {number} - Единственное уникальное число
 * @time O(n) - один проход по массиву
 * @space O(1) - используется только одна переменная-аккумулятор
 */
var singleNumber = function(nums) {
    let result = 0;

    for (let i = 0; i < nums.length; i++) {
        result ^= nums[i];
    }

    return result;
};

console.log(singleNumber([2,2,1])); // 1