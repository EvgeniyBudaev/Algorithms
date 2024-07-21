/* https://leetcode.com/problems/target-sum/description/

Вам дан целочисленный массив nums и целочисленная цель.
Вы хотите построить выражение из чисел, добавив один из символов «+» и «-» перед каждым целым числом в числах, а затем
объединить все целые числа.
Например, если nums = [2, 1], вы можете добавить «+» перед 2 и «-» перед 1 и объединить их, чтобы построить выражение
«+2-1».
Верните количество различных выражений, которые вы можете создать и которые достигают целевого значения.

Input: nums = [1,1,1,1,1], target = 3
Output: 5
Пояснение: Есть 5 способов назначить символы, чтобы сумма чисел соответствовала целевому значению 3.
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3

Input: nums = [1], target = 1
Output: 1
*/

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var findTargetSumWays = function(nums, target) {
    // Memo
    this.memo = new Map();
    return dp(nums, target, 0, 0);
};

const dp = function(nums, target, index, sum) {
    let key = `${index}_${sum}`;

    // Базовый случай, дойти до конца `nums`
    if (index === nums.length) {
        // Сумма - это то, что мы ищем
        if (sum === target) return 1;

        // Сумма не та, которую мы ищем
        return 0;
    }

    // Return from memo
    if (this.memo.has(key) === true) {
        return this.memo.get(key);
    }

    // К каждому числу в `nums` мы можем добавить "+" или "-"
    // Чтобы добавить знак «+», мы просто оставляем все как есть, поскольку все числа в `nums` положительны, кроме 0
    // Чтобы добавить знак «-», мы вычитаем его, поскольку «-» приводит к вычитанию
    let positive = dp(nums, target, index + 1, sum + nums[index]);
    let negative = dp(nums, target, index + 1, sum - nums[index]);

    // Рассчитайте общую сумму
    let total = positive + negative;

    // Set memo
    this.memo.set(key, total);

    return total;
};

const nums = [1,1,1,1,1];
const target = 3;
console.log(findTargetSumWays(nums, target)); // 5