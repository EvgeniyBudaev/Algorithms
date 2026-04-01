package leetcode.java.dynamicProgramming.NumberOfLongestIncreasingSubsequence;

import java.util.Arrays;

/* 673. Number of Longest Increasing Subsequence
https://leetcode.com/problems/number-of-longest-increasing-subsequence/description/

Учитывая целочисленный массив nums, верните количество самых длинных возрастающих подпоследовательностей.
Обратите внимание, что последовательность должна быть строго возрастающей.

Input: nums = [1,3,5,4,7]
Output: 2
Пояснение: Две самые длинные возрастающие подпоследовательности — это [1, 3, 4, 7] и [1, 3, 5, 7].

Input: nums = [2,2,2,2,2]
Output: 5
Пояснение: Длина самой длинной возрастающей подпоследовательности равна 1, существует 5 возрастающих
подпоследовательностей длины 1, поэтому выведите 5.
*/

public class NumberOfLongestIncreasingSubsequence {
    public static void main(String[] args) {
        System.out.println(findNumberOfLIS(new int[]{1, 3, 5, 4, 7})); // 2
    }

    // findNumberOfLIS находит количество самых длинных возрастающих подпоследовательностей
    // time: O(n^2), space: O(n)
    private static int findNumberOfLIS(int[] nums) {
        if (nums == null || nums.length == 0) {
            return 0;
        }

        // lengths - хранит длины НВП, заканчивающихся на каждом индексе
        int[] lengths = new int[nums.length];
        // counts - хранит количество НВП для каждой длины
        int[] counts = new int[nums.length];

        // Инициализация: каждый элемент сам по себе является подпоследовательностью длины 1
        Arrays.fill(lengths, 1);
        Arrays.fill(counts, 1);

        for (int i = 0; i < nums.length; i++) {
            for (int j = 0; j < i; j++) {
                if (nums[j] < nums[i]) {
                    // Случай 1: нашли более длинную подпоследовательность
                    if (lengths[j] + 1 > lengths[i]) {
                        lengths[i] = lengths[j] + 1;
                        counts[i] = counts[j];
                    } else if (lengths[j] + 1 == lengths[i]) {
                        // Случай 2: нашли еще одну подпоследовательность такой же длины
                        counts[i] += counts[j];
                    }
                }
            }
        }

        // Находим максимальную длину
        int maxLength = 0;
        for (int l : lengths) {
            if (l > maxLength) {
                maxLength = l;
            }
        }

        // Суммируем количество подпоследовательностей с максимальной длиной
        int result = 0;
        for (int i = 0; i < counts.length; i++) {
            if (lengths[i] == maxLength) {
                result += counts[i];
            }
        }

        return result;
    }
}
