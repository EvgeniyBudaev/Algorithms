/* https://leetcode.com/problems/search-a-2d-matrix-ii/description/

Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix. This matrix has the
following properties:
Integers in each row are sorted in ascending from left to right.
Integers in each column are sorted in ascending from top to bottom.

Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
Output: true

Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
Output: false
*/

/**
 * @param {number[][]} matrix
 * @param {number} target
 * @return {boolean}
 */
var searchMatrix = function(matrix, target) {
    let n = matrix.length;
    for (let i=0; i < n; i++) {
        let row = matrix.pop();
        if(row.includes(target)) return true;
    }
    return false;
};

const matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]];
const target = 5;
console.log(searchMatrix(matrix, target)); // true