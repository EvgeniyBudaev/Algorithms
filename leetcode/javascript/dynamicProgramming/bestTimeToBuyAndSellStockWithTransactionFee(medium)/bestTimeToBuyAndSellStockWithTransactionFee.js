// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/description/

/*
Вам дан массив prices, где prices[i] — это цена данной акции на i-й день, а также целочисленная комиссия, представляющая
 комиссию за транзакцию.
Найдите максимальную прибыль, которую вы можете получить. Вы можете совершить столько транзакций, сколько захотите, но
 вам придется платить комиссию за каждую транзакцию.
Примечание:
Вы не можете совершать несколько транзакций одновременно (т.е. вы должны продать акции, прежде чем купить их снова).
Комиссия за транзакцию взимается только один раз за каждую покупку и продажу акций.
*/

/*
Input: prices = [1,3,2,8,4,9], fee = 2
Output: 8
Пояснение: Максимальную прибыль можно получить за счет:
- Покупка по ценам[0] = 1
- Продажа по ценам[3] = 8
- Покупка по ценам[4] = 4
- Продажа по ценам[5] = 9
Общая прибыль равна ((8 – 1) – 2) + ((9 – 4) – 2) = 8.
*/

/**
 * @param {number[]} prices
 * @param {number} fee
 * @return {number}
 */
var maxProfit = function(prices, fee) {
    let n = prices.length;
    let free = 0, hold = -prices[0];
    for (let i = 1; i < n; i++) {
        let tmp = hold;
        hold = Math.max(hold, free - prices[i]);
        free = Math.max(free, tmp + prices[i] - fee);
    }
    return free;
};

console.log(maxProfit([1,3,2,8,4,9], 2)); // 8