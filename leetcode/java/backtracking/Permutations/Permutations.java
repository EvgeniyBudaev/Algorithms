package leetcode.java.backtracking.Permutations;

import java.util.ArrayList;
import java.util.List;

/* 46. Permutations
https://leetcode.com/problems/permutations/description/

Учитывая массив различных целых чисел, верните все возможные перестановки. Вы можете вернуть ответ в любом порядке.

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Input: nums = [0,1]
Output: [[0,1],[1,0]]

Input: nums = [1]
Output: [[1]]
*/

public class Permutations {
    public static void main(String[] args) {
        System.out.println(permute(new int[]{1, 2, 3})); // [[1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1]]
    }

    // permute возвращает все возможные перестановки чисел из входного массива.
    // time: O(n!) - факториал времени, поскольку для каждого числа существует (n-1)! перестановок.
    // space: O(n!) - факториал памяти, поскольку для каждого числа существует (n-1)! перестановок.
    private static List<List<Integer>> permute(int[] nums) {
        List<List<Integer>> result = new ArrayList<>(); // Список для хранения результирующих перестановок
        int n = nums.length; // Длина входного массива

        backtrack(nums, new ArrayList<>(), result, n);
        return result;
    }

    // Рекурсивная функция для генерации перестановок
    private static void backtrack(int[] nums, List<Integer> current, List<List<Integer>> result, int n) {
        // Если текущая перестановка завершена, добавляем в результат
        if (current.size() == n) {
            // Создаем копию текущей перестановки и добавляем в результат
            // Важно: в Java списки передаются по ссылке, поэтому создаём новую копию
            result.add(new ArrayList<>(current));
            return;
        }

        // Перебираем все числа, которые еще не использованы
        for (int num : nums) {
            // Если число не использовалось, добавляем его в текущую перестановку
            if (!current.contains(num)) {
                // Добавляем число в текущую перестановку
                current.add(num);
                // Рекурсивно продолжаем построение перестановки
                backtrack(nums, current, result, n);
                // Удаляем последнее число (backtracking)
                current.remove(current.size() - 1);
            }
        }
    }
}
