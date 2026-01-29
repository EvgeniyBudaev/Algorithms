package leetcode.java.hashing.CheckIfTheSentenceIsPangram;

import java.util.HashSet;
import java.util.Set;

/* 1832. Check if the Sentence Is Pangram
https://leetcode.com/problems/check-if-the-sentence-is-pangram/description/

Панграмма – это предложение, в котором каждая буква английского алфавита встречается хотя бы один раз.
Учитывая строковое предложение, содержащее только строчные буквы английского языка, верните true, если предложение
является панграммой, или false в противном случае.

Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
Output: true

Input: sentence = "leetcode"
Output: false
*/

public class CheckIfTheSentenceIsPangram {
    public static void main(String[] args) {
        System.out.println(checkIfPangram("thequickbrownfoxjumpsoverthelazydog")); // true
        System.out.println(checkIfPangram("leetcode"));                            // false
    }

    // checkIfPangram проверяет, является ли предложение панграммой.
    // time: O(n), space: O(1)
    private static boolean checkIfPangram(String sentence) {
        // Оптимизация: если длина меньше 26, панграмма невозможна
        if (sentence.length() < 26) {
            return false;
        }

        Set<Character> letters = new HashSet<>(); // Для хранения уникальных букв
        for (char c : sentence.toCharArray()) {
            if (Character.isLetter(c)) { // Если символ является буквой
                letters.add(Character.toLowerCase(c));
                // Ранний выход: если уже собраны все 26 букв
                if (letters.size() == 26) {
                    return true;
                }
            }
        }

        return letters.size() == 26; // Если количество уникальных букв равно 26, то это панграмма
    }
}
