package slidingWindow.LongestRepeatingCharacterReplacement;

import java.util.HashMap;
import java.util.Map;

/* 424. Longest Repeating Character Replacement
https://leetcode.com/problems/longest-repeating-character-replacement/description/

Вам дана строка s и целое число k. Вы можете выбрать любой символ строки и заменить его на любой другой символ
английского верхнего регистра. Эту операцию можно выполнить не более k раз.

Верните длину самой длинной подстроки, содержащей ту же букву, которую вы можете получить после выполнения
вышеуказанных операций.

Input: s = "ABAB", k = 2
Output: 4
Пояснение: Замените две буквы «А» двумя буквами «B» или наоборот.

Input: s = "AABABBA", k = 1
Output: 4
Пояснение: Замените букву «A» посередине на «B» и получите «AABBBBA».
Подстрока «BBBB» имеет самую длинную повторяющуюся букву — 4.
Могут существовать и другие способы получить этот ответ.
*/

public class LongestRepeatingCharacterReplacement {
    public static void main(String[] args) {
        System.out.println(characterReplacement("ABAB", 2)); // 4
    }

    // characterReplacement находит длину самой длинной подстроки, в которой можно сделать не более k замен символов.
    // time: O(n), space: O(1)
    private static int characterReplacement(String s, int k) {
        Map<Character, Integer> charCount = new HashMap<>(); // Мапа для подсчета количества каждого символа в подстроке
        int left = 0, result = 0, maxFreq = 0; // Левый индекс подстроки, результат и максимальная частота в подстроке

        for (int right = 0; right < s.length(); right++) { // Проходим по строке
            char current = s.charAt(right); // Текущий символ
            charCount.put(current, charCount.getOrDefault(current, 0) + 1); // Увеличиваем счетчик текущего символа
            maxFreq = Math.max(maxFreq, charCount.get(current)); // Обновляем максимальную частоту
            int windowSize = right - left + 1; // Текущая длина подстроки

            // Если текущая подстрока не подходит по условию
            if (windowSize - maxFreq > k) {
                char leftChar = s.charAt(left);
                charCount.put(leftChar, charCount.get(leftChar) - 1); // Уменьшаем счетчик левого символа
                left++; // Сдвигаем левый индекс
            }

            // Обновляем результат текущей длиной подстроки
            result = Math.max(result, right - left + 1);
        }

        return result; // Возвращаем результат
    }
}
