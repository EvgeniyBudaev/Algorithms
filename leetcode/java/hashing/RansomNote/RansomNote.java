package leetcode.java.hashing.RansomNote;

import java.util.HashMap;
import java.util.Map;

/* 383. Ransom Note
https://leetcode.com/problems/ransom-note/description/

Учитывая две строки ransomNote и magazine, верните true, если ransomNote можно создать с использованием букв из magazine,
и false в противном случае.
Каждое письмо в magazine можно использовать в ransomNote только один раз.

Input: ransomNote = "a", magazine = "b"
Output: false

Input: ransomNote = "aa", magazine = "ab"
Output: false

Input: ransomNote = "aa", magazine = "aab"
Output: true
*/

public class RansomNote {
    public static void main(String[] args) {
        System.out.println(canConstruct("a", "b"));    // false
        System.out.println(canConstruct("aa", "aab")); // true
    }

    // canConstruct проверяет, можно ли построить строку ransomNote из букв строки magazine.
    // time: O(n), space: O(1)
    private static boolean canConstruct(String ransomNote, String magazine) {
        // Создаем map для подсчета количества каждой буквы в magazine
        Map<Character, Integer> letterCounts = new HashMap<>();

        // Заполняем map: ключ - буква, значение - количество таких букв в magazine
        for (char letter : magazine.toCharArray()) {
            letterCounts.put(letter, letterCounts.getOrDefault(letter, 0) + 1);
        }

        // Проверяем каждую букву в ransomNote
        for (char letter : ransomNote.toCharArray()) {
            // Если буквы нет в magazine или ее недостаточно, возвращаем false
            if (!letterCounts.containsKey(letter) || letterCounts.get(letter) == 0) {
                return false;
            }
            // Уменьшаем количество оставшихся букв
            letterCounts.put(letter, letterCounts.get(letter) - 1);
        }

        // Если все буквы ransomNote есть в magazine в достаточном количестве
        return true;
    }
}
