package slidingWindow.SlidingWindowMaximum;

import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Deque;
import java.util.List;

/* 239. Sliding Window Maximum
https://leetcode.com/problems/sliding-window-maximum/description/
https://leetcode.com/problems/minimum-window-substring/solutions/3061723/js-fastest-easy-commented/

Вам дан массив целых чисел nums, есть скользящее окно размера k, которое перемещается из самого левого края массива в
самый правый. В окне вы можете увидеть только k чисел. Каждый раз скользящее окно перемещается вправо на одну позицию.
Верните максимальное число в скользящем окне.

Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]
Explanation:
Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

Input: nums = [1], k = 1
Output: [1]
*/

public class SlidingWindowMaximum {
    public static void main(String[] args) {
        int[] nums = {1, 3, -1, -3, 5, 3, 6, 7};
        System.out.println(Arrays.toString(maxSlidingWindow(nums, 3))); // [3,3,5,5,6,7]
    }

    // maxSlidingWindow - функция для нахождения максимального значения в каждом окне.
    // time: O(n), space: O(k), где n - длина входного массива, k - размер окна.
    private static int[] maxSlidingWindow(int[] nums, int k) {
        if (nums == null || nums.length == 0 || k <= 0) {
            return new int[0];
        }

        List<Integer> result = new ArrayList<>(); // Инициализируем массив для хранения максимальных значений
        Deque<Integer> deque = new ArrayDeque<>(); // Инициализируем дек (двустороннюю очередь) для хранения индексов элементов в текущем окне

        for (int i = 0; i < nums.length; i++) {
            // Удаляем из дека индексы элементов, которые меньше или равны текущему элементу
            while (!deque.isEmpty() && nums[deque.peekLast()] <= nums[i]) {
                deque.pollLast();
            }

           // Удаляем индексы элементов из дека, находящихся за пределами текущего окна
            if (!deque.isEmpty() && deque.peekFirst() <= i - k) {
                deque.pollFirst();
            }

            deque.offerLast(i); // Добавляем текущий индекс в конец

            // Начинаем записывать результат, как только заполнилось первое окно
            if (i >= k - 1) {
                result.add(nums[deque.peekFirst()]);
            }
        }

        // Преобразуем List<Integer> в int[]
        return result.stream().mapToInt(Integer::intValue).toArray();
    }
}
