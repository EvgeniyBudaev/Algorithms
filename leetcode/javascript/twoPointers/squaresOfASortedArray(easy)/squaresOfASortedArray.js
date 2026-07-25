/* 977. Squares of a Sorted Array
https://leetcode.com/problems/squares-of-a-sorted-array/description/

Учитывая целочисленный массив nums, отсортированный в неубывающем порядке, верните массив квадратов каждого числа,
отсортированного в неубывающем порядке.

Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].

Input: nums = [-7,-3,2,3,11]
Output: [4,9,9,49,121]

Two pointers
Time complexity: O(n)
Space complexity: O(n)
*/

/**
 * @param {number[]} nums
 * @return {number[]}
 */
var sortedSquares = function (nums) {
    const n = nums.length;
    const result = new Array(n);

    let p1 = 0;          // Указатель на начало (самые маленькие отрицательные числа)
    let p2 = n - 1;      // Указатель на конец (самые большие положительные числа)

    // Заполняем результирующий массив с конца (от самых больших квадратов к самым маленьким)
    for (let i = n - 1; p1 <= p2; i--) {
        const abs1 = Math.abs(nums[p1]);
        const abs2 = Math.abs(nums[p2]);

        if (abs1 > abs2) {
            result[i] = abs1 * abs1;
            p1++;
        } else {
            result[i] = abs2 * abs2;
            p2--;
        }
    }

    return result;
};