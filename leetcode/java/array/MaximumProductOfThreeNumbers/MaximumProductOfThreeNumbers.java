package array.MaximumProductOfThreeNumbers;

import java.util.Arrays;

/* 628. Maximum Product of Three Numbers
https://leetcode.com/problems/maximum-product-of-three-numbers/description/

Дан целочисленный массив nums. Найдите три числа, произведение которых максимально, и верните максимальное произведение.

Input: nums = [1,2,3]
Output: 6

Input: nums = [1,2,3,4]
Output: 24

Input: nums = [-1,-2,-3]
Output: -6

Input: nums = [-100,-98,-1,2,3,4]
Output: 39200
*/

public class MaximumProductOfThreeNumbers {
    public static void main(String[] args) {
        int[] nums = {1, 2, 3};
        System.out.println(maximumProduct(nums)); // 6
    }

    // maximumProduct возвращает максимальное произведения трех чисел в массиве nums.
    // time: O(n * log(n)), space: O(1)
    private static int maximumProduct(int[] nums) {
        Arrays.sort(nums); // Сортируем массив в порядке возрастания
        int n = nums.length; // Получаем длину массива

        // Вариант 1: три самых больших числа
        int option1 = nums[n - 1] * nums[n - 2] * nums[n - 3];

        // Вариант 2: два самых маленьких (возможно отрицательных) и самое большое
        int option2 = nums[0] * nums[1] * nums[n - 1];

        // Возвращаем максимальное из двух вариантов
        return Math.max(option1, option2);
    }
}
