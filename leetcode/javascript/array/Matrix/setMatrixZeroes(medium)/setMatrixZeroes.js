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
    const rows = matrix.length;
    const cols = matrix[0].length;
    const zeroRow = new Set();
    const zeroCol = new Set();

    for (let row = 0; row < rows; row++) {
        for (let col = 0; col < cols; col++) {
            if (matrix[row][col] === 0) {
                zeroRow.add(row);
                zeroCol.add(col);
            }
        }
    }

    zeroRow.forEach((row) => {
        for (let col = 0; col < cols; col++) {
            matrix[row][col] = 0;
        }
    })

    zeroCol.forEach((col) => {
        for (let row = 0; row < rows; row++) {
            matrix[row][col] = 0;
        }
    })
};

const matrix = [[1,1,1],[1,0,1],[1,1,1]];
console.log(setZeroes(matrix));