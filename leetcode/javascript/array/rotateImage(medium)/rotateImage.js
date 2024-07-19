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

    for (let row = 0; row < n / 2; row++) {
        for (let col = row; col < n - row - 1; col++) {
            // Поменяйте местами верхнюю левую и верхнюю правую ячейки в текущей группе.
            [matrix[row][col], matrix[col][n - 1 - row]] = [
                matrix[col][n - 1 - row],
                matrix[row][col],
            ];

            // Поменяйте местами верхнюю левую и нижнюю правую ячейки в текущей группе.
            [matrix[row][col], matrix[n - 1 - row][n - 1 - col]] =
                [
                    matrix[n - 1 - row][n - 1 - col],
                    matrix[row][col],
                ];

            // Поменяйте местами верхнюю левую и нижнюю левую ячейки в текущей группе.
            [matrix[row][col], matrix[n - 1 - col][row]] = [
                matrix[n - 1 - col][row],
                matrix[row][col],
            ];
        }
    }
    return matrix;
};

const matrix = [[1,2,3],[4,5,6],[7,8,9]];
console.log(rotate(matrix)); // [[7,4,1],[8,5,2],[9,6,3]]