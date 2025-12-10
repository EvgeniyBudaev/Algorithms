package slidingWindow.MinimumWindowSubstring;

import java.util.HashMap;
import java.util.Map;

/* 76. Minimum Window Substring
https://leetcode.com/problems/minimum-window-substring/description/
solution https://leetcode.com/problems/minimum-window-substring/solutions/4673781/easy-sliding-window-solution-explained/

Учитывая две строки s и t длиной m и n соответственно, верните минимальную подстроку окна s, такую, что каждый символ
в t (включая дубликаты) включен в окно. Если такой подстроки нет, верните пустую строку «».

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Объяснение: Подстрока минимального окна "BANC" включает "A", "B" и "C" из строки t.

Input: s = "a", t = "a"
Output: "a"
Объяснение: Вся строка s представляет собой минимальное окно.

Input: s = "a", t = "aa"
Output: ""
Объяснение: В окно должны быть включены обе буквы 'a' из t.
Поскольку самое большое окно s имеет только одну букву «а», верните пустую строку.
*/

public class MinimumWindowSubstring {
    public static void main(String[] args) {
        System.out.println(minWindow("ADOBECODEBANC", "ABC")); // "BANC"
    }

    // minWindow - функция для нахождения минимального окна в строке s, содержащего все символы из строки t.
    // time: O(n), space: O(1)
    private static String minWindow(String s, String t) {
        // Если длина t больше s, то нет смысла искать
        if (s.length() < t.length()) {
            return "";
        }

        // Создаем карту для подсчета символов в t
        Map<Character, Integer> charMap = new HashMap<>();
        for (char c : t.toCharArray()) {
            charMap.put(c, charMap.getOrDefault(c, 0) + 1);
        }

        int left = 0;
        int minLen = s.length() + 1; // Инициализируем минимальную длину как максимально возможную
        int ansLeft = 0, ansRight = 0; // Индексы для минимального окна
        int required = t.length(); // Количество символов, которые нужно найти

        for (int right = 0; right < s.length(); right++) { // Перебираем все символы в s
            char current = s.charAt(right); // Получаем текущий символ

            if (charMap.getOrDefault(current, 0) > 0) { // Если текущий символ есть в charMap
                required--; // Уменьшаем количество требуемых символов
            }
            charMap.put(current, charMap.getOrDefault(current, 0) - 1); // Уменьшаем количество текущего символа

            // Когда все символы найдены, пытаемся сузить окно
            while (required == 0) {
                // Обновляем минимальное окно
                if (right - left + 1 < minLen) { // Если новое окно короче предыдущего
                    minLen = right - left + 1; // Обновляем минимальную длину
                    ansLeft = left; // Обновляем индексы
                    ansRight = right; // Обновляем индексы
                }

                // Убираем левый символ из окна
                char leftChar = s.charAt(left);
                charMap.put(leftChar, charMap.get(leftChar) + 1); // Увеличиваем количество требуемых символов
                if (charMap.get(leftChar) > 0) { // Если символ был частью t, уменьшаем
                    required++; // Если символ был частью t, увеличиваем required
                }
                left++; // Сдвигаем левый указатель
            }
        }

        // Если минимальное окно не найдено, возвращаем пустую строку
        if (minLen > s.length()) {
            return "";
        }

        return s.substring(ansLeft, ansRight + 1); // Возвращаем минимальное окно
    }
}
