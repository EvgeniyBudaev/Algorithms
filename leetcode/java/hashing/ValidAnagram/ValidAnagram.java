package hashing.ValidAnagram;

import java.util.HashMap;
import java.util.Map;

/* 242. Valid Anagram
https://leetcode.com/problems/valid-anagram/description/

Учитывая две строки s и t, верните true, если t является анаграммой s, и false в противном случае.
Анаграмма — это слово или фраза, образованная путем перестановки букв другого слова или фразы, обычно с использованием
всех исходных букв ровно один раз.

Input: s = "anagram", t = "nagaram"
Output: true

Input: s = "rat", t = "car"
Output: false
*/

public class ValidAnagram {
    public static void main(String[] args) {
        System.out.println(isAnagram("anagram", "nagaram")); // true
        System.out.println(isAnagram("rat", "car"));         // false
    }

    // isAnagram проверяет, являются ли строки анаграммами.
    // time: O(n), space: O(n)
    private static boolean isAnagram(String s, String t) {
        // Если длины строк разные, они не могут быть анаграммами
        if (s.length() != t.length()) {
            return false;
        }

        // Создаем map для подсчета символов
        Map<Character, Integer> charCount = new HashMap<>();

        // Подсчитываем количество каждого символа в первой строке
        for (char c : s.toCharArray()) {
            charCount.put(c, charCount.getOrDefault(c, 0) + 1);
        }

        // Проверяем вторую строку
        for (char c : t.toCharArray()) {
            // Если символа нет в первой строке или его недостаточно
            if (!charCount.containsKey(c) || charCount.get(c) <= 0) {
                return false;
            }
            // Уменьшаем счетчик символа
            charCount.put(c, charCount.get(c) - 1);
        }

        return true; // Все символы в обеих строках присутствуют одинаковое количество раз
    }
}
