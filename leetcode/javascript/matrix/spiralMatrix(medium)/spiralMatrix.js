/* https://leetcode.com/problems/spiral-matrix/description/

Учитывая матрицу размера m x n, верните все элементы матрицы в спиральном порядке.

Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]
*/

/**
 * @param {number[][]} matrix
 * @return {number[]}
 */
var spiralOrder = function(matrix) {
    // Посчитать общее количество строк и столбцов
    let rows = matrix.length;
    let cols = matrix[0].length;

    // Настройка указателей для перемещения по матрице
    let row = 0;
    let col = -1;

    // Установите начальное направление на 1 для перемещения слева направо.
    let direction = 1;

    // Создайте массив для хранения элементов в спиральном порядке.
    const result = [];

    // Обход матрицы по спирали
    while (rows > 0 && cols > 0) {
        // Двигаемся горизонтально в одном из двух направлений:
        // 1. Слева направо (если направление == 1)
        // 2. Справа налево (если направление == -1) Увеличиваем указатель col для перемещения по горизонтали
        for (let i = 0; i < cols; i++) {
            col += direction;
            console.log("result: ", result);
            result.push(matrix[row][col]);
        }
        rows--;

        // Двигаемся вертикально в одном из двух направлений:
        // 1. Сверху вниз (если направление == 1)
        // 2. Снизу вверх (если направление == -1) Увеличиваем указатель строки для перемещения по вертикали
        for (let i = 0; i < rows; i++) {
            row += direction;
            result.push(matrix[row][col]);
        }
        cols--;

        // Изменение направления для следующего обхода
        direction *= -1;
    }

    return result;
};

const matrix = [[1,2,3],[4,5,6],[7,8,9]];
console.log(spiralOrder(matrix)); // [1,2,3,6,9,8,7,4,5]