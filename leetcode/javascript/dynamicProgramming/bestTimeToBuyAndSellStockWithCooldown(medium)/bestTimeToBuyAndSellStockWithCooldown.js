// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/

/*
Вам дан массив цен, где цены[i] — это цена данной акции на i-й день.
Найдите максимальную прибыль, которую вы можете получить. Вы можете совершить столько транзакций, сколько захотите
 (т.е. купить одну и продать одну акцию несколько раз) со следующими ограничениями:
После того, как вы продадите свои акции, вы не сможете купить их на следующий день
 (т.е. время восстановления один день).
Примечание. Вы не можете совершать несколько транзакций одновременно
 (т.е. вы должны продать акции, прежде чем купить их снова).
*/

/*
Input: prices = [1,2,3,0,2]
Output: 3
Explanation: transactions = [buy, sell, cooldown, buy, sell]
*/

/*
Input: prices = [1]
Output: 0
*/

/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function(prices) {
    let [coolDown, sell, hold] = [0, 0, Number.NEGATIVE_INFINITY];

    for ( let stockPrice_Day_i of prices) {
        let [prevCoolDown, prevSell, prevHold] = [coolDown, sell, hold];

        // Максимальная прибыль от восстановления в День i получается либо от охлаждения Day_i-1, либо от продажи в
        // Day_i-1, и сегодня Day_i — день охлаждения.
        coolDown = Math.max(prevCoolDown, prevSell);

        // Максимальная прибыль от продажи в Day_i получается при удержании Day_i-1 и продаже в Day_i.
        sell = prevHold + stockPrice_Day_i;

        //  Максимальная прибыль от удержания Day_i получается либо от удержания Day_i-1, либо от охлаждения в Day_i-1 и
        //  покупки в Day_i.
        hold = Math.max( prevHold, prevCoolDown - stockPrice_Day_i );
    }

    // Действием заключительного торгового дня должна быть либо продажа, либо охлаждение.
    return Math.max(sell, coolDown);
};

console.log(maxProfit([1,2,3,0,2])); // 3