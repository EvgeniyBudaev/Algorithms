package leetcode.java.dynamicProgramming.CoinChange;

/* 322. Coin Change
https://leetcode.com/problems/coin-change/description/

Вам дан целочисленный массив coin, представляющий монеты разного номинала, и целочисленная сумма, представляющая общую
 сумму денег.
Верните наименьшее количество монет, которое вам нужно, чтобы составить эту сумму. Если эта сумма денег не может быть
 составлена ни одной комбинацией монет, верните -1.
Вы можете предположить, что у вас есть бесконечное количество монет каждого вида.

Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1

Input: coins = [2], amount = 3
Output: -1

Input: coins = [1], amount = 0
Output: 0
*/

public class CoinChange {
    public static void main(String[] args) {
        System.out.println(coinChange(new int[]{1, 2, 5}, 11)); // 3
    }

    // coinChange вычисляет минимальное количество монет, необходимых для составления суммы.
    // time: O(n*m) где n - количество монет, а m - сумма
    // space: O(m) где m - сумма
    private static int coinChange(int[] coins, int amount) {
        // Если сумма равна 0, монет не требуется
        if (amount < 1) {
            return 0;
        }

        // Создаем массив для хранения минимального количества монет для каждой суммы
        // Изначально заполняем "бесконечностью" (недостижимым значением)
        int[] dp = new int[amount + 1]; // +1 потому что индексы начинаются с 0
        for (int i = 0; i < dp.length; i++) {
            dp[i] = Integer.MAX_VALUE; // "бесконечность" (MAX_VALUE)
        }

        dp[0] = 0; // Базовый случай: для суммы 0 нужно 0 монет

        // Перебираем все номиналы монет
        for (int coin : coins) {
            // Для текущего номинала обновляем значения для всех сумм от coin до amount
            for (int i = coin; i <= amount; i++) {
                // Минимум между текущим значением и значением для суммы (i-coin) + 1 монета
                if (dp[i - coin] != Integer.MAX_VALUE && dp[i - coin] + 1 < dp[i]) {
                    dp[i] = dp[i - coin] + 1; // Обновляем значение в массиве
                }
            }
        }

        // Если последний элемент остался "бесконечным", значит сумму составить невозможно
        if (dp[amount] == Integer.MAX_VALUE) {
            return -1;
        }

        return dp[amount]; // Возвращаем минимальное количество монет для суммы amount
    }
}
