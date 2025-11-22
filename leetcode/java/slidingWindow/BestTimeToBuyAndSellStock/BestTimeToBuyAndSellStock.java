package slidingWindow.BestTimeToBuyAndSellStock;

/* 121. Best Time to Buy and Sell Stock
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/

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

public class BestTimeToBuyAndSellStock {
    public static void main(String[] args) {
        int[] prices = {7, 1, 5, 3, 6, 4};
        System.out.println(maxProfit(prices)); // 5
    }

    // maxProfit вычисляет максимальную возможную прибыль от покупки и продажи акций на основе массива цен.
    // time: O(n), space: O(1)
    private static int maxProfit(int[] prices) {
        int left = 0, right = 1, profit = 0; // left - индекс дня покупки, right - индекс дня продажи

        while (right < prices.length) { // проходимся по массиву цен
            // если цена акции меньше или равна цене акции в день покупки,
            // то сдвигаем левую границу на один шаг вправо
            if (prices[right] <= prices[left]) {
                left = right;
            }
            int window = prices[right] - prices[left]; // прибыль от текущей сделки
            profit = Math.max(profit, window);         // обновляем максимальную прибыль
            right++;
        }

        return profit; // возвращаем максимальную прибыль
    }
}
