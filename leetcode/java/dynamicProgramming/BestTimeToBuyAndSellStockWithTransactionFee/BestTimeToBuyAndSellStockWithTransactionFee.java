package leetcode.java.dynamicProgramming.BestTimeToBuyAndSellStockWithTransactionFee;

/* 714. Best Time to Buy and Sell Stock with Transaction Fee
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/description/

Вам дан массив prices, где prices[i] — это цена данной акции на i-й день, а также целочисленная комиссия, представляющая
 комиссию за транзакцию.
Найдите максимальную прибыль, которую вы можете получить. Вы можете совершить столько транзакций, сколько захотите, но
 вам придется платить комиссию за каждую транзакцию.
Примечание:
Вы не можете совершать несколько транзакций одновременно (т.е. вы должны продать акции, прежде чем купить их снова).
Комиссия за транзакцию взимается только один раз за каждую покупку и продажу акций.

Input: prices = [1,3,2,8,4,9], fee = 2
Output: 8
Пояснение: Максимальную прибыль можно получить за счет:
- Покупка по ценам[0] = 1
- Продажа по ценам[3] = 8
- Покупка по ценам[4] = 4
- Продажа по ценам[5] = 9
Общая прибыль равна ((8 – 1) – 2) + ((9 – 4) – 2) = 8.
*/

public class BestTimeToBuyAndSellStockWithTransactionFee {
    public static void main(String[] args) {
        System.out.println(maxProfit(new int[]{1, 3, 2, 8, 4, 9}, 2)); // 8
    }

    // maxProfit вычисляет максимальную прибыль от торговли акциями с учетом комиссии за транзакцию.
    // time: O(n), space: O(1)
    private static int maxProfit(int[] prices, int fee) {
        int n = prices.length;
        if (n < 2) {
            return 0; // Невозможно совершить сделку при менее чем 2 днях
        }

        // free - максимальная прибыль, когда у нас нет акции в конце дня (можем купить на следующий день)
        // hold - максимальная прибыль, когда у нас есть акция в конце дня
        int free = 0;
        int hold = -prices[0]; // Покупаем акцию в первый день

        for (int i = 1; i < n; i++) {
            // Сохраняем предыдущее значение hold перед обновлением
            int prevHold = hold;

            // Обновляем hold: либо продолжаем держать акцию, либо покупаем сегодня (из состояния free)
            hold = Math.max(hold, free - prices[i]);

            // Обновляем free: либо остаемся без акции, либо продаем сегодня (из состояния hold) с учетом комиссии
            free = Math.max(free, prevHold + prices[i] - fee);
        }

        // В конце оптимально остаться без акции (состояние free)
        return free;
    }
}
