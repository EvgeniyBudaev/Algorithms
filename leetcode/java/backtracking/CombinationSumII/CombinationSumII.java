package leetcode.java.backtracking.CombinationSumII;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/* 40. Combination Sum II
https://leetcode.com/problems/combination-sum-ii/description/

Учитывая набор номеров кандидатов (кандидатов) и целевое число (цель), найдите все уникальные комбинации среди
кандидатов, в которых сумма чисел кандидатов равна целевому значению.
Каждое число кандидатов можно использовать в комбинации только один раз.
Примечание. Набор решений не должен содержать повторяющихся комбинаций.

Input: candidates = [10,1,2,7,6,1,5], target = 8
Output: [[1,1,6], [1,2,5], [1,7], [2,6]]
*/

public class CombinationSumII {
    public static void main(String[] args) {
        int[] candidates = {10, 1, 2, 7, 6, 1, 5};
        System.out.println(combinationSum2(candidates, 8)); // [[1, 1, 6], [1, 2, 5], [1, 7], [2, 6]] 
    }

    // combinationSum2 находит все уникальные комбинации чисел из candidates,
    // которые в сумме дают target. Каждое число может использоваться только один раз
    // time: O(2^n * n + n log n), space: O(n)
    private static List<List<Integer>> combinationSum2(int[] candidates, int target) {
        Arrays.sort(candidates);
        List<List<Integer>> result = new ArrayList<>();
        dfs(candidates, target, 0, new ArrayList<>(), result);
        return result;
    }

    /**
     * Вспомогательная рекурсивная функция с обратным ходом (backtracking)
     * <p>
     * time: O(2^n * n) — перебор подмножеств с копированием комбинаций
     * space: O(n) — глубина рекурсивного стека
     */
    private static void dfs(int[] candidates, int remainingTarget, int start,
                            List<Integer> current, List<List<Integer>> result) {
        // Базовый случай: достигли целевой суммы
        if (remainingTarget == 0) {
            result.add(new ArrayList<>(current)); // добавляем копию комбинации
            return;
        }

        // Перебираем кандидатов, начиная с указанной позиции
        for (int i = start; i < candidates.length; i++) {
            // Пропускаем дубликаты на одном уровне рекурсии
            if (i > start && candidates[i] == candidates[i - 1]) {
                continue;
            }
            // Если текущий кандидат превышает оставшуюся цель — прерываем (оптимизация)
            if (candidates[i] > remainingTarget) {
                break;
            }

            // Выбираем: добавляем кандидата в текущую комбинацию
            current.add(candidates[i]);
            // Исследуем: рекурсивный вызов с уменьшенной целью и следующим индексом
            dfs(candidates, remainingTarget - candidates[i], i + 1, current, result);
            // Отменяем выбор: удаляем последний элемент (backtracking)
            current.remove(current.size() - 1);
        }
    }
}
