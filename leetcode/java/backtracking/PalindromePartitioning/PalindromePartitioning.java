package leetcode.java.backtracking.PalindromePartitioning;

import java.util.ArrayList;
import java.util.List;

/* 131. Palindrome Partitioning
https://leetcode.com/problems/palindrome-partitioning/description/

Дана строка s, раздел s такой, что каждый подстрока раздела представляет собой палиндром.
Вернуть все возможные палиндромные разбиения s.

Input: s = "aab"
Output: [["a","a","b"],["aa","b"]]

Input: s = "a"
Output: [["a"]]
*/

public class PalindromePartitioning {
    public static void main(String[] args) {
        System.out.println(partition("aab")); // [[a, a, b], [aa, b]]
        System.out.println(partition("a"));   // [[a]]
        System.out.println(partition("ab"));  // [[a, b]]
    }

    /**
     * Возвращает все возможные разбиения строки на палиндромные подстроки.
     * <p>
     * Алгоритм: обратное отслеживание (backtracking).
     * - Перебираем все возможные окончания текущей подстроки, начиная с позиции `start`
     * - Если подстрока [start, end] — палиндром, добавляем её в текущий путь и рекурсивно
     * обрабатываем оставшуюся часть строки
     * - После возврата из рекурсии выполняем «откат»: удаляем последнюю добавленную подстроку
     * - Когда достигнут конец строки — сохраняем копию текущего пути в результат
     *
     * @param s входная строка
     * @return список всех валидных разбиений, где каждое разбиение — список палиндромных подстрок
     * <p>
     * Time: O(n * 2^n), где n — длина строки.
     * - 2^n — верхняя оценка количества возможных разбиений (каждый разрыв — бинарный выбор)
     * - *n — проверка палиндрома + копирование пути в результат
     * <p>
     * Space: O(n) — вспомогательная память (глубина рекурсии + текущий путь),
     * не считая памяти под результат. С учётом вывода: O(n * 2^n).
     */
    private static List<List<String>> partition(String s) {
        List<List<String>> result = new ArrayList<>();
        backtrack(result, new ArrayList<>(), s, 0);
        return result;
    }

    /**
     * Рекурсивная функция обратного отслеживания для генерации палиндромных разбиений.
     * <p>
     * Принцип работы:
     * - На каждом шаге пытаемся выделить палиндромную подстроку, начинающуюся с `start`
     * - Для каждого валидного окончания `end` добавляем подстроку в путь и рекурсивно
     * продолжаем с позиции `end + 1`
     * - После возврата из рекурсии удаляем последнюю добавленную подстроку (откат)
     * - Когда `start == s.length()` — все символы обработаны, сохраняем копию пути
     *
     * @param result список для накопления итоговых разбиений
     * @param path   текущий путь — список палиндромных подстрок, собранных на данном этапе
     * @param s      исходная строка
     * @param start  текущая позиция в строке, с которой начинаем искать следующую подстроку
     */
    private static void backtrack(List<List<String>> result, List<String> path, String s, int start) {
        // Базовый случай: все символы строки обработаны — сохраняем текущее разбиение
        if (start == s.length()) {
            result.add(new ArrayList<>(path)); // Важно: добавляем копию, а не ссылку!
            return;
        }

        // Перебираем все возможные окончания подстроки, начинающейся с `start`
        for (int end = start; end < s.length(); end++) {
            // Если подстрока [start, end] — палиндром, продолжаем рекурсию
            if (isPalindrome(s, start, end)) {
                path.add(s.substring(start, end + 1)); // Выбираем: добавляем подстроку в путь
                backtrack(result, path, s, end + 1);   // Исследуем: рекурсия к оставшейся части
                path.remove(path.size() - 1);          // Откат: удаляем последнюю подстроку
            }
        }
    }

    /**
     * Проверяет, является ли подстрока s[left..right] палиндромом.
     * <p>
     * Использует двухуказательный подход: сравниваем символы с концов к центру.
     *
     * @param s     исходная строка
     * @param left  левая граница подстроки (включительно)
     * @param right правая граница подстроки (включительно)
     * @return true если подстрока — палиндром, false иначе
     * <p>
     * Time: O(m), где m = right - left + 1 — длина проверяемой подстроки
     */
    private static boolean isPalindrome(String s, int left, int right) {
        while (left < right) {
            if (s.charAt(left) != s.charAt(right)) {
                return false;
            }
            left++;
            right--;
        }
        return true;
    }
}
