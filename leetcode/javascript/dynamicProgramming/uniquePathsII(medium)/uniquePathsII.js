// https://leetcode.com/problems/unique-paths-ii/description/

/*
Вам дан массив целых чисел размером m x n. Изначально в верхнем левом углу находится робот (т. е. сетка[0][0]).
 Робот пытается переместиться в правый нижний угол (т. е. сетку[m - 1][n - 1]).
 В любой момент времени робот может двигаться только вниз или вправо.
Препятствие и пространство помечаются в сетке цифрами 1 или 0 соответственно. Путь, по которому движется робот, не может
 включать квадраты, являющиеся препятствиями.
Возвращает количество возможных уникальных путей, по которым робот может добраться до правого нижнего угла.
Тестовые примеры генерируются так, что ответ будет меньше или равен 2 * 109.
*/

/*
Input: obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
Output: 2
Пояснение: В центре сетки 3х3 выше находится одно препятствие.
Есть два способа добраться до правого нижнего угла:
1. Вправо -> Вправо -> Вниз -> Вниз.
2. Вниз -> Вниз -> Вправо -> Вправо.
*/

/**
 * @param {number[][]} obstacleGrid
 * @return {number}
 */
var uniquePathsWithObstacles = function(obstacleGrid) {
    if (!obstacleGrid.length || !obstacleGrid[0].length || obstacleGrid[0][0] === 1) {
        return 0;
    }

    let m = obstacleGrid.length;
    let n = obstacleGrid[0].length;

    let previous = new Array(n).fill(0);
    let current = new Array(n).fill(0);
    previous[0] = 1;

    for (let i = 0; i < m; i++) {
        current[0] = obstacleGrid[i][0] === 1 ? 0 : previous[0];
        for (let j = 1; j < n; j++) {
            current[j] = obstacleGrid[i][j] === 1 ? 0 : current[j-1] + previous[j];
        }
        previous = [...current];
    }

    return previous[n-1];
};

const obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]];
console.log(uniquePathsWithObstacles(obstacleGrid)); // 2