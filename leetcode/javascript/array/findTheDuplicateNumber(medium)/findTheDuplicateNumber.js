/* https://leetcode.com/problems/find-the-duplicate-number/description/

Дан массив целых чисел nums, содержащий n + 1 целых чисел, где каждое целое число находится в диапазоне [1, n]
включительно.
В nums есть только одно повторяющееся число, верните это повторяющееся число.
Вы должны решить проблему, не изменяя числа массива и используя только постоянное дополнительное пространство.

Input: nums = [1,3,4,2,2]
Output: 2

Input: nums = [3,1,3,4,2]
Output: 3
*/

/**
 * @param {number[]} nums
 * @return {number}
 */
var findDuplicate = function(nums) {
    let visited = new Array(nums.length).fill(false);
    for (let num of nums) {
        if (!visited[num]) visited[num] = true;
        else return num;
    }
};

var findDuplicate2 = function(nums) {
    const set = new Set();
    for (const num of nums) {
        if (set.has(num)) return num;
        else set.add(num);
    }
};

console.log(findDuplicate([1,3,4,2,2])); // 2