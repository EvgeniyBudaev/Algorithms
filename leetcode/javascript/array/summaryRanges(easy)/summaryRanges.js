/* https://leetcode.com/problems/summary-ranges/description/

Вам дан отсортированный уникальный целочисленный массив чисел.
Диапазон [a,b] — это набор всех целых чисел от a до b (включительно).

Верните наименьший отсортированный список диапазонов, который точно охватывает все числа в массиве. То есть каждый
элемент nums охватывается ровно одним из диапазонов, и не существует целого числа x такого, что x находится в одном из
диапазонов, но не в nums.

Каждый диапазон [a,b] в списке должен выводиться как:
"a->b" if a != b
"a" if a == b

Input: nums = [0,1,2,4,5,7]
Output: ["0->2","4->5","7"]
Пояснение: Диапазоны:
[0,2] --> "0->2"
[4,5] --> "4->5"
[7,7] --> "7"
*/

/**
 * @param {number[]} nums
 * @return {string[]}
 */
var summaryRanges = function(nums) {
    let result = [];
    let start = nums[0];
    for(let i = 1; i <= nums.length; i++) {
        if(nums[i] - nums[i-1] === 1) continue;
        if(start === nums[i-1]) result.push(`${start}`);
        else result.push(`${start}->${nums[i-1]}`);
        start = nums[i];
    }
    return result;
};

console.log(summaryRanges([0,1,2,4,5,7])); // ["0->2","4->5","7"]