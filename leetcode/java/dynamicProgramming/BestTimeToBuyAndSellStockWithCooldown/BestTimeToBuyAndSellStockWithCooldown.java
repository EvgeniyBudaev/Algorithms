package leetcode.java.dynamicProgramming.BestTimeToBuyAndSellStockWithCooldown;

/* 309. Best Time to Buy and Sell Stock with Cooldown
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/

Вам дан массив цен, где цены[i] — это цена данной акции на i-й день.
Найдите максимальную прибыль, которую вы можете получить. Вы можете совершить столько транзакций, сколько захотите
 (т.е. купить одну и продать одну акцию несколько раз) со следующими ограничениями:
После того, как вы продадите свои акции, вы не сможете купить их на следующий день
 (т.е. время восстановления один день).
Примечание. Вы не можете совершать несколько транзакций одновременно
 (т.е. вы должны продать акции, прежде чем купить их снова).

Input: prices = [1,2,3,0,2]
Output: 3
Explanation: transactions = [buy, sell, cooldown, buy, sell]

Input: prices = [1]
Output: 0
*/

public class BestTimeToBuyAndSellStockWithCooldown {
    public static void main(String[] args) {
        System.out.println(maxProfit(new int[]{1, 2, 3, 0, 2})); // 3
    }

    // maxProfit вычисляет максимальную прибыль от торговли акциями.
    // time: O(n), space: O(1), n - длина массива цен
    private static int maxProfit(int[] prices) {
        int coolDown = 0; // прибыль, если в текущий день ничего не делаем (после продажи или бездействия)
        int sell = 0; // прибыль, если продаем акцию в текущий день
        int hold = Integer.MIN_VALUE; // прибыль, если держим акцию (купили ранее или покупаем сегодня)

        for (int price : prices) {
            // Сохраняем предыдущие значения состояний
            int prevCoolDown = coolDown;
            int prevSell = sell;
            int prevHold = hold;

            // Обновляем coolDown: Максимум между бездействием в предыдущий день и продажей в предыдущий день
            coolDown = Math.max(prevCoolDown, prevSell);

            // Обновляем sell: Прибыль = цена продажи + прибыль от удержания акции до этого дня
            sell = prevHold + price;

            // Обновляем hold: Либо продолжаем держать акцию, либо покупаем сегодня (если был coolDown)
            hold = Math.max(prevHold, prevCoolDown - price);
        }

        // В последний день оптимальная стратегия — либо продажа, либо бездействие
        return Math.max(sell, coolDown);
    }
}
