package slidingWindow.LongestSubstringWithAtMostKDistinctCharacters;

import java.util.HashMap;
import java.util.Map;

/* https://github.com/EvgeniyBudaev/doocs_leetcode/tree/main/solution/0100-0199/0159.Longest%20Substring%20with%20At%20Most%20Two%20Distinct%20Characters

Учитывая строку s и целое число k, верните длину самой длинной подстроки s, содержащей не более k различных символов.

Input: s = "eceba", k = 2
Output: 3
Объяснение: Это подстрока «ece» длиной 3.

Input: s = "aa", k = 1
Output: 2
Объяснение: Это подстрока «aa» длиной 2.
*/


public class LongestSubstringWithAtMostKDistinctCharacters {
    public static void main(String[] args) {
        System.out.println(lengthOfLongestSubstringKDistinct("eceba", 2)); // Вывод: 3
    }

    // lengthOfLongestSubstringKDistinct возвращает длину самой длинной подстроки s, содержащей не более k различных символов.
    // time: O(n), space: O(k)
    private  static int lengthOfLongestSubstringKDistinct(String s, int k) {
        if (s == null || k < 0) {
            return 0;
        }
        if (k == 0) {
            return 0;
        }

        Map<Character, Integer> charMap = new HashMap<>(); // Хранит частоту символов в текущем окне
        int maxLength = 0; // Максимальная длина подстроки
        int left = 0; // Указатель на начало окна

        for (int right = 0; right < s.length(); right++) {
            char currentChar = s.charAt(right); // Текущий символ
            charMap.put(currentChar, charMap.getOrDefault(currentChar, 0) + 1); // Увеличиваем частоту текущего символа

            // Если количество уникальных символов превышает k, сдвигаем левый указатель
            while (charMap.size() > k) {
                char leftChar = s.charAt(left);
                charMap.put(leftChar, charMap.get(leftChar) - 1); // Уменьшаем частоту символа слева
                if (charMap.get(leftChar) == 0) {
                    charMap.remove(leftChar); // Удаляем символ, если его частота стала 0
                }
                left++; // Сдвигаем left
            }

            maxLength = Math.max(maxLength, right - left + 1); // Обновляем максимальную длину подстроки
        }

        return maxLength; // Возвращаем максимальную длину подстроки
    }
}
