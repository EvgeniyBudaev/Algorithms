package slidingWindow.MinimumSizeSubarraySum;

/* 209. Minimum Size Subarray Sum
https://leetcode.com/problems/minimum-size-subarray-sum/description/

Учитывая массив положительных целых чисел nums и цель положительного целого числа, верните минимальную длину
подмассива, сумма которых больше или равна целевой. Если такого подмассива нет, вместо него верните 0.

Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Пояснение: Подмассив [4,3] имеет минимальную длину при условии ограничения задачи

Input: target = 4, nums = [1,4,4]
Output: 1

Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0
*/

public class MinimumSizeSubarraySum {
    public static void main(String[] args) {
        int[] nums = {2, 3, 1, 2, 4, 3};
        System.out.println(minSubArrayLen(7, nums)); // 2
    }

    // minSubArrayLen возвращает минимальную длину подмассива, сумма которой больше или равна целевой.
    // time: O(n), space: O(1)
    private static int minSubArrayLen(int target, int[] nums) {
        if (nums == null || nums.length == 0) {
            return 0;
        }

        int left = 0, right = 0;
        int minLength = Integer.MAX_VALUE; // Инициализируем максимально возможным значением
        int subarraySum = nums[0];         // Начинаем с первого элемента

        while (right < nums.length) {
            // Если текущая сумма больше или равна целевой
            if (subarraySum >= target) {
                minLength = Math.min(minLength, right - left + 1); // Обновляем минимальную длину, если текущая длина меньше
                subarraySum -= nums[left]; // Уменьшаем сумму, убирая левый элемент
                left++; // Сдвигаем левый указатель
            } else {
                right++; // Увеличиваем окно, добавляя правый элемент
                // Если правый указатель не вышел за пределы массива, добавляем элемент в сумму
                if (right < nums.length) {
                    subarraySum += nums[right];
                }
            }
        }

        // Если minLength так и не была обновлена, значит подмассив не найден
        // Иначе возвращаем минимальную длину
        return minLength == Integer.MAX_VALUE ? 0 : minLength;
    }
}
