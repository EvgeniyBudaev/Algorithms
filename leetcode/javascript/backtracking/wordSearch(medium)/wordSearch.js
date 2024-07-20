/* https://leetcode.com/problems/word-search/description/

Учитывая сетку символов m x n и строковое слово, верните true, если слово существует в сетке.
Слово может быть составлено из букв последовательно соседних ячеек, где соседние ячейки являются соседними по
горизонтали или по вертикали. Одна и та же буквенная ячейка не может использоваться более одного раза.

Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true
*/

/**
 * @param {string[][]} board
 * @param {string} word
 * @return {boolean}
 */
var exist = function(board, word) {
    const m = board.length;
    const n = board[0].length;

    const backtrack = (i, j, k) => {
        if (k === word.length) return true;
        if (i < 0 || i >= m || j < 0 || j >= n || board[i][j] !== word.charAt(k)) return false;

        const temp = board[i][j];
        board[i][j] = '\0';

        const result = backtrack(i + 1, j, k + 1) ||
            backtrack(i - 1, j, k + 1) ||
            backtrack(i, j + 1, k + 1) ||
            backtrack(i, j - 1, k + 1);

        board[i][j] = temp;
        return result;
    };

    for (let i = 0; i < m; ++i) {
        for (let j = 0; j < n; ++j) {
            if (backtrack(i, j, 0)) return true;
        }
    }
    return false;
};

const board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]];
const word = "ABCCED";
console.log(exist(board, word));