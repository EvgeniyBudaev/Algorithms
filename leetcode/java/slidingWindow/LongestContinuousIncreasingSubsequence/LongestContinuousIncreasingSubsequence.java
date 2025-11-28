package slidingWindow.LongestContinuousIncreasingSubsequence;

/* 674. Longest Continuous Increasing Subsequence
https://leetcode.com/problems/longest-continuous-increasing-subsequence/description/

Дан несортированный массив целых чисел nums, вернуть длину самой длинной непрерывной возрастающей подпоследовательности
(т. е. подмассива). Подпоследовательность должна быть строго возрастающей.
Непрерывная возрастающая подпоследовательность определяется двумя индексами l и r (l < r),
такими, что она равна [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] и для каждого l <= i < r, nums[i] < nums[i + 1].

Input: nums = [1,3,5,4,7]
Output: 3
Объяснение: Самая длинная непрерывная возрастающая подпоследовательность — [1,3,5] с длиной 3.
Хотя [1,3,5,7] является возрастающей подпоследовательностью, она не является непрерывной, поскольку элементы 5 и 7
разделены элементом 4.
*/

public class LongestContinuousIncreasingSubsequence {
    public static void main(String[] args) {
        int[] nums = {1, 3, 5, 4, 7};
        System.out.println(findLengthOfLCIS(nums)); // 3
    }

    // findLengthOfLCIS находит длину самой длинной непрерывной возрастающей подпоследовательности в массиве nums.
    // time: O(n), space: O(1)
    private static int findLengthOfLCIS(int[] nums) {
        if (nums == null || nums.length == 0) { // Если массив nums пуст, возвращаем 0
            return 0;
        }

        // left - левая граница текущей подпоследовательности,
        // maxLen - максимальная длина найденной подпоследовательности
        int left = 0, maxLen = 1;

        for (int right = 1; right < nums.length; right++) {
            // Если текущий элемент не больше предыдущего, сбрасываем left
            if (nums[right] <= nums[right - 1]) {
                left = right;
            }
            // Обновляем maxLen, если текущая подпоследовательность длиннее
            maxLen = Math.max(maxLen, right - left + 1);
        }

        return maxLen; // Возвращаем максимальную длину найденной подпоследовательности
    }
}
