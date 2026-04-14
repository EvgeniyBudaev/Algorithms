package leetcode.java.backtracking.CombinationSumIII;

import java.util.ArrayList;
import java.util.List;

/* 216. Combination Sum III
https://leetcode.com/problems/combination-sum-iii/description/

Найдите все допустимые комбинации k чисел, сумма которых равна n, такие, что выполняются следующие условия:
Используются только цифры от 1 до 9.
Каждое число используется не более одного раза.
Возвращает список всех возможных допустимых комбинаций. Список не должен содержать одну и ту же комбинацию дважды,
комбинации могут возвращаться в любом порядке.

Input: k = 3, n = 7
Output: [[1,2,4]]
Объяснение:
1 + 2 + 4 = 7
Других допустимых комбинаций нет

Input: k = 3, n = 9
Output: [[1,2,6],[1,3,5],[2,3,4]]
Объяснение:
1 + 2 + 6 = 9
1 + 3 + 5 = 9
2 + 3 + 4 = 9
Других допустимых комбинаций нет.
*/

public class CombinationSumIII {
    public static void main(String[] args) {
        System.out.println(combinationSum3(3, 7)); // [[1, 2, 4]]
    }

    // combinationSum3 находит все уникальные комбинации k различных чисел от 1 до 9,
    // сумма которых равна n. Каждое число используется не более одного раза.
    // time: O(C(9, k) * k), space: O(k)
    private static List<List<Integer>> combinationSum3(int k, int n) {
        List<List<Integer>> result = new ArrayList<>();
        backtrack(0, 0, k, n, new ArrayList<>(), result);
        return result;
    }

    /**
     * Вспомогательная рекурсивная функция с обратным ходом (backtracking)
     * для генерации комбинаций чисел от 1 до 9.
     * <p>
     * Time: O(C(9, k) * k) — перебор всех комбинаций с копированием
     * Space: O(k) — глубина рекурсивного стека
     */
    private static void backtrack(int currentDigit, int sum, int k, int n,
                                  List<Integer> elements, List<List<Integer>> result) {
        // Базовые случаи остановки рекурсии
        if (currentDigit > 9 || sum > n || elements.size() > k) {
            return;
        }

        // Если нашли подходящую комбинацию: сумма и количество элементов совпадают
        if (sum == n && elements.size() == k) {
            result.add(new ArrayList<>(elements)); // добавляем копию, чтобы избежать мутаций
            return;
        }

        // Перебираем следующие цифры от currentDigit + 1 до 9
        // Это гарантирует: 1) уникальность чисел 2) отсутствие дубликатов комбинаций
        for (int i = currentDigit + 1; i <= 9; i++) {
            // Выбираем: добавляем цифру в текущую комбинацию
            elements.add(i);
            // Исследуем: рекурсивный вызов с обновлёнными параметрами
            backtrack(i, sum + i, k, n, elements, result);
            // Отменяем выбор: удаляем последнюю цифру (backtracking)
            elements.remove(elements.size() - 1);
        }
    }
}
