package stacksAndQueues.DailyTemperatures;

import java.util.ArrayDeque;
import java.util.Arrays;
import java.util.Deque;

/* 739. Daily Temperatures
https://leetcode.com/problems/daily-temperatures/description/

Учитывая, что массив целых чисел представляет собой ежедневную температуру, верните ответ массива, такой,
что ответ [i] — это количество дней, которое вам нужно подождать после i-го дня, чтобы получить более высокую
температуру. Если нет будущего дня, для которого это возможно, вместо этого сохраните ответ [i] == 0.

Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]

Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]

Input: temperatures = [30,60,90]
Output: [1,1,0]
*/

public class DailyTemperatures {
    public static void main(String[] args) {
        int[] temps = {73, 74, 75, 71, 69, 72, 76, 73};
        System.out.println(Arrays.toString(dailyTemperatures(temps))); // [1, 1, 4, 2, 1, 1, 0, 0]
    }

    // dailyTemperatures возвращает количество дней, которое вам нужно подождать после i-го дня,
    // чтобы получить более высокую температуру.
    // time: O(n), space: O(n), где n - длина входного массива температур.
    private static int[] dailyTemperatures(int[] temperatures) {
        Deque<Integer> stack = new ArrayDeque<>(); // Стек для хранения индексов дней с пониженной температурой
        int[] result = new int[temperatures.length];

        for (int i = 0; i < temperatures.length; i++) {
            // Пока стек не пуст и текущая температура выше температуры на вершине стека
            while (!stack.isEmpty() && temperatures[i] > temperatures[stack.peek()]) {
                int idx = stack.pop();            // Запоминаем индекс последнего элемента в стеке и удаляем последний элемент
                result[idx] = i - idx;            // Записываем разницу в днях
            }
            stack.push(i); // Добавляем текущий индекс в стек
        }

        return result; // Возвращаем количество дней
    }
}
