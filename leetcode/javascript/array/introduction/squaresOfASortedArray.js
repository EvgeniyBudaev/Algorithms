// https://leetcode.com/problems/squares-of-a-sorted-array/description/
/* Squares of a Sorted Array */

/* Two pointers
Time complexity: O(n)
Space complexity: O(n)
 */

/* 
Учитывая целочисленный массив nums, отсортированный в неубывающем порядке, вернуть массив квадратов каждого числа,
отсортированных в неубывающем порядке.

Example 1:
Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]

Example 2:
Input: nums = [-7,-3,2,3,11]
Output: [4,9,9,49,121]
*/

/**
 * @param {number[]} nums
 * @return {number[]}
 */
const sortedSquares = function(nums){
    let res = new Array(nums.length).fill(0);
    let left = 0;
    let right = nums.length - 1;

    for (let i = nums.length - 1; i >= 0; i--) {
        if (Math.abs(nums[left]) > Math.abs(nums[right])) {
            res[i] = nums[left] ** 2;
            left++;
        } else {
            res[i] = nums[right] ** 2;
            right--;
        }
    }

    return res;
}

// const sortedSquares = function(nums) {
//     const listPositive = nums.map(item => {
//         if (item < 0) {
//             return -1 * item;
//         } else {
//             return item;
//         }
//     });
//
//     const listSorted = listPositive.sort((a, b) => a - b);
//     return listSorted.map(item => item * item);
// };

const nums = [-4, -1, 0, 3, 10];
console.log(sortedSquares(nums));
