package leetcode.java.dynamicProgramming.LongestIncreasingSubsequence;

import java.util.Arrays;

/* 300. Longest Increasing Subsequence
https://leetcode.com/problems/longest-increasing-subsequence/description/

Учитывая целочисленный массив nums, верните длину самого длинного, строго увеличивающегося последовательность.

Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Пояснение: Самая длинная возрастающая подпоследовательность — [2,3,7,101], поэтому длина равна 4.

Input: nums = [0,1,0,3,2,3]
Output: 4

Input: nums = [7,7,7,7,7,7,7]
Output: 1
*/

public class LongestIncreasingSubsequence {
    public static void main(String[] args) {
        System.out.println(lengthOfLIS(new int[]{10, 9, 2, 5, 3, 7, 101, 18})); // 4
    }

    // lengthOfLIS находит длину наибольшей возрастающей подпоследовательности.
    // time: O(n^2), space: O(n)
    private static int lengthOfLIS(int[] nums) {
        // Если массив пуст или null, возвращаем 0
        if (nums == null || nums.length == 0) {
            return 0;
        }

        // Создаем массив для хранения длин НВП для каждой позиции
        int[] dp = new int[nums.length];
        // Заполняем массив единицами (каждый элемент сам по себе - подпоследовательность длины 1)
        Arrays.fill(dp, 1);

        for (int i = 0; i < nums.length; i++) {
            // Проверяем все предыдущие элементы
            for (int j = 0; j < i; j++) {
                if (nums[i] > nums[j]) { // Если текущий элемент больше предыдущего
                    // Обновляем длину НВП для текущей позиции
                    dp[i] = Math.max(dp[i], dp[j] + 1);
                }
            }
        }

        // Находим максимальное значение в массиве dp
        int maxLen = 0; // Инициализируем максимальную длину
        for (int v : dp) {
            if (v > maxLen) { // Если текущее значение больше максимального, обновляем
                maxLen = v;
            }
        }

        return maxLen; // Возвращаем максимальную длину НВП
    }
}
