/* https://leetcode.com/problems/pacific-atlantic-water-flow/description/
solution https://leetcode.com/problems/pacific-atlantic-water-flow/solutions/1018375/explanation-dfs/

Существует прямоугольный остров размером m x n, который граничит как с Тихим, так и с Атлантическим океаном. Тихий океан
касается левого и верхнего края острова, а Атлантический океан касается правого и нижнего края острова.

Остров разделен на сетку квадратных ячеек. Вам дана целочисленная матрица высот размером m x n, где heights[r][c]
представляет высоту над уровнем моря ячейки в координате (r, c).

На острове выпадает много дождя, и дождевая вода может течь в соседние ячейки прямо на север, юг, восток и запад, если
высота соседней ячейки меньше или равна высоте текущей ячейки. Вода может перетекать из любой клетки, прилегающей к океану, в океан.

Возвращает двумерный список координат сетки, где result[i] = [ri, ci] означает, что дождевая вода может течь из ячейки
(ri, ci) как в Тихий, так и в Атлантический океаны.
Input: heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
Output: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
Пояснение: Следующие клетки могут течь в Тихий и Атлантический океаны, как показано ниже:
[0,4]: [0,4] -> Pacific Ocean
       [0,4] -> Atlantic Ocean
[1,3]: [1,3] -> [0,3] -> Pacific Ocean
       [1,3] -> [1,4] -> Atlantic Ocean
[1,4]: [1,4] -> [1,3] -> [0,3] -> Pacific Ocean
       [1,4] -> Atlantic Ocean
[2,2]: [2,2] -> [1,2] -> [0,2] -> Pacific Ocean
       [2,2] -> [2,3] -> [2,4] -> Atlantic Ocean
[3,0]: [3,0] -> Pacific Ocean
       [3,0] -> [4,0] -> Atlantic Ocean
[3,1]: [3,1] -> [3,0] -> Pacific Ocean
       [3,1] -> [4,1] -> Atlantic Ocean
[4,0]: [4,0] -> Pacific Ocean
       [4,0] -> Atlantic Ocean
Обратите внимание, что существуют и другие возможные пути попадания этих клеток в Тихий и Атлантический океаны.
*/

/**
 * @param {number[][]} heights
 * @return {number[][]}
 */
var pacificAtlantic = function(heights) {
    if (heights.length === 0) return [];
    let numRows = heights.length;
    let numCols = heights[0].length;

    let atlantic = [];
    let pacific = [];
    for (let i = 0; i < numRows; i++) {
        atlantic.push(new Array(numCols).fill(false));
        pacific.push(new Array(numCols).fill(false));
    }

    for (let col = 0; col < heights[0].length; col++) {
        dfs(heights, 0, col, Number.MIN_SAFE_INTEGER, pacific);
        dfs(heights, numRows - 1, col, Number.MIN_SAFE_INTEGER, atlantic);
    }

    for (let row = 0;row < heights.length; row++) {
        dfs(heights, row, 0, Number.MIN_SAFE_INTEGER, pacific);
        dfs(heights, row, numCols - 1, Number.MIN_SAFE_INTEGER, atlantic);
    }

    let res = [];
    for (let i = 0; i < numRows; i++) {
        for (let j = 0; j < numCols; j++){
            if (atlantic[i][j] && pacific[i][j]) res.push([i, j]);
        }
    }
    return res;
}



const dfs = (matrix, i, j, prev, ocean) =>{
    //checkOutOfBounds
    if (i<0 ||
        i > matrix.length -1 ||
        j < 0 ||
        j > matrix[i].length - 1
    ) return;


    if (matrix[i][j] < prev) return;
    if (ocean[i][j]) return;
    ocean[i][j] = true;

    dfs(matrix, i+1, j, matrix[i][j], ocean);
    dfs(matrix, i-1, j, matrix[i][j], ocean);
    dfs(matrix, i, j+1, matrix[i][j], ocean);
    dfs(matrix, i, j-1, matrix[i][j], ocean);
}

const heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]];
console.log(pacificAtlantic(heights)); // [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]