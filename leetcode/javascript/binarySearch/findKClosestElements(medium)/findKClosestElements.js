/* https://leetcode.com/problems/find-k-closest-elements/description/

Учитывая отсортированный массив целых чисел arr, два целых числа k и x, верните k ближайших к x целых чисел в массиве.
Результат также следует отсортировать по возрастанию.

Целое число a ближе к x, чем целое число b, если:
|a - x| < |b - x|, or
|a - x| == |b - x| and a < b

Input: arr = [1,2,3,4,5], k = 4, x = 3
Output: [1,2,3,4]

Input: arr = [1,2,3,4,5], k = 4, x = -1
Output: [1,2,3,4]
*/

/**
 * @param {number[]} arr
 * @param {number} k
 * @param {number} x
 * @return {number[]}
 */
var findClosestElements = function(arr, k, x) {
    let left = 0, right = arr.length - k;
    while (left < right) {
        const mid = Math.floor((left + right) / 2);
        if (x - arr[mid] > arr[mid + k] - x) left = mid + 1;
        else right = mid;
    }
    return arr.slice(left, left + k);
};

const arr = [1,2,3,4,5];
const k = 4;
const x = 3;
console.log(findClosestElements(arr, k, x)); // [1,2,3,4]