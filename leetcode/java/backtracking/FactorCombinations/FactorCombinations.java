package leetcode.java.backtracking.FactorCombinations;

import java.util.ArrayList;
import java.util.List;

/* 254. Factor Combinations

Числа можно рассматривать как произведение своих множителей. Например,
8 = 2 х 2 х 2;
  = 2 х 4.
Напишите функцию, которая принимает целое число n и возвращает все возможные комбинации его множителей.

Input: 1
Output: []

Input: 37
Output:[]

Input: 12
Output: [[2, 6], 2, 2, 3], [3, 4]]
*/


public class FactorCombinations {
    public static void main(String[] args) {
        System.out.println(getFactors(12)); // [[2, 2, 3], [2, 6], [3, 4]]
    }

    /**
     * Находит все возможные комбинации множителей числа n,
     * где каждый множитель больше или равен предыдущему.
     * Исключает тривиальный случай [n].
     * <p>
     * Time Complexity: O(√n * P(n))
     * - √n — перебор делителей на каждом уровне рекурсии
     * - P(n) — количество мультипликативных разбиений числа n
     * - В худшем случае экспоненциально от количества простых делителей
     * Space Complexity: O(log n) — не считая выходных данных
     * - O(log n) — глубина рекурсивного стека (макс. при делении на 2)
     * - O(log n) — хранение текущей комбинации в процессе рекурсии
     * - O(P(n) * log n) — хранение результата (выходные данные, не учитывается)
     */
    private static List<List<Integer>> getFactors(int n) {
        List<List<Integer>> result = new ArrayList<>();
        backtrack(n, 2, new ArrayList<>(), result);
        return result;
    }

    /**
     * Вспомогательная рекурсивная функция с обратным ходом (backtracking)
     * для генерации комбинаций множителей.
     *
     * @param remain  оставшееся число для разложения на множители
     * @param start   минимальный множитель для предотвращения дубликатов
     * @param current текущая формируемая комбинация множителей
     * @param result  список для хранения всех найденных валидных комбинаций
     *                <p>
     *                Time: O(√remain * P(remain)) — перебор делителей с рекурсией
     *                Space: O(log n) — глубина рекурсивного стека
     */
    private static void backtrack(int remain, int start, List<Integer> current,
                                  List<List<Integer>> result) {
        // Базовый случай: число полностью разложено на множители
        if (remain == 1) {
            // Исключаем тривиальный случай: комбинация из одного элемента [n]
            if (current.size() > 1) {
                result.add(new ArrayList<>(current)); // добавляем копию комбинации
            }
            return;
        }

        // Перебираем возможные множители от start до √remain
        // Оптимизация: если i > √remain, то remain/i < i, и такой делитель
        // уже был рассмотрен ранее или приведёт к нарушению порядка
        for (int i = start; i * i <= remain; i++) {
            if (remain % i == 0) { // Если i является делителем
                // Выбираем: добавляем множитель в текущую комбинацию
                current.add(i);
                // Исследуем: рекурсивно ищем множители для remain/i,
                // начиная с i (чтобы сохранить неубывающий порядок)
                backtrack(remain / i, i, current, result);
                // Отменяем выбор: удаляем последний множитель (backtracking)
                current.remove(current.size() - 1);
            }
        }

        // Важное дополнение: сам остаток (remain) тоже может быть множителем,
        // если он >= start. Это позволяет добавить комбинации типа [2, 6] для n=12.
        // Добавляем только если remain >= start и current не пуст (чтобы избежать [12])
        if (remain >= start && !current.isEmpty()) {
            current.add(remain);
            result.add(new ArrayList<>(current));
            current.remove(current.size() - 1);
        }
    }
}
