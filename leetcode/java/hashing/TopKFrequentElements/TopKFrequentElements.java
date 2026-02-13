package leetcode.java.hashing.TopKFrequentElements;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/* 347. Top K Frequent Elements
https://leetcode.com/problems/top-k-frequent-elements/description/

Учитывая целочисленный массив nums и целое число k, верните k наиболее часто встречающихся элементов. Вы можете вернуть
ответ в любом порядке.

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

Input: nums = [1], k = 1
Output: [1]
*/

public class TopKFrequentElements {
    public static void main(String[] args) {
        int[] nums = {1, 1, 1, 2, 2, 3};
        System.out.println(Arrays.toString(topKFrequent(nums, 2))); // [1, 2]
    }

    // topKFrequent возвращает k наиболее часто встречающихся чисел.
    // time: O(n log n), space: O(n)
    private static int[] topKFrequent(int[] nums, int k) {
        // Создаем map для подсчета частоты каждого числа
        Map<Integer, Integer> frequency = new HashMap<>();

        // Заполняем map: ключ - число, значение - частота встречаемости
        for (int num : nums) {
            frequency.put(num, frequency.getOrDefault(num, 0) + 1);
        }

        // Преобразуем map в список объектов для сортировки
        List<NumFreq> numFreqs = new ArrayList<>();
        for (Map.Entry<Integer, Integer> entry : frequency.entrySet()) {
            numFreqs.add(new NumFreq(entry.getKey(), entry.getValue()));
        }

        // Сортируем по убыванию частоты
        numFreqs.sort((a, b) -> b.freq - a.freq);
        // Результат: [{1, 3}, {2, 2}, {3, 1}]

        // Выбираем k самых частых чисел
        int[] result = new int[Math.min(k, numFreqs.size())];
        for (int i = 0; i < k && i < numFreqs.size(); i++) {
            result[i] = numFreqs.get(i).num;
        }

        return result; // Возвращаем k самых частых чисел
    }

    // // Создаем структуру для хранения чисел и их частот
    private static class NumFreq {
        int num;  // Число
        int freq; // Частота встречаемости
        
        NumFreq(int num, int freq) {
            this.num = num;
            this.freq = freq;
        }
    }
}
