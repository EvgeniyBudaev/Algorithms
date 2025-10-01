package twoPointers.Sum3Closest;

import java.util.Arrays;

/* 16. 3Sum Closest
https://leetcode.com/problems/3sum-closest/description/

Учитывая целочисленный массив nums длины n и целочисленную цель, найдите три целых числа в nums, сумма которых наиболее
близка к цели. Верните сумму трех целых чисел. Вы можете предположить, что каждый вход будет иметь ровно одно решение.

Input: nums = [-1,2,1,-4], target = 1
Output: 2
Пояснение: Ближайшая к цели сумма равна 2. (-1 + 2 + 1 = 2).

Input: nums = [0,0,0], target = 1
Output: 0
Пояснение: Сумма, ближайшая к целевой, равна 0. (0 + 0 + 0 = 0).
*/

public class 3SumClosest {
    public static void main (String[]args){
        int[] nums = {-1, 2, 1, -4};
        System.out.println(threeSumClosest(nums, 1)); // 2
    }

    // threeSumClosest возвращает сумму трех целых чисел, ближайшую к цели.
    // time: O(n^2), так как мы сортируем массив и проходим по нему дважды.
    // space: O(1), так как мы не используем дополнительную память.
    private static int threeSumClosest ( int[] nums, int target){
        Arrays.sort(nums); // Сортируем массив чисел // [-4,-1,1,2]
        int n = nums.length; // Длина массива чисел
        int closestSum = nums[0] + nums[1] + nums[2]; // Инициализируем ближайшую сумму тремя первыми элементами массива

        // n-2, чтобы гарантировать, что после текущего элемента nums[i] останется как минимум два элемента для формирования тройки.
        for (int i = 0; i < n - 2; i++) {
            int left = i + 1, right = n - 1;

            while (left < right) {
                int sum = nums[i] + nums[left] + nums[right];
                if (sum == target) { // Если сумма равна цели, возвращаем её
                    return sum;
                }
                // Обновляем closestSum, если текущая сумма ближе к цели
                if (Math.abs(sum - target) < Math.abs(closestSum - target)) {
                    closestSum = sum;
                }
                if (sum < target) { // Если сумма меньше цели, двигаем левый указатель вправо
                    left++;
                } else { // Если сумма больше цели, двигаем правый указатель влево
                    right--;
                }
            }
        }

        return closestSum;
    }
}
