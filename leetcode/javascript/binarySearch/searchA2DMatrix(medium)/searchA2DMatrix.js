/* https://leetcode.com/problems/search-a-2d-matrix/description/

Вам дана целочисленная матрица размера m x n со следующими двумя свойствами:
Каждая строка сортируется в порядке неубывания.
Первое целое число каждой строки больше последнего целого числа предыдущей строки.
Учитывая целочисленную цель, верните true, если цель находится в матрице, или false в противном случае.
Вы должны написать решение с временной сложностью O(log(m * n)).

Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true

Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false
*/

/**
 * @param {number[][]} matrix
 * @param {number} target
 * @return {boolean}
 */
var searchMatrix = function(matrix, target) {
    const rows = matrix.length, cols = matrix[0].length;
    let left = 0, right = rows * cols - 1;

    while (left <= right) {
        const mid = Math.floor((left + right) / 2); // 5 -> 2 -> 0 -> 1
        // Преобразование индекса в координаты: Значение mid затем преобразуется обратно в координаты строки и
        // столбца (row и col) в исходной матрице. Это делается путем деления mid на количество столбцов (cols) и
        // взятия остатка от деления соответственно.
        const [row, col] = [Math.floor(mid / cols), mid % cols];
        const guess = matrix[row][col]; // 11 -> 5 -> 1 -> 3

        if (guess === target) return true;
        else if (guess < target) left = mid + 1;
        else right = mid - 1;
    }

    return false;
};

const matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]];
console.log(searchMatrix(matrix, 3)); // true