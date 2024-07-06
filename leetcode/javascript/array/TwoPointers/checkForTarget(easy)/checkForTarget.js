/*
Учитывая отсортированный массив уникальных целых чисел и целевое целое число, верните true, если существует пара чисел,
сумма которых равна целевому значению, и false в противном случае. Эта задача аналогична задаче «Две суммы».
(В режиме Two Sum входные данные не сортируются).

Например, если заданы числа = [1, 2, 4, 6, 8, 9, 14, 15] и цель = 13, верните true, потому что 4 + 9 = 13.
 */

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {boolean}
 */
var checkForTarget = function(nums, target) {
    let left = 0, right = nums.length - 1;

    while (left < right) {
        // curr is the current sum
        let curr = nums[left] + nums[right];
        if (curr === target) return true;

        if (curr > target) {
            right--;
        } else {
            left++;
        }
    }

    return false;
}

console.log(checkForTarget([1, 2, 4, 6, 8, 9, 14, 15], 13)); // true