/* https://leetcode.com/problems/rotate-image/description/

Вам дана двумерная матрица размера n x n, представляющая изображение. Поверните изображение на 90 градусов
(по часовой стрелке).
Вам необходимо повернуть изображение на месте, а это значит, что вам придется напрямую изменить входную 2D-матрицу.
НЕ выделяйте еще одну 2D-матрицу и не выполняйте вращение.

Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]
*/

/**
 * @param {number[][]} matrix
 * @return {number[][]} Do not return anything, modify matrix in-place instead.
 */
var rotate = function(matrix) {
    const n = matrix.length;
    for (let i = 0; i < n; i++) {
        for (let j = i; j < n; j++) {
            let temp = matrix[i][j];
            matrix[i][j] = matrix[j][i];
            matrix[j][i] = temp;
        }
    }

    // Переверните каждую строку
    for (let i = 0; i < n; i++) {
        matrix[i].reverse();
    }

    return matrix;
};

const matrix = [[1,2,3],[4,5,6],[7,8,9]];
console.log(rotate(matrix)); // [[7,4,1],[8,5,2],[9,6,3]]