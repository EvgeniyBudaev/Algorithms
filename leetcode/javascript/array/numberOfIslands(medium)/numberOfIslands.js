/* https://leetcode.com/problems/number-of-islands/description/

Учитывая двумерную двоичную сетку размером m x n, которая представляет собой карту из «1» (суша) и «0» (вода), верните
количество островов.
Остров окружен водой и образован путем соединения соседних земель по горизонтали или вертикали. Вы можете предположить,
что все четыре края сетки окружены водой.

Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
*/

/**
 * @param {string[][]} grid
 * @return {number}
 */
var numIslands = function(grid) {
    if (!grid.length || !grid[0].length) return 0;

    const rows = grid.length;
    const cols = grid[0].length;
    let islands = 0;

    const dfs = (row, col) => {
        if (row < 0 || col < 0 || row >= rows || col >= cols || grid[row][col] !== '1') return;
        // Помечаем все ячейки, принадлежащие одному острову, как посещенные («0»), чтобы они больше не учитывались.
        grid[row][col] = '0';
        // Рекурсивно вызываем все четыре соседние ячейки (вверх, вниз, влево, вправо), чтобы продолжить исследование острова.
        dfs(row - 1, col);
        dfs(row + 1, col);
        dfs(row, col - 1);
        dfs(row, col + 1);
    };

    for (let row = 0; row < rows; row++) {
        for (let col = 0; col < cols; col++) {
            if (grid[row][col] === '1') {
                dfs(row, col);
                islands++;
            }
        }
    }

    return islands;
};

const grid = [
    ["1","1","1","1","0"],
    ["1","1","0","1","0"],
    ["1","1","0","0","0"],
    ["0","0","0","0","0"]
];
console.log(numIslands(grid)); // 1