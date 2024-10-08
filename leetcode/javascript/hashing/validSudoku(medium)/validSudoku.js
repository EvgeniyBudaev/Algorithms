/* https://leetcode.com/problems/valid-sudoku/description/

Определите, действительна ли доска для судоку 9 x 9. Только заполненные ячейки подлежат проверке согласно следующим
правилам:

Каждая строка должна содержать цифры от 1 до 9 без повторений.
Каждый столбец должен содержать цифры 1–9 без повторений.
Каждый из девяти подполей сетки размером 3х3 должен содержать цифры от 1 до 9 без повторений.

Примечание:
Доска судоку (частично заполненная) может быть допустимой, но не обязательно решаемой.
Только заполненные ячейки должны быть проверены в соответствии с указанными правилами.

Input: board =
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: true

Input: board =
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: false
*/

/**
 * @param {string[][]} board
 * @return {boolean}
 */
var isValidSudoku = function (board) {
  let boardLength = board.length;
  let rows = Array.from({length: boardLength}, () => new Set());
  let cols = Array.from({length: boardLength}, () => new Set());
  let boxes = Array.from({length: boardLength}, () => new Set());

  for (let r = 0; r < boardLength; r++) {
      for (let c = 0; c < boardLength; c++) {
          let val = board[r][c];
          if (val === ".") continue;
          let boxIndex = Math.floor(r / 3) * 3 + Math.floor(c / 3);
          if (rows[r].has(val) || cols[c].has(val) || boxes[boxIndex].has(val)) return false;
          rows[r].add(val);
          cols[c].add(val);
          boxes[boxIndex].add(val);
      }
  }
  return true;
};

const board = [
  ["5", "3", ".", ".", "7", ".", ".", ".", "."],
  ["6", ".", ".", "1", "9", "5", ".", ".", "."],
  [".", "9", "8", ".", ".", ".", ".", "6", "."],
  ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
  ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
  ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
  [".", "6", ".", ".", ".", ".", "2", "8", "."],
  [".", ".", ".", "4", "1", "9", ".", ".", "5"],
  [".", ".", ".", ".", "8", ".", ".", "7", "9"],
];
console.log(isValidSudoku(board)); // true
