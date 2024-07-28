/* https://leetcode.com/problems/sort-colors/description/

Дан массив nums с n объектами красного, белого или синего цвета, отсортируйте их по месту так, чтобы объекты одного
цвета были соседними, а цвета располагались в порядке красного, белого и синего.
Мы будем использовать целые числа 0, 1 и 2 для обозначения красного, белого и синего цветов соответственно.
Вы должны решить эту проблему, не используя функцию сортировки библиотеки.

Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]

Input: nums = [2,0,1]
Output: [0,1,2]
*/

/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var sortColors = function(nums) {
    let left = 0, right = nums.length - 1;
    let i = 0;

    while (left < right && i <= right) {
        if (nums[i] === 0) {
            [nums[left], nums[i]] = [nums[i], nums[left]];
            left++;
            i++;
        } else if (nums[i] === 2) {
            [nums[right], nums[i]] = [nums[i], nums[right]];
            right--;
        } else {
            i++;
        }
    }
};

var sortColorsV2 = function(nums) {
    nums.sort((a, b) => a - b);
};

sortColors([2,0,2,1,1,0]); // [0,0,1,1,2,2]