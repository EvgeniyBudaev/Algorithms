/*
Input: arr = [10,2,5,3]
Output: true
Explanation: For i = 0 and j = 2, arr[i] == 10 == 2 * 5 == 2 * arr[j]
*/

/*
Input: arr = [3,1,7,11]
Output: false
Explanation: There is no i and j that satisfy the conditions.
*/

/**
 * @param {number[]} arr
 * @return {boolean}
 */
var checkIfExist = function(arr) {
    let left = 0, right = 1;

    while (left < arr.length-1) {
        if (arr[left] === (arr[right] * 2) || arr[right] === (arr[left] * 2)) {
            return true
        } else if (right === arr.length - 1) {
            left++
            right = left + 1
        } else {
            right++
        }
    }
    return false;
};

console.log(checkIfExist([10,2,5,3]));