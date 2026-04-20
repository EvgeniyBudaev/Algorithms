package leetcode.java.backtracking.NumbersWithSameConsecutiveDifferences;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/* 967. Numbers With Same Consecutive Differences
https://leetcode.com/problems/numbers-with-same-consecutive-differences/description/

Учитывая два целых числа n и k, верните массив всех целых чисел длины n, где разница между всеми двумя последовательными
цифрами равна k. Вы можете вернуть ответ в любом порядке.
Обратите внимание, что целые числа не должны иметь ведущих нулей. Целые числа 02 и 043 не допускаются.

Input: n = 3, k = 7
Output: [181,292,707,818,929]
Объяснение: Обратите внимание: 070 не является допустимым числом, поскольку в нем есть ведущие нули.

Input: n = 2, k = 1
Output: [10,12,21,23,32,34,43,45,54,56,65,67,76,78,87,89,98]
*/

public class NumbersWithSameConsecutiveDifferences {
    public static void main(String[] args) {
        System.out.println(Arrays.toString(numsSameConsecDiff(3, 7))); // [181, 292, 707, 818, 929]
        System.out.println(Arrays.toString(numsSameConsecDiff(2, 1))); // [10, 12, 21, 23, 32, ...]
        System.out.println(Arrays.toString(numsSameConsecDiff(1, 0))); // [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
    }

    /**
     * Генерирует все n-значные числа, где абсолютная разница между соседними цифрами равна k.
     * <p>
     * Алгоритм: итеративное построение чисел «слоями» (BFS-подобный подход).
     * - Начинаем с однозначных чисел 1–9 (первая цифра не может быть 0)
     * - На каждом шаге расширяем каждое число, добавляя цифру, отличающуюся на ±k от последней
     * - Повторяем, пока не достигнем длины n
     *
     * @param n требуемая длина числа (количество цифр)
     * @param k абсолютная разница между соседними цифрами
     * @return массив всех валидных чисел
     * <p>
     * Time: O(2^n) — в худшем случае каждая цифра порождает до 2 продолжений.
     * Общее количество чисел: ≤ 9 * 2^(n-1).
     * <p>
     * Space: O(2^n) — память для хранения текущего и следующего «слоя» чисел,
     * не считая выходной массив. С учётом результата: O(2^n).
     */
    private static int[] numsSameConsecDiff(int n, int k) {
        // Начинаем с цифр 1–9 (первая цифра числа не может быть 0)
        List<Integer> current = new ArrayList<>(List.of(1, 2, 3, 4, 5, 6, 7, 8, 9));

        // Построение чисел по одной цифре за итерацию (нужно добавить n-1 цифр)
        for (int length = 1; length < n; length++) {
            List<Integer> next = new ArrayList<>();

            // Для каждого числа текущего слоя пытаемся добавить следующую цифру
            for (int num : current) {
                int lastDigit = num % 10; // Получаем последнюю цифру

                // Вариант 1: добавляем цифру lastDigit + k (если в диапазоне 0–9)
                int newDigitPlus = lastDigit + k;
                if (newDigitPlus < 10) {
                    next.add(num * 10 + newDigitPlus);
                }

                // Вариант 2: добавляем цифру lastDigit - k (если k != 0 и цифра валидна)
                // Проверка k != 0 нужна, чтобы не дублировать один и тот же переход
                if (k != 0) {
                    int newDigitMinus = lastDigit - k;
                    if (newDigitMinus >= 0) {
                        next.add(num * 10 + newDigitMinus);
                    }
                }
            }

            // Переходим к следующему слою
            current = next;
        }

        // Конвертируем List<Integer> в int[]
        return current.stream().mapToInt(Integer::intValue).toArray();
    }
}
