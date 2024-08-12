/* https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/

Предположим, что массив длины n, отсортированный в порядке возрастания, поворачивается от 1 до n раз. Например, массив
nums = [0,1,2,4,5,6,7] может выглядеть следующим образом:

[4,5,6,7,0,1,2], если его повернули 4 раза.
[0,1,2,4,5,6,7], если его повернули 7 раз.
Обратите внимание, что вращение массива [a[0], a[1], a[2], ..., a[n-1]] 1 раз приводит к созданию массива
[a[n-1], a[0] , а[1], а[2], ..., а[n-2]].

Учитывая отсортированные числа уникальных элементов повернутого массива, верните минимальный элемент этого массива.
Вы должны написать алгоритм, который работает за время O(log n).

Input: nums = [3,4,5,1,2]
Output: 1
Пояснение: Исходный массив был [1,2,3,4,5] повёрнут 3 раза.
*/

/**
 * @param {number[]} nums
 * @return {number}
 */
var findMin = function(nums) {
    nums.sort((a, b) => a - b);
    return nums[0];
};

var findMinBinarySearch = function(nums) {
    let left = 0, right = nums.length - 1;

    while (left < right) { // l: 0, r: 4 -> l: 3, r: 4
        const mid = Math.floor((left + right) / 2); // 2 -> 3
        const guess = nums[mid]; // 5 -> 1
        const [leftNum, rightNum] = [nums[left], nums[right]]; // l: 3, r: 2 -> l: 1, r: 2

        if (leftNum < rightNum) return leftNum;
        if (leftNum <= guess) left = mid + 1;
        if (leftNum > guess) right = mid;
    }

    return nums[left];
};

console.log(findMin([3,4,5,1,2])); // 1