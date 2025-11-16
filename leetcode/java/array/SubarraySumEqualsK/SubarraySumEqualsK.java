package array.SubarraySumEqualsK;

import java.util.HashMap;
import java.util.Map;

/* 560. Subarray Sum Equals K
https://leetcode.com/problems/subarray-sum-equals-k/description/

Учитывая массив целых чисел nums и целое число k, верните общее количество подмассивов, сумма которых равна k.
Подмассив — это непрерывная непустая последовательность элементов массива.

Input: nums = [1,1,1], k = 2
Output: 2

Input: nums = [1,2,3], k = 3
Output: 2
*/

public class SubarraySumEqualsK {
    public static void main(String[] args) {
        int[] nums = {1, 2, 3};
        System.out.println(subarraySum(nums, 3)); // 2
    }

    // subarraySum возвращает количество подмассивов с суммой, равной k.
    // time: O(n), space: O(n)
    private static int subarraySum(int[] nums, int k) {
        Map<Integer, Integer> sumMap = new HashMap<>(); // Создаем map для хранения сумм и их количества
        sumMap.put(0, 1); // Инициализируем с суммой 0, встречающейся 1 раз

        int sum = 0;   // Текущая сумма
        int count = 0; // Количество подмассивов с суммой k

        for (int num : nums) { // Проходим по массиву
            sum += num; // Добавляем текущее число к текущей сумме
            // Если (sum - k) существует в map, увеличиваем count
            if (sumMap.containsKey(sum - k)) {
                count += sumMap.get(sum - k);
            }
            // Обновляем map с текущей суммой
            sumMap.put(sum, sumMap.getOrDefault(sum, 0) + 1);
        }

        return count; // Возвращаем количество подмассивов с суммой k
    }
}
