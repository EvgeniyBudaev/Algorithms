package stacksAndQueues.LargestRectangleInHistogram;

import java.util.ArrayDeque;
import java.util.Deque;

/* 84. Largest Rectangle in Histogram
https://leetcode.com/problems/largest-rectangle-in-histogram/description/

Учитывая массив целых чисел, представляющих высоту столбца гистограммы, где ширина каждого столбца равна 1, верните
площадь самого большого прямоугольника в гистограмме.

Input: heights = [2,1,5,6,2,3]
Output: 10
Пояснение: Выше представлена гистограмма, ширина каждого столбца которой равна 1.
Самый большой прямоугольник показан в красной области, его площадь = 10 единиц.
*/

public class LargestRectangleInHistogram {
    public static void main(String[] args) {
        int[] heights = {2, 1, 5, 6, 2, 3};
        System.out.println(largestRectangleArea(heights)); // 10
    }

    // largestRectangleArea вычисляет площадь самого большого прямоугольника, который можно построить на гистограмме.
    // time: O(n), space: O(n),
    private static int largestRectangleArea(int[] heights) {
        Deque<Integer> stack = new ArrayDeque<>(); // Стек для хранения индексов столбцов гистограммы
        int maxArea = 0; // Максимальная площадь прямоугольника
        int index = 0; // Текущий индекс в массиве heights

        // Первый проход по всем столбцам гистограммы
        while (index < heights.length) {
            // Если стек пуст или текущий столбец выше или равен последнему в стеке
            if (stack.isEmpty() || heights[index] >= heights[stack.peek()]) {
                stack.push(index); // Добавляем индекс текущего столбца в стек
                index++;
            } else {
                // Если текущий столбец ниже, извлекаем последний индекс из стека
                int top = stack.pop();
                // Вычисляем высоту прямоугольника - высота извлеченного столбца
                int height = heights[top];
                // Ширина прямоугольника:
                // - если стек пуст, ширина равна текущему индексу (все предыдущие столбцы были выше)
                // - иначе ширина = расстояние между текущим индексом и предыдущим в стеке минус 1
                int width = stack.isEmpty() ? index : index - stack.peek() - 1;
                // Вычисляем площадь и обновляем максимум при необходимости
                int area = height * width;
                maxArea = Math.max(maxArea, area);
            }
        }

        // Второй проход - обрабатываем оставшиеся элементы в стеке
        while (!stack.isEmpty()) {
            int top = stack.pop();
            int height = heights[top];
            // Для оставшихся элементов ширина равна длине всей гистограммы минус позиция предыдущего элемента
            int width = stack.isEmpty() ? index : index - stack.peek() - 1;
            int area = height * width;
            maxArea = Math.max(maxArea, area);
        }

        return maxArea;
    }
}
