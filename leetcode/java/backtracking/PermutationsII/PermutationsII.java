package leetcode.java.backtracking.PermutationsII;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

/* 47. Permutations II
https://leetcode.com/problems/permutations-ii/description/

Учитывая набор чисел nums, который может содержать дубликаты, возвращает все возможные уникальные перестановки в любом
порядке.

Input: nums = [1,1,2]
Output: [[1,1,2], [1,2,1], [2,1,1]]

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/

public class PermutationsII {
    public static void main(String[] args) {
        int[] nums = {1, 1, 2};
        List<List<Integer>> result = permuteUnique(nums);
        System.out.println(result); // [[1, 1, 2], [1, 2, 1], [2, 1, 1]]
    }

    // permuteUnique возвращает все уникальные перестановки чисел с учетом дубликатов
    // time: O(n * n!/(n1!*n2!*...*nk!)), где ni — количество каждого повторяющегося элемента
    // space: O(n) для стека рекурсии + O(n * result_size) для хранения результата
    private static List<List<Integer>> permuteUnique(int[] nums) {
        List<List<Integer>> result = new ArrayList<>();
        backtrack(nums, 0, result);
        return result;
    }

    // time: O(1), space: O(1)
    private static void backtrack(int[] nums, int start, List<List<Integer>> result) {
        if (start == nums.length) {
            // Создаём копию текущей перестановки
            List<Integer> permutation = new ArrayList<>();
            for (int num : nums) {
                permutation.add(num);
            }
            result.add(permutation);
            return;
        }

        // Множество для отслеживания уже использованных чисел на текущей позиции
        Set<Integer> seen = new HashSet<>();

        for (int i = start; i < nums.length; i++) {
            // Пропускаем дубликаты
            if (seen.contains(nums[i])) {
                continue;
            }
            seen.add(nums[i]);

            // Меняем местами элементы
            swap(nums, start, i);
            // Рекурсивно генерируем перестановки для следующей позиции
            backtrack(nums, start + 1, result);
            // Возвращаем элементы на место (backtracking)
            swap(nums, start, i);
        }
    }

    // time: O(1), space: O(1)
    private static void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }
}
