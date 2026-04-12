package leetcode.java.backtracking.CombinationSum;

import java.util.ArrayList;
import java.util.List;

/* 39. Combination Sum
https://leetcode.com/problems/combination-sum/description/

Учитывая массив различных целых чисел-кандидатов и целевую целочисленную цель, верните список всех уникальных комбинаций
кандидатов, в которых сумма выбранных чисел равна целевому значению. Вы можете возвращать комбинации в любом порядке.

Одно и то же число может быть выбрано из кандидатов неограниченное количество раз. Две комбинации уникальны, если
частота хотя бы одно из выбранных чисел отличается.

Тестовые примеры генерируются таким образом, чтобы количество уникальных комбинаций, дающих целевое значение, составляло
менее 150 комбинаций для данного входного сигнала.

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Объяснение:
2 и 3 являются кандидатами, а 2 + 2 + 3 = 7. Обратите внимание, что 2 можно использовать несколько раз.
7 — кандидат, а 7 = 7.
Это единственные две комбинации.
*/

public class CombinationSum {
    public static void main(String[] args) {
        int[] candidates = {2, 3, 6, 7};
        System.out.println(combinationSum(candidates, 7));
    }

    // combinationSum находит все уникальные комбинации чисел из candidates,
    // которые в сумме дают target. Одно и то же число может использоваться многократно.
    // time: O(n^target) в худшем случае, где n — количество уникальных чисел в candidates.
    // space: O(target) для хранения текущей комбинации и рекурсивных вызовов.
    private static List<List<Integer>> combinationSum(int[] candidates, int target) {
        List<List<Integer>> result = new ArrayList<>();
        backtrack(candidates, target, 0, new ArrayList<>(), 0, result);
        return result;
    }

    /**
     * Внутренняя рекурсивная функция для генерации комбинаций
     *
     * @param candidates массив кандидатов
     * @param target     целевая сумма
     * @param start      индекс начала перебора
     * @param current    текущая комбинация
     * @param sum        текущая сумма
     * @param result     список для хранения результатов
     */
    private static void backtrack(int[] candidates, int target, int start,
                                  List<Integer> current, int sum,
                                  List<List<Integer>> result) {
        // Если сумма совпадает с целевой, добавляем комбинацию в результат
        if (sum == target) {
            result.add(new ArrayList<>(current)); // Создаём копию комбинации
            return;
        }

        // Если сумма превысила целевое значение или дошли до конца массива
        if (sum > target || start >= candidates.length) {
            return;
        }

        // Включаем текущий элемент в комбинацию (можно использовать многократно)
        current.add(candidates[start]);
        backtrack(candidates, target, start, current, sum + candidates[start], result);

        // Исключаем текущий элемент (backtracking)
        current.remove(current.size() - 1);
        backtrack(candidates, target, start + 1, current, sum, result);
    }
}
