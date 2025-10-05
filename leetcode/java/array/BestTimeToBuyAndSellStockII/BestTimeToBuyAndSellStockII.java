package array.BestTimeToBuyAndSellStockII;

/* 122. Best Time to Buy and Sell Stock II
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/

Вам дан целочисленный массив цен, где цены[i] — это цена данной акции на i-й день.

Каждый день вы можете решить купить и/или продать акцию. Вы можете держать не более одной акции в любой момент времени.
Однако вы можете купить ее, а затем немедленно продать в тот же день.

Найдите и верните максимальную прибыль, которую вы можете получить.

Input: prices = [7,1,5,3,6,4]
Output: 7
Объяснение: Покупаем на 2-й день (цена = 1) и продаем на 3-й день (цена = 5), прибыль = 5-1 = 4.
Затем покупаем на 4-й день (цена = 3) и продаем на 5-й день (цена = 6), прибыль = 6-3 = 3.
Общая прибыль составляет 4 + 3 = 7.
*/

public class BestTimeToBuyAndSellStockII {
    public static void main(String[] args) {
        int[] prices = {7, 1, 5, 3, 6, 4};
        System.out.println(maxProfit(prices)); // 7  // 4 + 3
    }

    // maxProfit находим максимальную прибыль.
    // time: O(n), space: O(1)
    private static int maxProfit(int[] prices) {
        int profit = 0; // 0 - если нет прибыли

        for (int i = 1; i < prices.length; i++) {
            // если цена на i-й день больше цены на i-1-й день, то прибавляем разницу к прибыли
            if (prices[i] > prices[i - 1]) {
                profit += prices[i] - prices[i - 1];
            }
        }

        return profit; // Возвращаем максимальную прибыль
    }
}
