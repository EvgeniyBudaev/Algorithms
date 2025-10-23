package array.IntersectionOfTwoArrays;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

/* 349. Intersection of Two Arrays
https://leetcode.com/problems/intersection-of-two-arrays/description/

Даны два целочисленных массива nums1 и nums2, вернуть массив их пересечения. Каждый элемент в результате должен быть
уникальным, и вы можете вернуть результат в любом порядке.

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
*/


public class IntersectionOfTwoArrays {
    public static void main(String[] args) {
        int[] nums1 = {1, 2, 2, 1};
        int[] nums2 = {2, 2};
        System.out.println(Arrays.toString(intersection(nums1, nums2))); // [2]
    }

    // intersection возвращает массив пересечения элементов двух массивов nums1 и nums2.
    // time: O(n+m), space: O(n)
    private static int[] intersection(int[] nums1, int[] nums2) {
        // Создаем множества из массивов
        Set<Integer> set1 = new HashSet<>();
        for (int num : nums1) {
            set1.add(num);
        }

        Set<Integer> set2 = new HashSet<>();
        for (int num : nums2) {
            set2.add(num);
        }

        // Находим пересечение множеств
        List<Integer> resultList = new ArrayList<>();
        for (int num : set1) {
            if (set2.contains(num)) { // Если элемент присутствует в обоих множествах, добавляем его в результат
                resultList.add(num);
            }
        }

        // Конвертируем List в массив
        int[] result = new int[resultList.size()];
        for (int i = 0; i < resultList.size(); i++) {
            result[i] = resultList.get(i);
        }

        return result; // Возвращаем результат
    }
}
