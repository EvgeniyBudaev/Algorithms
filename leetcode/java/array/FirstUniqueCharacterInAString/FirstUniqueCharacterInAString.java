package array.FirstUniqueCharacterInAString;

import java.util.HashMap;
import java.util.Map;

/* 387. First Unique Character in a String
https://leetcode.com/problems/first-unique-character-in-a-string/description/

Дана строка s, найти в ней первый неповторяющийся символ и вернуть его индекс. Если он не существует, вернуть -1.

Input: s = "leetcode"
Output: 0
Пояснение: Символ «l» в индексе 0 — это первый символ, который не встречается ни в каком другом индексе.
*/

public class FirstUniqueCharacterInAString {
    public static void main(String[] args) {
        System.out.println(firstUniqChar("leetcode")); // 0
    }

    // firstUniqChar возвращает индекс первого уникального символа в строке s.
    // time: O(n), space: O(n)
    private static int firstUniqChar(String s) {
        // Создаем мап для подсчета количества вхождений каждого символа
        Map<Character, Integer> charCount = new HashMap<>();

        // Первый проход: подсчитываем количество каждого символа
        for (char c : s.toCharArray()) {
            charCount.put(c, charCount.getOrDefault(c, 0) + 1);
        }

        // Второй проход: ищем первый символ с количеством 1
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (charCount.get(c) == 1) {
                return i;
            }
        }

        return -1; // Если не нашли уникальных символов
    }
}
