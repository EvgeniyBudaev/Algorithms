// https://leetcode.com/problems/duplicate-zeros/description/

/*
Учитывая целочисленный массив фиксированной длины arr, дублируйте каждое вхождение нуля, сдвигая оставшиеся элементы
вправо.
Обратите внимание, что элементы, превышающие длину исходного массива, не записываются. Внесите вышеуказанные изменения
во входной массив и ничего не возвращайте.
*/

/*
Input: arr = [1,0,2,3,0,4,5,0]
Output: [1,0,0,2,3,0,0,4]
Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]
*/

/*
Input: arr = [1,2,3]
Output: [1,2,3]
Explanation: After calling your function, the input array is modified to: [1,2,3]
*/

/**
 * @param {number[]} arr
 * @return {void} Do not return anything, modify arr in-place instead.
 */
var duplicateZeros = function(arr) {
    for (let i = 0; i < arr.length; i++) {
        if (arr[i] === 0) {
            arr.splice(i, 0, 0);
            arr.pop();
            i++;
        }
    }
};

/**
 * @param {number[]} arr
 * @return {void} Do not return anything, modify arr in-place instead.
 */
// Two pointers
// O(n) time | O(1) space
// var duplicateZeros(easy) = function(arr) {
//     for (let i = 0; i < arr.length - 1; i++) {
//         if (arr[i] == 0) {
//             shiftArray(arr, i + 1, arr.length - 1);
//             arr[++i] = 0;
//         }
//     }
// };

// function shiftArray(array, left, right) {
//     while (left < right) {
//         array[right] = array[right - 1];
//         right--;
//     }
// }

const arr = [1,0,2,3,0,4,5,0];
duplicateZeros(arr);