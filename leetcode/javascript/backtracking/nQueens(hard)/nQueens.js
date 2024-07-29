/* https://leetcode.com/problems/n-queens/description/

Загадка с n ферзями — это задача о том, как разместить n ферзей на шахматной доске размера n x n так, чтобы никакие два
ферзя не атаковали друг друга.

Учитывая целое число n, найдите все различные решения головоломки с n ферзями. Вы можете вернуть ответ в любом порядке.

Каждое решение содержит отдельную конфигурацию доски для размещения n ферзей, где «Q» и «.». оба указывают на ферзя и
пустое место соответственно.

Input: n = 4
Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
Пояснение: существует два различных решения головоломки с четырьмя ферзями, как показано выше.

Input: n = 1
Output: [["Q"]]
*/

/**
 * @param {number} n
 * @return {string[][]}
 */
var solveNQueens = function(n) {
  const chessBoard = new Array(n);
  for (let i = 0; i < n; i++) {
    chessBoard[i] = new Array(n).fill(".");
  }

  const result = [];

  const isValidQueen = function(row, col) {
    for (let i = 0; i < row; i++) {
      if(chessBoard[i][col] === "Q") return false;
    }
    for (let i = row - 1, j = col - 1; i >= 0 && j >= 0; i--, j--) {
      if(chessBoard[i][j] === "Q") return false;
    }
    for (let i = row - 1, j = col + 1; i >= 0 && j <= n - 1; i--, j++) {
      if(chessBoard[i][j] === "Q") return false;
    }
    return true;
  }

  const backtrack = function(row) {
    if (row === n) {
      result.push([...chessBoard].map(row => row.join('')));
      return;
    }
    for( var col = 0; col < n; col++) {
      if(isValidQueen(row, col)) {
        chessBoard[row][col] = "Q";
        backtrack(row + 1);
        chessBoard[row][col] = ".";
      }
    }
  }

  backtrack(0);
  return result;
};

console.log(solveNQueens(4)); // [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]