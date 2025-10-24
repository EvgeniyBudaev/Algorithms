package array.IsomorphicStrings;

import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

/* 205. Isomorphic Strings
https://leetcode.com/problems/isomorphic-strings/description/

Даны две строки s и t. Определите, изоморфны ли они.
Две строки s и t изоморфны, если символы в s можно заменить, чтобы получить t.
Все вхождения символа должны быть заменены другим символом с сохранением порядка символов. Никакие два символа не могут
сопоставляться одному и тому же символу, но символ может сопоставляться сам с собой.
Например, рассмотрим строки ACAB и XCXY.
Они изоморфны, поскольку мы можем отобразить 'A' —> 'X', 'B' —> 'Y'и 'C' —> 'C'.

Input: s = "egg", t = "add"
Output: true

Input: s = "foo", t = "bar"
Output: false

Input: s = "paper", t = "title"
Output: true
*/

public class IsomorphicStrings {
    public static void main(String[] args) {
        System.out.println(isIsomorphic("egg", "add")); // true
        System.out.println(isIsomorphic("abc", "aaa")); // false
    }

    // isIsomorphic определяет, изоморфны ли строки.
    // time: O(n), space: O(n), где n - длина строки.
    private static boolean isIsomorphic(String s, String t) {
        // Если длины строк не совпадают, то они не изоморфны.
        if (s.length() != t.length()) {
            return false;
        }

        // Используем карту для отображения символов.
        Map<Character, Character> charMap = new HashMap<>();
        // Используем множество для отслеживания уже сопоставленных символов.
        Set<Character> valuesSet = new HashSet<>();

        for (int i = 0; i < s.length(); i++) { // Проходим по символам в строке s.
            char sLetter = s.charAt(i); // Текущий символ в строке s.
            char tLetter = t.charAt(i); // Текущий символ в строке t.

            if (!charMap.containsKey(sLetter)) {
                // Если символ уже существует, но не соответствует текущему символу строки t, возвращаем false.
                if (valuesSet.contains(tLetter)) {
                    return false;
                }

                // Символ не существует, добавляем его в карту.
                valuesSet.add(tLetter); // Добавляем в множество, чтобы отслеживать уже сопоставленные символы.
                charMap.put(sLetter, tLetter); // Добавляем в карту соответствия символов.
                continue; // Переходим к следующему символу.
            }

            // Если символ уже существует, но не соответствует текущему символу строки t, возвращаем false.
            if (charMap.get(sLetter) != tLetter) {
                return false;
            }
        }

        return true; // Строки изоморфны.
    }
}
