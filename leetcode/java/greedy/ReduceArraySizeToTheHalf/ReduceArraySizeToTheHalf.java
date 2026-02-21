package greedy.ReduceArraySizeToTheHalf;

import java.util.ArrayList;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/* 1338. Reduce Array Size to The Half
https://leetcode.com/problems/reduce-array-size-to-the-half/description/

Вам дан целочисленный массив. Вы можете выбрать набор целых чисел и удалить все вхождения этих целых чисел в массиве.
Верните минимальный размер набора, чтобы была удалена хотя бы половина целых чисел массива.

Input: arr = [3,3,3,3,5,5,5,2,2,7]
Output: 2
Объяснение: При выборе {3,7} будет создан новый массив [5,5,5,2,2] размером 5
(т.е. равный половине размера старого массива).
Возможные наборы размера 2: {3,5},{3,2},{5,2}.
Выбор набора {2,7} невозможен, так как при этом будет создан новый массив [3,3,3,3,5,5,5],
размер которого превышает половину размера старого массива.

Input: arr = [7,7,7,7,7,7]
Output: 1
Пояснение: Единственный возможный набор, который вы можете выбрать, — это {7}. Это сделает новый массив пустым.
*/

public class ReduceArraySizeToTheHalf {
    public static void main(String[] args) {
        int[] arr = {3, 3, 3, 3, 5, 5, 5, 2, 2, 7};
        System.out.println(minSetSize(arr)); // 2
    }

    // minSetSize возвращает минимальный размер набора, чтобы была удалена хотя бы половина целых чисел массива.
    // time O(n), space O(n)
    private static int minSetSize(int[] arr) {
        Map<Integer, Integer> freqMap = new HashMap<>(); // Создаем карту для подсчета частот
        // Подсчитываем частоту каждого числа
        for (int num : arr) {
            freqMap.put(num, freqMap.getOrDefault(num, 0) + 1);
        }
        // freqMap: {2=2, 3=4, 5=3, 7=1}

        List<Integer> freqs = new ArrayList<>(freqMap.values()); // Извлекаем частоты в список
        // freqs: [2, 4, 3, 1]
  
        Collections.sort(freqs, Collections.reverseOrder()); // Сортируем частоты по убыванию
        // freqs: [4, 3, 2, 1]

        int total = arr.length; // Общее количество элементов
        int half = total / 2;   // Половина общего количества чисел
        int sum = 0;            // Сумма частот
        int setSize = 0;        // Размер набора

        // Суммируем частоты, пока не достигнем половины
        for (int count : freqs) {
            sum += count; // Добавляем частоту числа к сумме
            setSize++; // Увеличиваем размер набора
            if (sum >= half) { // Если сумма частот больше или равна половине, возвращаем размер набора
                return setSize;
            }
        }

        return setSize; // Возвращаем минимальный размер набора
    }
}
