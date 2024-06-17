/*
Вам дан матричный лабиринт размером m x n (с индексом 0) с пустыми ячейками (обозначенными как «.») и стенами
(обозначенными как «+»). Вам также дан вход в лабиринт, где entrance = [entrancerow, entrancecol] обозначает строку и
столбец ячейки, в которой вы изначально стоите.

За один шаг вы можете переместить одну ячейку вверх, вниз, влево или вправо. Вы не можете войти в клетку со стеной и не
можете выйти за пределы лабиринта. Ваша цель — найти ближайший выход из входа. Выход определяется как пустая ячейка,
находящаяся на границе лабиринта. Вход не считается выходом.

Возвращает количество шагов по кратчайшему пути от входа до ближайшего выхода или -1, если такого пути не существует.
*/

/*
Input: maze = [["+","+",".","+"],[".",".",".","+"],["+","+","+","."]], entrance = [1,2]
Output: 1
Explanation: There are 3 exits in this maze at [1,0], [0,2], and [2,3].
Initially, you are at the entrance cell [1,2].
- You can reach [1,0] by moving 2 steps left.
- You can reach [0,2] by moving 1 step up.
It is impossible to reach [2,3] from the entrance.
Thus, the nearest exit is [0,2], which is 1 step away.
 */

/**
 * @param {character[][]} maze
 * @param {number[]} entrance
 * @return {number}
 */
var nearestExit = function(maze, entrance) {
    // bfs
    const [y0, x0] = entrance;
    maze[y0][x0] = '@';
    const queue = [[y0, x0, 0]];

    while (queue.length) {
        const [y, x, step] = queue.shift();
        for (const [dy, dx] of [[-1, 0], [0, -1], [1, 0], [0, 1]]) {
            const ny = y + dy, nx = x + dx;
            if (!maze[ny]?.[nx]) {
                if (step) return step
            } else if (maze[ny][nx] === '.') {
                queue.push([ny, nx, step + 1]);
                maze[ny][nx] = '*';
            }
        }
    }

    return -1;
};

const maze = [["+","+",".","+"],[".",".",".","+"],["+","+","+","."]];
const entrance = [1,2];
console.log(nearestExit(maze, entrance));