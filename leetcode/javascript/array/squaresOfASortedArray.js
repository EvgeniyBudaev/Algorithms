/* Squares of a Sorted Array */
/* Квадраты отсортированного массива */
/* 
Учитывая целочисленный массив nums, отсортированный в неубывающем порядке, вернуть массив квадратов каждого числа, отсортированных в неубывающем порядке.

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

const compareNumbers = (a, b) => {
    return a - b;
};

const sortedSquares = function(nums) {
    const listPositive = nums.map(item => {
        if (item < 0) {
            return -1 * item;
        } else {
            return item;
        }
    }); 

    const listSorted = listPositive.sort(compareNumbers); 
    return listSorted.map(item => item * item);
};

const nums = [-4, -1, 0, 3, 10];
console.log(sortedSquares(nums));
