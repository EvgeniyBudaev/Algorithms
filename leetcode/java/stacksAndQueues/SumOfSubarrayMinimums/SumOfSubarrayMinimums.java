package stacksAndQueues.SumOfSubarrayMinimums;

import java.util.ArrayList;
import java.util.List;

/* 907. Sum of Subarray Minimums
https://leetcode.com/problems/sum-of-subarray-minimums/description/

Дан массив целых чисел arr, найти сумму min(b), где b пробегает каждый (смежный) подмассив arr.
Поскольку ответ может быть большим, верните ответ по модулю 109 + 7.

Input: arr = [3,1,2,4]
Output: 17
Пояснение:
Подмассивы — это [3], [1], [2], [4], [3,1], [1,2], [2,4], [3,1,2], [1,2,4], [3,1,2,4].
Минимумы — 3, 1, 2, 4, 1, 1, 2, 1, 1, 1.
Сумма — 17.
*/

public class SumOfSubarrayMinimums {
    public static void main(String[] args) {
        int[] arr = {3, 1, 2, 4};
        System.out.println(sumSubarrayMins(arr)); // 17
    }

    // sumSubarrayMins находит сумму минимумов каждого подмассива.
    // time: O(n), space: O(n)
    private static int sumSubarrayMins(int[] arr) {
        int n = arr.length;
        long[] sums = new long[n]; // Суммы каждого подмассива
        List<Integer> stack = new ArrayList<>(); // Стэк индексов

        for (int i = 0; i < n; i++) {
            // Сравниваем текущее число с верхним индексом стэка
            // Если текущее число меньше верхнего индекса стэка, удаляем верхний индекс стэка
            while (!stack.isEmpty() && arr[i] < arr[stack.get(stack.size() - 1)]) {
                stack.remove(stack.size() - 1);
            }

            stack.add(i); // Добавляем текущий индекс стэку

            if (stack.size() == 1) { // Если стэк пустой, то сумма подмассива равна текущему числу
                sums[i] = (long) (i + 1) * arr[i]; // Сумма подмассива равна текущему числу
            } else { // Если стэк не пустой, то сумма подмассива равна сумме подмассива с верхним индексом стэка и текущему числу
                int prev = stack.get(stack.size() - 2); // Предыдущий индекс стэка
                sums[i] = (long) (i - prev) * arr[i] + sums[prev]; // Сумма подмассива равна сумме подмассива с верхним индексом стэка и текущему числу
            }
        }

        long result = 0; // Результат
        for (long sum : sums) {
            result += sum; // Суммируем все суммы подмассивов
        }

        // Модуль для избежания переполнения
        return (int) (result % 1_000_000_007);
    }
}
