package leetcode.java.backtracking.Combinations;

import java.util.ArrayList;
import java.util.List;

/* 77. Combinations
https://leetcode.com/problems/combinations/description/

Учитывая два целых числа n и k, верните все возможные комбинации k чисел, выбранных из диапазона [1, n].
Вы можете вернуть ответ в любом порядке.

Input: n = 4, k = 2
Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
Пояснение: всего есть 4 комбинации на выбор 2 = 6.
Обратите внимание, что комбинации неупорядочены, т. е. [1,2] и [2,1] считаются одной и той же комбинацией.
*/

public class Combinations {
    public static void main(String[] args) {
        System.out.println(combine(4, 2)); // [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
    }

    // combine генерирует все возможные комбинации k чисел из диапазона [1, n]
    // time: O(C(n,k) * k), space: O(C(n,k) * k) for output + O(k) for recursion stack
    private static List<List<Integer>> combine(int n, int k) {
        List<List<Integer>> result = new ArrayList<>();
        backtrack(n, k, 1, new ArrayList<>(), result);
        return result;
    }

    private static void backtrack(int n, int k, int start, List<Integer> current, List<List<Integer>> result) {
        // Если текущая комбинация достигла нужной длины, добавляем копию в результат
        if (current.size() == k) {
            result.add(new ArrayList<>(current));
            return;
        }

        // Генерируем комбинации, начиная с числа start
        for (int i = start; i <= n; i++) {
            current.add(i);                           // Добавляем текущее число
            backtrack(n, k, i + 1, current, result);  // Рекурсивный вызов для следующих чисел
            current.remove(current.size() - 1);       // Backtracking: удаляем последнее число
        }
    }
}
