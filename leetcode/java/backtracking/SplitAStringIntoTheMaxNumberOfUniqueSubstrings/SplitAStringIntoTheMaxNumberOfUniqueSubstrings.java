package leetcode.java.backtracking.SplitAStringIntoTheMaxNumberOfUniqueSubstrings;

import java.util.HashSet;
import java.util.Set;

/* 1593. Split a String Into the Max Number of Unique Substrings
https://leetcode.com/problems/split-a-string-into-the-max-number-of-unique-substrings/description/

Учитывая строку s, верните максимальное количество уникальных подстрок, на которые можно разбить данную строку.
Вы можете разделить строку на любой список непустых подстрок, где объединение подстрок образует исходную строку.
Однако вам необходимо разделить подстроки так, чтобы все они были уникальными.
Подстрока — это непрерывная последовательность символов внутри строки.

Input: s = "ababccc"
Output: 5
Пояснение: Один из способов максимального разделения — это ['a', 'b', 'ab', 'c', 'cc'].
Разделение типа ['a', 'b', 'a', 'b', 'c', 'cc'] недопустимо, поскольку у вас есть 'a' и 'b' несколько раз.

Input: s = "aba"
Output: 2
Пояснение: Один из способов максимального разделения — ['a', 'ba'].

Input: s = "aa"
Output: 1
Объяснение: Дальнейшее разделение строки невозможно.
*/

public class SplitAStringIntoTheMaxNumberOfUniqueSubstrings {
    public static void main(String[] args) {
        System.out.println(maxUniqueSplit("ababccc")); // 5
    }

    /**
     * Возвращает максимальное количество уникальных подстрок,
     * на которые можно разбить строку s.
     * <p>
     * Использует рекурсивный перебор с возвратом (backtracking).
     * <p>
     * Time Complexity: O(2^n * n), где n — длина строки.
     * - В худшем случае перебираются все возможные способы разбиения строки (экспоненциально).
     * - Каждая операция substring() и работа с HashSet требует O(n) времени.
     * <p>
     * Space Complexity: O(n^2)
     * - HashSet хранит до O(n) подстрок длиной до O(n) в текущем пути рекурсии.
     * - Глубина стека рекурсии: O(n).
     *
     * @param s входная строка
     * @return максимальное количество уникальных подстрок в разбиении
     */
    private static int maxUniqueSplit(String s) {
        return getMax(0, s, new HashSet<>());
    }

    /**
     * Вспомогательная рекурсивная функция для поиска максимального разбиения
     * с использованием backtracking.
     * <p>
     * Time Complexity: O(2^n * n) в худшем случае для всего дерева рекурсии.
     * Space Complexity: O(n^2) с учётом хранения подстрок в set и стека вызовов.
     *
     * @param start текущая позиция в строке, с которой начинаем брать подстроку
     * @param s     исходная строка
     * @param set   множество уже использованных уникальных подстрок в текущем пути
     * @return максимальное количество уникальных подстрок для оставшейся части строки
     */
    private static int getMax(int start, String s, Set<String> set) {
        // Базовый случай: если дошли до конца строки, возвращаем количество уникальных подстрок
        if (start == s.length()) {
            return set.size();
        }

        int maxCount = 0;

        // Перебираем все возможные подстроки, начиная с текущей позиции
        for (int i = start + 1; i <= s.length(); i++) {
            String sub = s.substring(start, i);

            // Если подстрока ещё не встречалась в текущем пути
            if (!set.contains(sub)) {
                set.add(sub); // Добавляем подстроку в множество

                // Рекурсивно ищем максимальное разбиение для оставшейся части строки
                int currentCount = getMax(i, s, set);
                maxCount = Math.max(maxCount, currentCount);

                set.remove(sub); // Удаляем подстроку (backtracking)
            }
        }

        return maxCount;
    }
}
