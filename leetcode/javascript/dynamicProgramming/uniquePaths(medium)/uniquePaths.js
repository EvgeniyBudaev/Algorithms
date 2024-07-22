/* https://leetcode.com/problems/unique-paths/description/

На сетке m x n находится робот. Робот изначально расположен в верхнем левом углу (т. е. сетка[0][0]). Робот пытается
переместиться в правый нижний угол (т. е. сетку[m - 1][n - 1]). В любой момент времени робот может двигаться только вниз
или вправо.

Учитывая два целых числа m и n, верните количество возможных уникальных путей, по которым робот может добраться до
правого нижнего угла.

Тестовые случаи генерируются так, что ответ будет меньше или равен 2 * 109.

Input: m = 3, n = 7
Output: 28

Input: m = 3, n = 2
Output: 3
Пояснение: Из верхнего левого угла можно попасть в правый нижний угол тремя способами:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down
*/

/**
 * @param {number} m
 * @param {number} n
 * @return {number}
 */
var uniquePaths = function(m, n) {
    let aboveRow = Array(n).fill(1);

    for (let row = 1; row < m; row++) {
        let currentRow = Array(n).fill(1);
        for (let col = 1; col < n; col++) {
            currentRow[col] = currentRow[col - 1] + aboveRow[col];
        }
        aboveRow = currentRow;
    }

    return aboveRow[n - 1];
};

console.log(uniquePaths(3, 7)); // 28