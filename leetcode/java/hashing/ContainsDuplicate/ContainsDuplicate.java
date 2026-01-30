package leetcode.java.hashing.ContainsDuplicate;

import java.util.HashSet;
import java.util.Set;

/* 217. Contains Duplicate
https://leetcode.com/problems/contains-duplicate/description/

Учитывая числа целочисленного массива, верните true, если какое-либо значение встречается в массиве хотя бы дважды, и
верните false, если каждый элемент различен.

Input: nums = [1,2,3,1]
Output: true

Input: nums = [1,2,3,4]
Output: false

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true
*/

public class ContainsDuplicate {
    public static void main(String[] args) {
        System.out.println(containsDuplicate(new int[]{1, 2, 3, 1})); // true
        System.out.println(containsDuplicate(new int[]{1, 2, 3, 4})); // false
        System.out.println(containsDuplicate(new int[]{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})); // true
    }

    // containsDuplicate проверяет, содержит ли массив повторяющиеся элементы.
    // time: O(n), space: O(n)
    private static boolean containsDuplicate(int[] nums) {
         Set<Integer> seen = new HashSet<>(); // Множество для хранения уже встреченных значений

        for (int num : nums) { // Проходим по каждому элементу массива
            // Если значение уже встречалось, возвращаем true
            if (!seen.add(num)) {
                return true;
            }
            // Метод add() возвращает false, если элемент уже существует в множестве
        }

        return false; // Если ни одно значение не повторилось, возвращаем false 
    }
}
