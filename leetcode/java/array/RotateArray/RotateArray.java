package array.RotateArray;

import java.util.Arrays;

/* 189. Rotate Array
https://leetcode.com/problems/rotate-array/description/

Дан целочисленный массив nums. Повернуть массив вправо на k шагов, где k — неотрицательное число.

Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]

Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]
*/

public class RotateArray {
    public static void main(String[] args) {
        int[] nums1 = {1, 2, 3, 4, 5, 6, 7};
        int k1 = 3;
        rotate(nums1, k1);
        System.out.println(Arrays.toString(nums1)); // [5,6,7,1,2,3,4]

        int[] nums2 = {-1, -100, 3, 99};
        int k2 = 2;
        rotate(nums2, k2);
        System.out.println(Arrays.toString(nums2)); // [3,99,-1,-100]
    }

    // rotate поворачивает массив от индекса 0 до индекса len(nums)-1.
    // time: O(n) — проходим по всем элементам, space: O(1) — не используем дополнительную память.
    private static void rotate(int[] nums, int k) {
        int n = nums.length; // Длина массива
        k = k % n;           // Нормализуем k, чтобы он был в пределах длины массива
        if (k == 0) {        // Если k = 0, вращение не нужно
            return;
        }
        rotateSubArray(nums, 0, n); // Поворачиваем весь массив. // [7, 6, 5, 4, 3, 2, 1]
        rotateSubArray(nums, 0, k); // [5, 6, 7, 4, 3, 2, 1]
        rotateSubArray(nums, k, n); // [5, 6, 7, 1, 2, 3, 4]
    }

    // rotateSubArray поворачивает массив от индекса i до индекса j (не включительно).
    // time: O(n) — проходим по всем элементам, space: O(1) — не используем дополнительную память.
    public static void rotateSubArray(int[] nums, int i, int j) {
        j -= 1;      // Чтобы не включать последний индекс в поворот
        while (i < j) { // Пока индексы не встретятся
            int temp = nums[i];
            nums[i] = nums[j];
            nums[j] = temp;
            i++;
            j--;
        }
    }
}
