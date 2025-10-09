package array.FindAllDuplicatesInAnArray;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/* 442. Find All Duplicates in an Array
https://leetcode.com/problems/find-all-duplicates-in-an-array/description/

Учитывая целочисленный массив nums длины n, где все целые числа nums находятся в диапазоне [1, n] и каждое целое число
встречается один или два раза, верните массив всех целых чисел, которые встречаются дважды.
Вы должны написать алгоритм, который работает за время O(n) и использует только постоянное дополнительное пространство.

Input: nums = [4,3,2,7,8,2,3,1]
Output: [2,3]

Input: nums = [1,1,2]
Output: [1]

Input: nums = [1]
Output: []
*/

public class FindAllDuplicatesInAnArray {
    public static void main(String[] args) {
        int[] nums = {4, 3, 2, 7, 8, 2, 3, 1};
        System.out.println(findDuplicates(nums)); // [2,3]
    }

    // findDuplicates находит дубликаты чисел.
    // time: O(n), space: O(1)
    private static List<Integer> findDuplicates(int[] nums) {
        Arrays.sort(nums);  // [1, 2, 2, 3, 3, 4, 7, 8]
        List<Integer> result = new ArrayList<>(); // Создаем пустой слайс для хранения дубликатов

        for (int i = 1; i < nums.length; i++) {
            // Если текущий элемент равен предыдущему, добавляем его в результат
            if (nums[i] == nums[i - 1]) {
                result.add(nums[i]);
            }
        }

        return result; // Возвращаем список дубликатов
    }
}
