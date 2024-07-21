/* https://leetcode.com/problems/maximum-product-subarray/description/

Учитывая целочисленный массив чисел, найдите подмассив который имеет самый большой продукт, и возвращает продукт.
Тестовые случаи генерируются таким образом, чтобы ответ соответствовал 32-битному целому числу.

Input: nums = [2,3,-2,4]
Output: 6
Пояснение: [2,3] имеет наибольшее произведение 6.

Input: nums = [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
*/

/**
 * @param {number[]} nums
 * @return {number}
 */
var maxProduct = function(nums) {
    if (nums.length === 0) return 0;

    let maxSoFar = nums[0];
    let minSoFar = nums[0];
    let result = maxSoFar;

    for (let i = 1; i < nums.length; i++) {
        const curr = nums[i];

        const tempMaxSoFar = Math.max(
            curr,
            maxSoFar * curr,
            minSoFar * curr
        );

        minSoFar = Math.min(
            curr,
            maxSoFar * curr,
            minSoFar * curr
        );

        maxSoFar = tempMaxSoFar;
        result = Math.max(maxSoFar, result);
    }

    return result;
};

console.log(maxProduct([2,3,-2,4])); // 6