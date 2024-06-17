/*
Вам дана доска с целочисленной матрицей размера n x n, где ячейки помечены от 1 до n2 в стиле бустрофедона, начиная с
нижнего левого угла доски (т.е. доска[n - 1][0]) и меняя направление в каждой строке.

Вы начинаете с клетки 1 доски. В каждом ходе, начиная с квадрата, делайте следующее:

Далее выберите целевой квадрат с меткой в ​​диапазоне [curr + 1, min(curr + 6, n2)].
Этот выбор имитирует результат стандартного 6-гранного броска кубика: т. е. всегда существует не более 6 пунктов
назначения, независимо от размера игрового поля.
Если рядом есть змея или лестница, вы должны перейти к месту назначения этой змеи или лестницы. В противном случае вы
переходите к следующему.
Игра заканчивается, когда вы достигаете квадрата n2.
На квадрате доски в строке r и столбце c есть змея или лестница, если board[r][c] != -1. Пункт назначения этой змеи или
лестницы — доска[r][c]. В квадратах 1 и n2 нет ни змеи, ни лестницы.

Обратите внимание, что вы берете змею или лестницу не более одного раза за ход. Если пунктом назначения змеи или
лестницы является начало другой змеи или лестницы, вы не следуете за следующей змейкой или лестницей.

Например, предположим, что доска — это [[-1,4],[-1,3]], и при первом ходе ваша клетка назначения — 2. Вы следуете по
лестнице до клетки 3, но не следуете по следующей лестнице. до 4.
Возвращает наименьшее количество ходов, необходимое для достижения квадрата n2. Если дойти до квадрата невозможно,
верните -1.
*/

/*
Input: board = [[-1,-1,-1,-1,-1,-1],[-1,-1,-1,-1,-1,-1],[-1,-1,-1,-1,-1,-1],[-1,35,-1,-1,13,-1],[-1,-1,-1,-1,-1,-1],[-1,15,-1,-1,-1,-1]]
Output: 4
Explanation:
In the beginning, you start at square 1 (at row 5, column 0).
You decide to move to square 2 and must take the ladder to square 15.
You then decide to move to square 17 and must take the snake to square 13.
You then decide to move to square 14 and must take the ladder to square 35.
You then decide to move to square 36, ending the game.
This is the lowest possible number of moves to reach the last square, so return 4.
 */

/**
 * @param {number[][]} board
 * @return {number}
 */
var snakesAndLadders = function(board) {
    let n = board.length;
    let set = new Set();
    let getPos = (pos) => {
        let row = Math.floor((pos - 1) / n);
        let col = (pos - 1) % n;
        col = row % 2 === 1 ? n - 1 - col : col;
        row = n - 1 - row;
        return [row, col];
    }

    let q = [[1, 0]];

    while(q.length > 0) {
        const [pos, moves] = q.shift();
        for (let i  = 1; i < 7; i++) {
            let newPos = i + pos;
            let [r, c] = getPos(newPos);
            if (board[r][c] !== -1 ) newPos = board[r][c];
            if (newPos === n * n) return moves + 1;
            if (!set.has(newPos)) {
                set.add(newPos);
                q.push([newPos,moves + 1]);
            }
        }
    }

    return -1;
};

const board = [
    [-1,-1,-1,-1,-1,-1],
    [-1,-1,-1,-1,-1,-1],
    [-1,-1,-1,-1,-1,-1],
    [-1,35,-1,-1,13,-1],
    [-1,-1,-1,-1,-1,-1],
    [-1,15,-1,-1,-1,-1]
];
console.log(snakesAndLadders(board));