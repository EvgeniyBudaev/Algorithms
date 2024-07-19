/* https://leetcode.com/problems/set-matrix-zeroes/description/

Учитывая целочисленную матрицу размера m x n, если элемент равен 0, установите всю ее строку и столбец в 0.
Вы должны сделать это на месте.

Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]
*/

/**
 * @param {number[][]} matrix
 * @return {number[][]} Do not return anything, modify matrix in-place instead.
 */
var setZeroes = function(matrix) {
    let row = new Array(matrix.length).fill(0);
    let col = new Array(matrix[0].length).fill(0);
    for (let i = 0;i < matrix.length; i++) {
        for (let j = 0; j < matrix[i].length; j++) {
            if( matrix[i][j] === 0) {
                row[i] = 1;
                col[j] = 1;
            }
        }
    }

    for (let i = 0;i < matrix.length; i++) {
        for (let j = 0;j < matrix[i].length; j++) {
            if (row[i]||col[j]) {
                matrix[i][j] = 0;
            }
        }
    }
    return matrix;
};

const matrix = [[1,1,1],[1,0,1],[1,1,1]];
console.log(setZeroes(matrix));