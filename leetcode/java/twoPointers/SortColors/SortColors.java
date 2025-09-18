package twoPointers.SortColors;

import java.util.Arrays;

/* 75. Sort Colors
https://leetcode.com/problems/sort-colors/description/

Дан массив nums с n объектами красного, белого или синего цвета, отсортируйте их по месту так, чтобы объекты одного
цвета были соседними, а цвета располагались в порядке красного, белого и синего.
Мы будем использовать целые числа 0, 1 и 2 для обозначения красного, белого и синего цветов соответственно.
Вы должны решить эту проблему, не используя функцию сортировки библиотеки.

Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]

Input: nums = [2,0,1]
Output: [0,1,2]
*/

public class SortColors {
    public static void main(String[] args) {
        int[] nums1 = {2, 0, 2, 1, 1, 0};
        sortColors(nums1);
        System.out.println(Arrays.toString(nums1)); // [0,0,1,1,2,2]
        
        int[] nums2 = {2, 0, 1};
        sortColors(nums2);
        System.out.println(Arrays.toString(nums2)); // [0,1,2]
    }

    // sortColors сортирует массив целых чисел по возрастанию.
    // time: O(n), space: O(1)
    public static void sortColors(int[] nums) {
        int i = 0; // Индекс текущего элемента
        int left = 0, right = nums.length - 1; // Левая и правая границы

        while (left < right && i <= right) { // пока левая граница меньше правой и индекс меньше правой границы
            // Если текущий элемент равен 0, то меняем его местами с элементом в начале массива
            if (nums[i] == 0) {
                swap(nums, left, i);
                left++;
                i++;
                // Если текущий элемент равен 2, то меняем его местами с элементом в конце массива
            } else if (nums[i] == 2) {
                swap(nums, right, i);
			    right--;
                // Если текущий элемент равен 1, то оставляем его на месте и переходим к следующему элементу
            } else {
                i++;
            }
        }
    }

    // swap функция для обмена элементов.
    private static void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }
}
