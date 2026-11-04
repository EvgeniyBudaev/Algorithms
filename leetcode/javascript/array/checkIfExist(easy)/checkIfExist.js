/* 1346. Check If N and Its Double Exist
https://leetcode.com/problems/check-if-n-and-its-double-exist/description/

Учитывая массив целых чисел, проверьте, существуют ли два индекса i и j такие, что:
i != j
0 <= i, j < arr.length
arr[i] == 2 * arr[j]

Input: arr = [10,2,5,3]
Output: true
Explanation: For i = 0 and j = 2, arr[i] == 10 == 2 * 5 == 2 * arr[j]

Input: arr = [3,1,7,11]
Output: false
Пояснение: Не существует i и j, удовлетворяющих этим условиям.
*/

/**
 * @param {number[]} arr
 * @return {boolean}
 * @time O(n)
 * @space O(n)
 */
var checkIfExist = function(arr) {
    const seen = new Set();

    for (let num of arr) {
        if (seen.has(num * 2) || (num % 2 === 0 && seen.has(num / 2))) {
            return true;
        }
        seen.add(num);
    }

    return false;
};

console.log(checkIfExist([10,2,5,3])); // true
console.log(checkIfExist([5, 10])); // true