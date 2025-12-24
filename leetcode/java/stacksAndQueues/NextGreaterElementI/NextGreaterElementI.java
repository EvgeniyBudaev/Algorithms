package stacksAndQueues.NextGreaterElementI;

import java.util.ArrayDeque;
import java.util.Arrays;
import java.util.Deque;
import java.util.HashMap;
import java.util.Map;

/* 496. Next Greater Element I
https://leetcode.com/problems/next-greater-element-i/description/

Следующий больший элемент некоторого элемента x в массиве — это первый больший элемент, находящийся справа от x в том же
массиве.
Вам даны два различных целочисленных массива с нулевым индексом nums1 и nums2, где nums1 — это подмножество nums2.
Для каждого 0 <= i < nums1.length найдите индекс j такой, что nums1[i] == nums2[j] и определите следующий больший
элемент nums2[j] в nums2. Если следующего большего элемента нет, то ответом на этот запрос будет -1.
Верните массив ans длины nums1.length такой, что ans[i] является следующим большим элементом, как описано выше.

Input: nums1 = [4,1,2], nums2 = [1,3,4,2]
Output: [-1,3,-1]
Объяснение: Следующий больший элемент для каждого значения nums1 выглядит следующим образом:
- 4 подчеркнуто в nums2 = [1,3,4,2]. Следующего большего элемента нет, поэтому ответ - -1.
- 1 подчеркнута в nums2 = [1,3,4,2]. Следующий больший элемент — 3.
- 2 подчеркнуто в nums2 = [1,3,4,2]. Следующего большего элемента нет, поэтому ответ - -1.

Input: nums1 = [2,4], nums2 = [1,2,3,4]
Output: [3,-1]
Объяснение: Следующий больший элемент для каждого значения nums1 выглядит следующим образом:
- 2 подчеркнуто в nums2 = [1,2,3,4]. Следующий больший элемент — 3.
- 4 подчеркнуто в nums2 = [1,2,3,4]. Следующего большего элемента нет, поэтому ответ - -1.
*/

public class NextGreaterElementI {
    public static void main(String[] args) {
        int[] nums1 = {4, 1, 2};
        int[] nums2 = {1, 3, 4, 2};
        int[] result = nextGreaterElement(nums1, nums2);
        System.out.println(Arrays.toString(result)); // [-1, 3, -1]
    }

    // nextGreaterElement находит следующий больший элемент для каждого элемента в nums1 из nums2.
    // time: O(n), space: O(n), где n — количество элементов в nums2.
    private static int[] nextGreaterElement(int[] nums1, int[] nums2) {
        Deque<Integer> stack = new ArrayDeque<>(); // Используем стек для хранения элементов
        Map<Integer, Integer> mapping = new HashMap<>(); // Маппинг для следующего большего элемента

        // Проходим по nums2 справа налево
        for (int i = nums2.length - 1; i >= 0; i--) {
            int num = nums2[i]; // Текущий элемент

            // Удаляем из стека все элементы, меньшие или равные текущему
            while (!stack.isEmpty() && stack.peek() <= num) {
                stack.pop();
            }

            // Записываем в маппинг следующий больший элемент или -1
            if (!stack.isEmpty()) {
                mapping.put(num, stack.peek());
            } else {
                mapping.put(num, -1);
            }

            // Добавляем текущий элемент в стек
            stack.push(num);
        }

        // Строим результат для nums1
        int[] result = new int[nums1.length];
        for (int i = 0; i < nums1.length; i++) {
            result[i] = mapping.get(nums1[i]);
        }

        return result;
    }
}
