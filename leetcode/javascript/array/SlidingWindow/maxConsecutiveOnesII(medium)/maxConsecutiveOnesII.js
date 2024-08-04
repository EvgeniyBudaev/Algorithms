/* https://github.com/EvgeniyBudaev/doocs_leetcode/blob/main/solution/0400-0499/0487.Max%20Consecutive%20Ones%20II/README.md

Учитывая двоичный массив nums, если не более одного можно перевернуть 0, верните 1 максимальное количество
последовательных чисел в массиве.

Input: nums = [1,0,1,1,0]
Output: 4
Объяснение: переверните первый 0, чтобы получить самую длинную последовательную 1.
     После переворачивания максимальное количество последовательных единиц равно 4.

Input: nums = [1,0,1,1,0,1]
Output: 4
*/

/**
 * @param {number[]} nums
 * @return {number}
 */
var findMaxConsecutiveOnes = function (nums) {
    let left = 0, zeroCount = 0, maxCountOnes = 0, maxFlipOperations = 1;
    for (let right = 0; right < nums.length; right++) {
        if (nums[right] === 0) zeroCount++;
        while (zeroCount > maxFlipOperations) {
            if (nums[left] === 0) zeroCount--;
            left++;
        }
        maxCountOnes = Math.max(maxCountOnes, right - left + 1);
    }
    return maxCountOnes;
};

console.log(findMaxConsecutiveOnes([1,0,1,1,0])); // 4