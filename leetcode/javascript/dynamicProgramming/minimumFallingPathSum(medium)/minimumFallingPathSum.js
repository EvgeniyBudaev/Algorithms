// https://leetcode.com/problems/unique-paths-ii/description/

/*
Учитывая матрицу целых чисел размера n x n, верните минимальную сумму любого падающего пути через матрицу.
Падающий путь начинается с любого элемента в первой строке и выбирает элемент в следующей строке, который находится либо
 непосредственно под ним, либо по диагонали влево/вправо. В частности, следующим элементом с позиции (row, col) будет
 (row + 1, col - 1), (row + 1, col) или (row + 1, col + 1).
*/

/*
Input: matrix = [[2,1,3],[6,5,4],[7,8,9]]
Output: 13
Пояснение: Есть два пути падения с минимальной суммой, как показано.
*/

/**
 * @param {number[][]} matrix
 * @return {number}
 */
var minFallingPathSum = function(matrix) {
    const dp = [matrix[0]];

    for (let i = 1; i < matrix.length; i++) {
        dp.push([]);
        for (let j = 0; j < matrix.length; j++) {
            let prev = Infinity;
            let curr = matrix[i][j] + dp[i - 1][j];
            let next = Infinity;
            if (j > 0) {
                prev = matrix[i][j] + dp[i - 1][j - 1];
            }
            if (j < length - 1) {
                next = matrix[i][j] + dp[i - 1][j + 1];
            }
            dp[i].push(Math.min(prev, curr, next));
        }
    }

    return Math.min(...dp[dp.length - 1]);
};

const matrix = [[2,1,3],[6,5,4],[7,8,9]];
console.log(minFallingPathSum(matrix)); // 13