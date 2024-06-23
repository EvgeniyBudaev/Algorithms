// https://leetcode.com/problems/coin-change/description/

/*
Вам дан целочисленный массив coin, представляющий монеты разного номинала, и целочисленная сумма, представляющая общую
 сумму денег.
Верните наименьшее количество монет, которое вам нужно, чтобы составить эту сумму. Если эта сумма денег не может быть
 составлена ни одной комбинацией монет, верните -1.
Вы можете предположить, что у вас есть бесконечное количество монет каждого вида.
*/

/*
Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
*/

/*
Input: coins = [2], amount = 3
Output: -1
*/

/*
Input: coins = [1], amount = 0
Output: 0
*/

/**
 * @param {number[]} coins
 * @param {number} amount
 * @return {number}
 */
var coinChange = function(coins, amount) {
    if (amount < 1) {
        return 0;
    }

    let dp = new Array(amount + 1).fill(Infinity);

    dp[0] = 0;

    for (let coin of coins) {

        for(let i = coin; i <= amount; i++) {
            dp[i] = Math.min(dp[i], dp[i - coin] + 1);
        }
    }

    return dp[amount] === Infinity ? -1: dp[amount];
};

console.log(coinChange([1,2,5], 11)); // 3