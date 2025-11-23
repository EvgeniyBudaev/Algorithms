package slidingWindow.FindAllAnagramsInAString;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/* 438. Find All Anagrams in a String
https://leetcode.com/problems/find-all-anagrams-in-a-string/description/

Учитывая две строки s и p, верните массив всех начальных индексов анаграмм p в s. Вы можете вернуть ответ в любом
порядке.
Анаграмма — это слово или фраза, образованная путем перестановки букв другого слова или фразы, обычно с использованием
всех исходных букв ровно один раз.

Input: s = "cbaebabacd", p = "abc"
Output: [0,6]
Объяснение:
Подстрока с начальным индексом = 0 — это «cba», которая является анаграммой слова «abc».
Подстрока с начальным индексом = 6 — это «bac», которая является анаграммой слова «abc».
*/

public class FindAllAnagramsInAString {
    public static void main(String[] args) {
        System.out.println(findAnagrams("cbaebabacd", "abc")); // [0, 6]
    }

    // findAnagrams находит все анаграммы строки p в строке s и возвращает их начальные индексы.
    // time: O(n), space: O(k), где k — размер алфавита (в худшем случае O(1) при ограниченном алфавите)
    private static List<Integer> findAnagrams(String s, String p) {
        List<Integer> result = new ArrayList<>();

        if (s == null || p == null || s.length() < p.length()) {
            return result;
        }

        Map<Character, Integer> charMap = new HashMap<>();

        // Заполняем map для строки p
        for (char c : p.toCharArray()) {
            charMap.put(c, charMap.getOrDefault(c, 0) + 1);
        }

        int left = 0, right = 0; // left - левый указатель, right - правый указатель
        int count = p.length(); // count - количество символов в p

        while (right < s.length()) {
            char rightChar = s.charAt(right);

            // Если символ есть в charMap и его количество > 0, уменьшаем count
            if (charMap.getOrDefault(rightChar, 0) > 0) {
                count--;
            }
            // Уменьшаем счётчик символа в map
            charMap.put(rightChar, charMap.getOrDefault(rightChar, 0) - 1);
            right++;

            // Если count == 0, значит, нашли анаграмму
            if (count == 0) {
                result.add(left);
            }

            // Если окно достигло размера p, сдвигаем левый указатель
            if (right - left == p.length()) {
                char leftChar = s.charAt(left);
                // Если символ в левой части окна был "полезным" (его счётчик >= 0), восстанавливаем count
                if (charMap.getOrDefault(leftChar, 0) >= 0) {
                    count++;
                }
                // Восстанавливаем счётчик символа
                charMap.put(leftChar, charMap.getOrDefault(leftChar, 0) + 1);
                left++;
            }
        }

        return result; // Возвращает список индексов анаграмм
    }

    // findAnagrams находит все анаграммы строки p в строке s и возвращает их начальные индексы.
    // time: O(n), space: O(1)
    // private static List<Integer> findAnagrams(String s, String p) {
    //     List<Integer> result = new ArrayList<>();

    //     if (s == null || p == null || s.length() < p.length()) {
    //         return result;
    //     }

    //     // Используем массив вместо HashMap для улучшения производительности и соответствия O(1) по памяти
    //     int[] charCount = new int[26]; // Предполагаем, что строки содержат только строчные английские буквы

    //     // Заполняем массив частотами символов из строки p
    //     for (char c : p.toCharArray()) {
    //         charCount[c - 'a']++;
    //     }

    //     int left = 0;
    //     int right = 0;
    //     int count = p.length(); // Сколько символов ещё нужно "покрыть"

    //     while (right < s.length()) {
    //         char rightChar = s.charAt(right);
    //         // Если данный символ есть в p (то есть после уменьшения charCount он не стал отрицательным),
    //         // значит, мы "использовали" один нужный символ
    //         if (charCount[rightChar - 'a'] > 0) {
    //             count--;
    //         }
    //         charCount[rightChar - 'a']--;
    //         right++;

    //         if (count == 0) {
    //             result.add(left);
    //         }

    //         // Сдвигаем окно, когда его размер достиг длины p
    //         if (right - left == p.length()) {
    //             char leftChar = s.charAt(left);
    //             if (charCount[leftChar - 'a'] >= 0) {
    //                 count++;
    //             }
    //             charCount[leftChar - 'a']++;
    //             left++;
    //         }
    //     }

    //     return result;
    // }
}
