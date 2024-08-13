/* https://leetcode.com/problems/top-k-frequent-elements/description/

Учитывая целочисленный массив nums и целое число k, верните k наиболее часто встречающихся элементов. Вы можете вернуть
ответ в любом порядке.

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

Input: nums = [1], k = 1
Output: [1]
*/

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number[]}
 */
var topKFrequent = function(nums, k) {
    const map = new Map();

    for (let num of nums) {
        map.set(num, (map.get(num) ?? 0) + 1);
    }

    // sortedByValue: [[1,3], [2,2], [3,1]]
    const sortedByValue = [...map.entries()].sort((a, b) => b[1] - a[1]);
    return sortedByValue.slice(0, k).map(entry => entry[0]);
};

console.log(topKFrequent([1,1,1,2,2,3], 2)); // [1,2]