package leetcode.java.dynamicProgramming.MaximumProductSubarray;

/* 152. Maximum Product Subarray
https://leetcode.com/problems/maximum-product-subarray/description/

Учитывая целочисленный массив чисел, найдите подмассив который имеет самый большой продукт, и возвращает продукт.
Тестовые случаи генерируются таким образом, чтобы ответ соответствовал 32-битному целому числу.

Input: nums = [2,3,-2,4]
Output: 6
Пояснение: [2,3] имеет наибольшее произведение 6.

Input: nums = [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
*/

public class MaximumProductSubarray {
    public static void main(String[] args) {
        System.out.println(maxProduct(new int[]{2, 3, -2, 4})); // 6 (подмассив [2,3])
    }

    // maxProduct находит максимальное произведение подмассива
    // time: O(n), space: O(1)
    private static int maxProduct(int[] nums) {
        if (nums == null || nums.length == 0) {
            return 0;
        }

        int maxSoFar = nums[0]; // максимальное произведение на текущий момент
        int minSoFar = nums[0]; // минимальное произведение на текущий момент (нужно для отрицательных чисел)
        int result = maxSoFar; // итоговый результат

        for (int i = 1; i < nums.length; i++) {
            int current = nums[i];

            // Временная переменная для хранения нового maxSoFar перед обновлением.
            // Это необходимо, потому что для расчета minSoFar нам нужно старое значение maxSoFar.
            int tempMaxSoFar = Math.max(
                current,
                Math.max(
                    maxSoFar * current,
                    minSoFar * current
                )
            );

            // Обновляем minSoFar с учетом текущего числа
            minSoFar = Math.min(
                current,
                Math.min(
                    maxSoFar * current,
                    minSoFar * current
                )
            );

            // Обновляем maxSoFar
            maxSoFar = tempMaxSoFar;

            // Обновляем результат, если нашли большее произведение
            if (maxSoFar > result) {
                result = maxSoFar;
            }
        }

        return result;
    }
}
