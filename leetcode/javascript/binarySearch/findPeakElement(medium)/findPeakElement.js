/* https://leetcode.com/problems/find-peak-element/description/

Пиковый элемент — это элемент, который строго больше своих соседей.
Учитывая целочисленный массив с нулевым индексом, найдите пиковый элемент и верните его индекс. Если массив содержит
несколько пиков, верните индекс любого из пиков.
Вы можете представить, что nums[-1] = nums[n] = -∞. Другими словами, элемент всегда считается строго большим, чем сосед,
находящийся за пределами массива.
Вы должны написать алгоритм, который работает за время O(log n).

Input: nums = [1,2,3,1]
Output: 2
Объяснение: 3 — это пиковый элемент, и ваша функция должна возвращать индексный номер 2.

Input: nums = [1,2,1,3,5,6,4]
Output: 5
Пояснение: Ваша функция может возвращать либо индексный номер 1, где пиковый элемент равен 2, либо индексный номер 5,
где пиковый элемент равен 6.
*/

/**
 * @param {number[]} nums
 * @return {number}
 */
var findPeakElement = function(nums) {
    let l = 0, r = nums.length - 1;
    while (l < r) {
        const mid = Math.floor((l + r) / 2);
        if (nums[mid + 1] > nums[mid]) l = mid + 1;
        else r = mid;
    }
    return l;
};

console.log(findPeakElement([1,2,3,1])); // 2