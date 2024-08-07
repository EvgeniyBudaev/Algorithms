/* https://leetcode.com/problems/3sum-closest/description/

Учитывая целочисленный массив nums длины n и целочисленную цель, найдите три целых числа в nums, сумма которых наиболее
близка к цели. Верните сумму трех целых чисел. Вы можете предположить, что каждый вход будет иметь ровно одно решение.

Input: nums = [-1,2,1,-4], target = 1
Output: 2
Пояснение: Ближайшая к цели сумма равна 2. (-1 + 2 + 1 = 2).

Input: nums = [0,0,0], target = 1
Output: 0
Пояснение: Сумма, ближайшая к целевой, равна 0. (0 + 0 + 0 = 0).
*/

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var threeSumClosest = function(nums, target) {
    nums.sort((a, b) => a - b); // [-4,-1,1,2]
    const n = nums.length;
    let closestSum = nums[0] + nums[1] + nums[2];
    for (let i = 0; i < n - 2; i++) {
        let left = i + 1, right = n - 1;
        while (left < right) {
            let sum = nums[i] + nums[left] + nums[right];
            if (sum === target) return sum;
            else if (sum < target) left++;
            else right--;
            if (Math.abs(sum - target) < Math.abs(closestSum - target)) closestSum = sum;
        }
    }
    return closestSum;
};

console.log(threeSumClosest([-1,2,1,-4], 1)); // 2