package slidingWindow.PermutationInString;

import java.util.HashMap;
import java.util.Map;

/* 567. Permutation in String
https://leetcode.com/problems/permutation-in-string/description/

Учитывая две строки s1 и s2, верните true, если s2 содержит перестановку s1, или false в противном случае.
Другими словами, верните true, если одна из перестановок s1 является подстрокой s2.

Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Пояснение: s2 содержит одну перестановку s1 («ba»).

Input: s1 = "ab", s2 = "eidboaoo"
Output: false
*/

public class PermutationInString {
    public static void main(String[] args) {
        System.out.println(checkInclusion("ab", "eidbaooo")); // true
        System.out.println(checkInclusion("ab", "eidboaoo")); // false
    }

    // checkInclusion проверяет, содержит ли строка s2 перестановку строки s1.
    // time: O(n), space: O(1)
    private static boolean checkInclusion(String s1, String s2) {
        // Если длина s1 больше длины s2, то s1 не может быть перестановкой s2
        if (s1.length() > s2.length()) {
            return false;
        }

        // Создаем карту для подсчета символов в s1
        Map<Character, Integer> neededChar = new HashMap<>();
        for (char c : s1.toCharArray()) {
            neededChar.put(c, neededChar.getOrDefault(c, 0) + 1);
        }

        int left = 0, right = 0; // Инициализируем левую и правую границы окна
        int requiredLength = s1.length(); // Инициализируем счетчик необходимых символов

        while (right < s2.length()) { // Проходим по s2
            char currentChar = s2.charAt(right); // Текущий символ в s2

            // Если текущий символ найден в s1, уменьшаем счетчик необходимых символов
            if (neededChar.getOrDefault(currentChar, 0) > 0) {
                requiredLength--; // Уменьшаем счетчик необходимых символов
            }
            neededChar.put(currentChar, neededChar.getOrDefault(currentChar, 0) - 1); // Уменьшаем счетчик текущего символа
            right++; // Сдвигаем правую границу окна

            // Если все символы s1 найдены в текущем окне
            if (requiredLength == 0) {
                return true;
            }

            // Если окно достигло длины s1, сдвигаем левую границу
            if (right - left == s1.length()) {
                char leftChar = s2.charAt(left);
                // Если символ на левой границе был найден в s1, увеличиваем счетчик необходимых символов
                if (neededChar.get(leftChar) >= 0) {
                    requiredLength++; // Увеличиваем счетчик необходимых символов
                }
                neededChar.put(leftChar, neededChar.get(leftChar) + 1); // Восстанавливаем счетчик текущего символа
                left++; // Сдвигаем левую границу окна
            }
        }

        return false; // Если перестановка не найдена
    }
}
