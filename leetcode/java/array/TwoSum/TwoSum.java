package array.TwoSum;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/* 1. Two Sum
https://leetcode.com/problems/two-sum/description/

Учитывая массив целых чисел nums и целочисленную target, верните индексы двух чисел так, чтобы их сумма составляла
target. Вы можете предположить, что каждый вход будет иметь ровно одно решение, и вы не можете использовать один и тот
же элемент дважды. Вы можете вернуть ответ в любом порядке.

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Input: nums = [3,2,4], target = 6
Output: [1,2]

Input: nums = [3,3], target = 6
Output: [0,1]
*/

public class TwoSum {
    public static void main(String[] args) {
        int[] nums = {2, 7, 11, 15};
        System.out.println(Arrays.toString(twoSum(nums, 9))); // [0,1]
    }

    // twoSum возвращает индексы двух чисел так, чтобы их сумма составляла target.
    // time: O(n), space: O(n)
    private static int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> numMap = new HashMap<>(); // Создаем map для хранения чисел и их индексов

        for (int i = 0; i < nums.length; i++) {
            int num = nums[i];
            int diff = target - num; // Вычисляем разницу между целевым значением и текущим числом
            // Проверяем, есть ли разница в map
            if (numMap.containsKey(diff)) {
                // Если есть, возвращаем индексы
                return new int[]{numMap.get(diff), i};
            }
            numMap.put(num, i); // Сохраняем текущее число и его индекс в map
        }

        return new int[0]; // Если решение не найдено, возвращаем пустой массив
    }
}