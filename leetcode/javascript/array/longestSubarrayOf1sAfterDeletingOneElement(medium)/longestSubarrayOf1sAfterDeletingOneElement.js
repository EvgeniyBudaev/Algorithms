/* https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/description/

Учитывая двоичный массив nums, вам следует удалить из него один элемент.
Возвращает размер самого длинного непустого подмассива, содержащего только 1 в результирующем массиве. Верните 0, если
такого подмассива нет.

Input: nums = [1,1,0,1]
Output: 3
Объяснение: После удаления числа в позиции 2 [1,1,1] содержит 3 числа со значением 1.

Input: nums = [0,1,1,1,0,1,1,0,1]
Output: 5
Пояснение: После удаления числа в позиции 4, [0,1,1,1,1,1,0,1] самым длинным подмассивом со значением единиц будет [1,1,1,1,1].

Input: nums = [1,1,1]
Output: 2
Объяснение: Вам необходимо удалить один элемент.
*/

/**
 * @param {number[]} nums
 * @return {number}
 */
var longestSubarray = function(nums) {
    let leftPtr = 0;
    let result = 0;
    let zeroCount = 0;

    for (let rightPtr = 0; rightPtr < nums.length; rightPtr++) {
        if (nums[rightPtr] === 0) zeroCount++;
        if (zeroCount > 1) {
            if (nums[leftPtr] === 0) zeroCount--;
            leftPtr++;
        }
        if (zeroCount <= 1) result = Math.max(result, rightPtr - leftPtr);
    }

    return result;
};

console.log(longestSubarray([1,1,0,1])); // 3