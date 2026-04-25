package leetcode.java.backtracking.Subsets;

import java.util.ArrayList;
import java.util.List;

/* 78. Subsets
https://leetcode.com/problems/subsets/description/

Учитывая целочисленный массив чисел уникальных элементов, верните все возможные подмножества (набор мощности).
Набор решений не должен содержать повторяющихся подмножеств. Верните решение в любом порядке.

Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Input: nums = [0]
Output: [[],[0]]
*/

public class Subsets {
    public static void main(String[] args) {
        // [[], [1], [2], [1, 2], [3], [1, 3], [2, 3], [1, 2, 3]]  
        System.out.println(subsets(new int[]{1, 2, 3}));
    }

    // subsets возвращает все возможные подмножества (powerset) входного массива.
    // time: O(2^n), где n - количество элементов в массиве.
    // space: O(2^n), так как результат содержит все возможные подмножества.
    private static List<List<Integer>> subsets(int[] nums) {
        List<List<Integer>> result = new ArrayList<>();
        dfs(nums, 0, new ArrayList<>(), result);
        return result;
    }

    /**
     * Рекурсивная вспомогательная функция для генерации подмножеств.
     *
     * @param nums    исходный массив
     * @param index   текущий индекс в массиве
     * @param current текущее формируемое подмножество
     * @param result  итоговый список всех подмножеств
     */
    private static void dfs(int[] nums, int index, List<Integer> current, List<List<Integer>> result) {
        // Базовый случай: достигли конца массива
        if (index == nums.length) {
            // Добавляем копию текущего подмножества в результат
            result.add(new ArrayList<>(current));
            return;
        }

        // Вариант 1: включаем текущий элемент
        current.add(nums[index]);
        dfs(nums, index + 1, current, result);

        // Откат (backtrack): удаляем последний добавленный элемент
        current.remove(current.size() - 1);

        // Вариант 2: не включаем текущий элемент
        dfs(nums, index + 1, current, result);
    }
}
