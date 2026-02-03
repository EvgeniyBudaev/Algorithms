package leetcode.java.hashing.FindAllNumbersDisappearedInAnArray;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

/* 448. Find All Numbers Disappeared in an Array
https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/

Учитывая массив nums из n целых чисел, где nums[i] находится в диапазоне [1, n], верните массив всех целых чисел в
диапазоне [1, n], которые не встречаются в nums.

Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]

Input: nums = [1,1]
Output: [2]
*/

public class FindAllNumbersDisappearedInAnArray {
    public static void main(String[] args) {
        int[] nums = {4, 3, 2, 7, 8, 2, 3, 1};
        System.out.println(findDisappearedNumbers(nums)); // [5, 6]
    }

    // findDisappearedNumbers находит пропущенные числа в массиве nums.
    // time: O(n), space: O(n)
    private static List<Integer> findDisappearedNumbers(int[] nums) {
        Set<Integer> numSet = new HashSet<>(); // Множество для отслеживания существующих чисел
        List<Integer> result = new ArrayList<>(); // Список для хранения пропущенных чисел

        // Отмечаем все существующие числа во множестве
        for (int num : nums) {
            numSet.add(num);
        }

        // Проверяем наличие пропущенных цифр от 1 до n
        for (int i = 1; i <= nums.length; i++) {
            // Если число не отмечено на карте, добавляем его в множество
            if (!numSet.contains(i)) {
                result.add(i);
            }
        }

        return result; // Возвращаем срез с пропущенными числами
    }
}
