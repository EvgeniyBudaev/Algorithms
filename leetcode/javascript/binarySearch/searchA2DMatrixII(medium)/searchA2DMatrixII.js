/* https://leetcode.com/problems/search-a-2d-matrix-ii/description/

Напишите эффективный алгоритм, который ищет целевое значение в целочисленной матрице размером m x n. Эта матрица имеет
следующие свойства:
Целые числа в каждой строке сортируются по возрастанию слева направо.
Целые числа в каждом столбце сортируются по возрастанию сверху вниз.

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
   const n = matrix.length;
    for (let i = 0; i < n; i++) {
        const row = matrix.pop();
        if (row.includes(target)) return true;
    }
    return false;
};

const matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]];
const target = 5;
console.log(searchMatrix(matrix, target)); // true