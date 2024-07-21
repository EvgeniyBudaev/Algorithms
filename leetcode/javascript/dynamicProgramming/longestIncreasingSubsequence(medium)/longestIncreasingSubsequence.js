/* https://leetcode.com/problems/longest-increasing-subsequence/description/

Учитывая целочисленный массив nums, верните длину самого длинного, строго увеличивающегося последовательность.

Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Пояснение: Самая длинная возрастающая подпоследовательность — [2,3,7,101], поэтому длина равна 4.

Input: nums = [0,1,0,3,2,3]
Output: 4

Input: nums = [7,7,7,7,7,7,7]
Output: 1
*/

/**
 * @param {number[]} nums
 * @return {number}
 */
var lengthOfLIS = function(nums) {
    let result = Array.from(nums).fill(1);

    for (let i = 0; i < nums.length; i++) {
        for (let j = i; j >= 0; j--) {
            if (nums[i] > nums[j]) {
                result[i] = Math.max(result[i], result[j] + 1);
            }
        }
    }

    return Math.max(...result);
};

const nums = [10,9,2,5,3,7,101,18];
console.log(lengthOfLIS(nums)); // 4