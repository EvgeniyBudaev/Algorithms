package leetcode.java.hashing.ContiguousArray;

import java.util.HashMap;
import java.util.Map;

/* 525. Contiguous Array
https://leetcode.com/problems/contiguous-array/description/

Учитывая числа двоичного массива, верните максимальную длину непрерывного подмассива с равным количеством 0 и 1.

Input: nums = [0,1]
Output: 2
Пояснение: [0, 1] — это самый длинный непрерывный подмассив с равным количеством 0 и 1.

Input: nums = [0,1,0]
Output: 2
Пояснение: [0, 1] (или [1, 0]) — это самый длинный непрерывный подмассив с равным количеством 0 и 1.
*/

public class ContiguousArray {
    public static void main(String[] args) {
        int[] nums1 = {0, 1};
        System.out.println(findMaxLength(nums1)); // 2

        int[] nums2 = {0, 1, 0};
        System.out.println(findMaxLength(nums2)); // 2

        int[] nums3 = {0, 1, 1, 0, 1, 1, 1, 0};
        System.out.println(findMaxLength(nums3)); // 4
    }

    // findMaxLength находит максимальную длину непрерывного подмассива с равным количеством 0 и 1.
    // time: O(n), space: O(n)
    private static int findMaxLength(int[] nums) {
        int maxLen = 0; // Текущая максимальная длина подмассива
        int count = 0; 	// count - Отслеживает разницу между количеством единиц и нулей в рассматриваемой части массива.
	                    // Если count положительно, это означает, что больше единиц, чем нулей; если отрицательно — наоборот.
        Map<Integer, Integer> countMap = new HashMap<>(); // Хранит индексы для каждой уникальной разницы между единицами и нулями
        countMap.put(0, -1); // Инициализация для случая, когда count равен 0.

        for (int i = 0; i < nums.length; i++) {
            // Увеличиваем count для 1, уменьшаем для 0
            count += (nums[i] == 1) ? 1 : -1;

            // Поиск подмассивов с равным количеством нулей и единиц
            // Если да, то вычисляем длину текущего подмассива как разницу между текущим индексом i и индексом последнего
            // вхождения такого же count (map[count]). Обновляем maxLen, если текущая длина больше предыдущего максимума.
            if (countMap.containsKey(count)) {
                int currentLen = i - countMap.get(count);
                maxLen = Math.max(maxLen, currentLen);
            } else {
                // Если нет, то сохраняем текущий индекс i в map под ключом count.
                countMap.put(count, i);
            }
        }

        return maxLen;
    }
}
