/* https://leetcode.com/problems/convert-1d-array-into-2d-array/description/

Вам дан исходный одномерный (1D) целочисленный массив с нулевым индексом и два целых числа: m и n. Вам предстоит создать
двумерный (2D) массив с m строками и n столбцами, используя все элементы оригинала.

Элементы с индексами от 0 до n - 1 (включительно) исходного должны образовывать первую строку построенного 2D-массива,
элементы с индексами от n до 2 * n - 1 (включительно) должны образовывать вторую строку построенного 2D-массива,
и так далее.

Верните двумерный массив m x n, созданный в соответствии с описанной выше процедурой, или пустой двумерный массив,
если это невозможно.

Input: original = [1,2,3,4], m = 2, n = 2
Output: [[1,2],[3,4]]
Пояснение: Построенный двумерный массив должен содержать 2 строки и 2 столбца.
Первая группа из n=2 элементов в оригинале, [1,2], становится первой строкой в построенном двумерном массиве.
Вторая группа из n=2 элементов в оригинале, [3,4], становится второй строкой в построенном двумерном массиве.
*/

/**
 * @param {number[]} original
 * @param {number} m
 * @param {number} n
 * @return {number[][]}
 */

var construct2DArray = function(original, m, n) {
    if (m * n !== original.length) return [];
    const arr = [];
    for (let i = 0; i < original.length; i += n) {
        arr.push(original.slice(i, i + n));
    }
    return arr;
};

const construct2DArray2 = (original, m, n) => {
    if (m * n !== original.length) return [];
    const result = [];
    let idx = 0;
    while (idx < original.length) {
        result.push(original.slice(idx, idx + n));
        idx += n;
    }
    return result;
}

var construct2DArray3 = function(original, m, n) {
    if (m * n !== original.length) return [];
    const arr = [];
    for (let i = 0, x = 0; i < m; i++) {
        arr[i] = [];
        for (let j = 0; j < n; j++) {
            arr[i][j] = original[x++];
        }
    }
    return arr;
};

console.log(construct2DArray([1,2,3,4], 2, 2)); // [[1,2],[3,4]]