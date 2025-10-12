package array.FindTheDifferenceOfTwoArrays;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

/* 2215. Find the Difference of Two Arrays
https://leetcode.com/problems/find-the-difference-of-two-arrays/description/

Учитывая два целочисленных массива nums1 и nums2 с индексом 0, верните ответ в виде списка размером 2, где:

ответ[0] — это список всех различных целых чисел в nums1, которых нет в nums2.
ответ[1] — это список всех различных целых чисел в nums2, которых нет в nums1.
Обратите внимание, что целые числа в списках могут возвращаться в любом порядке.

Input: nums1 = [1,2,3], nums2 = [2,4,6]
Output: [[1,3],[4,6]]
Объяснение:
Для nums1 nums1[1] = 2 присутствует в индексе 0 в nums2, тогда как nums1[0] = 1 и nums1[2] = 3 отсутствуют в nums2.
Следовательно, ответ[0] = [1,3].
Для nums2 nums2[0] = 2 присутствует в индексе 1 nums1, тогда как nums2[1] = 4 и nums2[2] = 6 отсутствуют в nums2.
Следовательно, ответ[1] = [4,6].
*/

public class FindTheDifferenceOfTwoArrays {
    public static void main(String[] args) {
        int[] arr1 = {1, 2, 3};
        int[] arr2 = {2, 4, 6};
        System.out.println(findDifference(arr1, arr2)); // [[1,3],[4,6]]
    }

    // findDifference возвращает списки элементов, которых нет в другом масиве.
    // time: O(n), space: O(n)
    private static List<List<Integer>> findDifference(int[] nums1, int[] nums2) {
        Set<Integer> set1 = new HashSet<>(); // Множество элементов первого массива
        Set<Integer> set2 = new HashSet<>(); // Множество элементов второго массива

        // Заполняем множества
        for (int num : nums1) set1.add(num);
        for (int num : nums2) set2.add(num);

        // Находим уникальные элементы для каждого списка
        List<Integer> unique1 = new ArrayList<>();
        List<Integer> unique2 = new ArrayList<>();

        for (int num : set1) {
            if (!set2.contains(num)) {
                unique1.add(num);
            }
        }

        for (int num : set2) {
            if (!set1.contains(num)) {
                unique2.add(num);
            }
        }

        return Arrays.asList(unique1, unique2); // Возвращаем списки
    }
}
