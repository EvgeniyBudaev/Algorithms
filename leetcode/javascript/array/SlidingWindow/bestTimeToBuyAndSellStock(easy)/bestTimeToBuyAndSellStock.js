/* https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
solution https://leetcode.com/problems/best-time-to-buy-and-sell-stock/solutions/1735550/python-javascript-easy-solution-with-very-clear-explanation/

Вам дан массив цен, где цены[i] — это цена данной акции на i-й день.
Вы хотите максимизировать свою прибыль, выбрав один день для покупки одной акции и выбрав другой день в будущем для
продажи этой акции.
Верните максимальную прибыль, которую вы можете получить от этой сделки. Если вы не можете получить прибыль, верните 0.

Input: prices = [7,1,5,3,6,4]
Output: 5
Пояснение: Покупайте в день 2 (цена = 1) и продавайте в день 5 (цена = 6), прибыль = 6-1 = 5.
Обратите внимание, что покупка во второй день и продажа в первый день не разрешены, поскольку вы должны купить, прежде
чем продавать.

Input: prices = [7,6,4,3,1]
Output: 0
Объяснение: В этом случае транзакции не выполняются, а максимальная прибыль = 0.
*/

/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function(prices) {
    let [left, right, profit] = [0, 1, 0]; // left - для покупки, right - для продажи
    while (right < prices.length) {
        const canSlide = prices[right] <= prices[left];
        if (canSlide) left = right;
        const window = prices[right] - prices[left];
        profit = Math.max(profit, window);
        right++;
    }
    return profit;
};

var maxProfit2 = function(prices) {
    let buyPrice = prices[0], profit = 0;
    for (let i = 1; i < prices.length; i++) {
        if (prices[i] < buyPrice) buyPrice = prices[i];
        profit = Math.max(profit, prices[i] - buyPrice);
    }
    return profit;
};

console.log(maxProfit([7,1,5,3,6,4])); // 5