/* https://leetcode.com/problems/peak-index-in-a-mountain-array/description/

Вам дан целочисленный горный массив длиной n, значения которого увеличиваются до пикового элемента, а затем уменьшаются.
Верните индекс пикового элемента.
Ваша задача — решить ее за время сложности O(log(n)).

Input: arr = [0,1,0]
Output: 1

Input: arr = [0,2,1,0]
Output: 1

Input: arr = [0,10,5,2]
Output: 1
*/

/**
 * @param {number[]} arr
 * @return {number}
 */
var peakIndexInMountainArray = function(arr) {
    let left = 0, right = arr.length - 1;

    while (left <= right) {
        const mid = Math.floor((left + right) / 2);
        if (arr[mid] < arr[mid + 1]) left = mid + 1;
        else right = mid - 1;
    }

    return left;
};

console.log(peakIndexInMountainArray([0,1,0])); // 1