/* https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/description/
solution https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/solutions/2369167/javascript-one-liner-solution/

Дана матрица размера n x n, в которой каждая строка и столбец отсортированы в порядке возрастания, верните
k-й наименьший элемент матрицы.
Обратите внимание, что это k-й наименьший элемент в отсортированном порядке, а не k-й отдельный элемент.
Вам необходимо найти решение со сложностью памяти выше O(n2).

Input: matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
Output: 13
Пояснение: Элементы матрицы — [1,5,9,10,11,12,13,13,15], а восьмое наименьшее число — 13.

Input: matrix = [[-5]], k = 1
Output: -5
*/

/**
 * @param {number[][]} matrix
 * @param {number} k
 * @return {number}
 */
var kthSmallest = function(matrix, k) {
    return matrix.flat().sort((a, b) => a - b)[k - 1];
};

const matrix = [[1,5,9],[10,11,13],[12,13,15]];
const k = 8;
console.log(kthSmallest(matrix, k));