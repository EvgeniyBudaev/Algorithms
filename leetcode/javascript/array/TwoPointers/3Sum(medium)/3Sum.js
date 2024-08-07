/* https://leetcode.com/problems/3sum/description/

Учитывая целочисленный массив nums, верните все тройки [nums[i], nums[j], nums[k]] такие,
что i != j, i != k и j != k и nums[i] + nums[j] + nums[к] == 0.
Обратите внимание, что набор решений не должен содержать повторяющихся троек.

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
Отдельные тройки — это [-1,0,1] и [-1,-1,2].
Обратите внимание, что порядок вывода и порядок троек не имеют значения.

Input: nums = [0,1,1]
Output: []
Пояснение: Единственная возможная тройка не дает в сумме 0.

Input: nums = [0,0,0]
Output: [[0,0,0]]
Пояснение: Сумма единственно возможной тройки равна 0.
*/

/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var threeSum = function(nums) {
    const res = [];
    nums.sort((a, b) => a - b); // [-4,-1,-1,0,1,2]

    for (let i = 0; i < nums.length; i++) {
        const num = nums[i];
        if (num > 0) break;
        if (i > 0 && num === nums[i - 1]) continue;

        let left = i + 1;
        let right = nums.length - 1;
        while (left < right) {
            const threeSum = num + nums[left] + nums[right];
            if (threeSum > 0) {
                right--;
            } else if (threeSum < 0) {
                left++;
            } else {
                res.push([num, nums[left], nums[right]]);
                left++;
                right--;
                while (nums[left] === nums[left - 1] && left < right) {
                    left++;
                }
            }
        }
    }

    return res;
};

console.log(threeSum([-1,0,1,2,-1,-4]));