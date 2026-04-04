package leetcode.java.dynamicProgramming.PartitionEqualSubsetSum;

/* 416. Partition Equal Subset Sum
https://leetcode.com/problems/partition-equal-subset-sum/description/

Учитывая числа целочисленного массива, верните true, если вы можете разделить массив на два подмножества так, чтобы
сумма элементов в обоих подмножествах была равна или false в противном случае.

Input: nums = [1,5,11,5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11].

Input: nums = [1,2,3,5]
Output: false
Explanation: The array cannot be partitioned into equal sum subsets.
*/

public class PartitionEqualSubsetSum {
    public static void main(String[] args) {
        System.out.println(canPartition(new int[]{1, 5, 11, 5})); // true (11 = 1 + 5 + 5)
    }

    // canPartition проверяет, можно ли разбить массив на две части с равной суммой
    // time: O(n * s), где n - количество элементов, s - целевая сумма (totalSum/2)
    // space: O(s), где s - целевая сумма (размер DP массива)
    private static boolean canPartition(int[] nums) {
        // Вычисляем общую сумму элементов
        int totalSum = 0;
        for (int num : nums) {
            totalSum += num;
        }

        // Если сумма нечетная - разбиение невозможно
        if (totalSum % 2 != 0) {
            return false;
        }

        int target = totalSum / 2;
        // Создаем массив для динамического программирования
        boolean[] dp = new boolean[target + 1];
        dp[0] = true; // Базовый случай - сумма 0 всегда достижима

        // Перебираем все числа в массиве
        for (int num : nums) {
            // Идем от целевой суммы вниз до текущего числа
            for (int j = target; j >= num; j--) {
                // Если можно достичь суммы (j - num), то можно достичь и суммы j
                if (dp[j - num]) {
                    dp[j] = true;
                }
            }
        }

        return dp[target];
    }
}
