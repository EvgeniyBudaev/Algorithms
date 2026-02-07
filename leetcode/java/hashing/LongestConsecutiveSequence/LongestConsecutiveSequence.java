package leetcode.java.hashing.LongestConsecutiveSequence;

import java.util.Arrays;

/* 128. Longest Consecutive Sequence
https://leetcode.com/problems/longest-consecutive-sequence/description/

Учитывая несортированный массив целых чисел nums, верните длину самой длинной последовательной последовательности
элементов.
Вы должны написать алгоритм, который работает за время O(n).

Input: nums = [100,4,200,1,3,2]
Output: 4
Объяснение: Самая длинная последовательная последовательность элементов — [1,2,3,4]. Следовательно, его длина равна 4

Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9
*/

public class LongestConsecutiveSequence {
    public static void main(String[] args) {
        int[] nums = {100, 4, 200, 1, 3, 2};
        System.out.println(longestConsecutive(nums)); // 4
    }

    // longestConsecutive находит длину самой длинной последовательности последовательных чисел.
    // time: O(n), space: O(1)
    private static int longestConsecutive(int[] nums) {
        if (nums == null || nums.length == 0) { // Если массив пуст, возвращаем 0.
            return 0;
        }

        // Сортируем массив по возрастанию
        Arrays.sort(nums); // time: O(n log n), space: O(log n) для сортировки
        return search(nums); // Ищем самую длинную последовательность и возвращаем
    }

    // search проходит по отсортированному массиву и находит самую длинную последовательность.
    // time: O(n), space: O(1)
    private static int search(int[] nums) {
        int maxScore = 1;   // Максимальная длина последовательности
        int currScore = 1;  // Текущая длина последовательности

        for (int i = 1; i < nums.length; i++) {
            // Пропускаем дубликаты
            if (nums[i - 1] == nums[i]) {
                continue;
            }

            // Проверяем, является ли текущее число следующим в последовательности
            if (nums[i] == nums[i - 1] + 1) {
                currScore++; // Увеличиваем текущую длину последовательности
            } else {
                // Обновляем максимальную длину если текущая последовательность закончилась
                maxScore = Math.max(maxScore, currScore);
                currScore = 1; // Сбрасываем счётчик
            }
        }

        // Проверяем последнюю последовательность (возможно, самую длинную)
        return Math.max(maxScore, currScore);
    }
}
