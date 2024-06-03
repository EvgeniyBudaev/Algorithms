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


const arr = [1,0,2,3,0,4,5,0];

/**
 * @param {number[]} arr
 * @return {void} Do not return anything, modify arr in-place instead.
 */
var duplicateZeros = function(arr) {
    let length = arr.length;

    for (let i = 0; i < length; i++) {
        if(arr[i] === 0) {
            for (let j = length - 1; j > i; j--) {
                arr[j] = arr[j-1];
            }
            i++;
        }
    }
};

duplicateZeros(arr);