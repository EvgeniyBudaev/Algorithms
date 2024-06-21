/*
Учитывая отсортированный массив различных целых чисел и целевое значение, верните индекс, если цель найдена.
Если нет, верните индекс там, где он был бы, если бы он был вставлен по порядку.
Вы должны написать алгоритм со сложностью выполнения O(log n)
*/

/*
Input: nums = [1,3,5,6], target = 5
Output: 2
 */

/*
Input: nums = [1,3,5,6], target = 2
Output: 1
 */

/*
Input: nums = [1,3,5,6], target = 7
Output: 4
 */

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var searchInsert = function(nums, target) {
    let low = 0;
    let high = nums.length - 1;

    while (low <= high) {
        const mid = Math.floor((low + high) / 2);
        const guess = nums[mid];

        if (guess === target) return mid;

        if (guess > target) {
            high = mid - 1;
        } else {
            low = mid + 1;
        }
    }

    return low;
};

const myList = [1,3,5,6];
console.log(searchInsert(myList, 5)); // 2