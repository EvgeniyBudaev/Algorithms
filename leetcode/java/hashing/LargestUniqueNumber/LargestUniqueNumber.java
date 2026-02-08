package leetcode.java.hashing.LargestUniqueNumber;

import java.util.HashMap;
import java.util.Map;

/* 1133. Largest Unique Number

Учитывая массив целых чисел A, верните наибольшее целое число, которое встречается только один раз.
Если ни одно целое число не встречается один раз, верните -1.

Input: [5,7,3,9,4,9,8,3,1]
Output: 8
Пояснение:
Максимальное целое число в массиве — 9, но оно повторяется. Число 8 встречается только один раз, поэтому это ответ.

Input: [9,9,8,8]
Output: -1
Пояснение:
Нет числа, которое встречается только один раз.
*/

public class LargestUniqueNumber {
    public static void main(String[] args) {
        int[] nums = {5, 7, 3, 9, 4, 9, 8, 3, 1};
        System.out.println(largestUniqueNumber(nums)); // 8
    }

    // largestUniqueNumber находит наибольшее уникальное число в массиве.
    // Если такого числа нет, возвращает -1.
    // time: O(n), space: O(n)
    private static int largestUniqueNumber(int[] nums) {
        Map<Integer, Integer> freq = new HashMap<>(); // Создаем мапу для подсчета количества вхождений каждого числа
        for (int num : nums) { // Заполняем мапу: ключ - число, значение - количество вхождений
            freq.put(num, freq.getOrDefault(num, 0) + 1);
        }

        int maxNum = -1; // Инициализируем максимальное значение как -1

        for (Map.Entry<Integer, Integer> entry : freq.entrySet()) { // Проходим по всем записям в мапе
            int num = entry.getKey();
            int count = entry.getValue();
            // Если число встречается 1 раз и больше текущего максимума, обновляем максимум
            if (count == 1 && num > maxNum) {
                maxNum = num;
            }
        }

        return maxNum; // Возвращаем максимальное уникальное число
    }
}
