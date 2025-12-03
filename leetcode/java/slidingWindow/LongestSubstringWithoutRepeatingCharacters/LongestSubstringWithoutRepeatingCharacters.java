package slidingWindow.LongestSubstringWithoutRepeatingCharacters;

import java.util.HashSet;
import java.util.Set;

/* 3. Longest Substring Without Repeating Characters
https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

Дана строка s. Найдите длину самой длинной подстрока без повторения символов.

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

public class LongestSubstringWithoutRepeatingCharacters {
    public static void main(String[] args) {
        System.out.println(lengthOfLongestSubstring("abcabcbb")); // 3
    }

    // lengthOfLongestSubstring находит длину самой длинной подстрока без повторения символов.
    // time: O(n), space: O(n)
    private static int lengthOfLongestSubstring(String s) {
        // Если строка пустая или состоит из одного символа, то длина самой длинной подстроки равна длине строки.
        if (s == null || s.length() <= 1) {
            return s == null ? 0 : s.length();
        }

        Set<Character> charSet = new HashSet<>(); // charSet - карта для хранения уникальных символов
        int left = 0, right = 0, maxLength = 0; // left - левая граница, right - правая граница, maxLength - максимальная длина подстроки

        while (right < s.length()) {
            char currentChar = s.charAt(right); // текущий символ

            if (!charSet.contains(currentChar)) { // если символ не найден в charSet, то добавляем его в charSet
                charSet.add(currentChar); // добавляем символ в charSet
                maxLength = Math.max(maxLength, right - left + 1); // вычисляем максимальную длину подстроки
                right++; // сдвигаем правую границу
            } else { // если символ найден в charSet, то удаляем символ с левой границы и сдвигаем левую границу
                charSet.remove(s.charAt(left));
                left++;
            }
        }

        return maxLength; // возвращаем максимальную длину подстроки
    }
}
