/* https://leetcode.com/problems/binary-search/description/

Учитывая массив целых чисел nums, отсортированный в порядке возрастания, и целочисленную цель, напишите функцию для
поиска цели в nums. Если цель существует, верните ее индекс. В противном случае верните -1.
Вы должны написать алгоритм со сложностью выполнения O(log n).

Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4

Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1
*/

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var search = function(nums, target) {
    let low = 0, high = nums.length - 1;

    while (low <= high) {
        const mid = Math.floor((low + high) / 2);
        const guess = nums[mid];

        if (guess === target) return mid;
        if (guess > target) high = mid - 1;
        else low = mid + 1;
    }

    return -1;
};

console.log(search([-1,0,3,5,9,12], 9)); // 4