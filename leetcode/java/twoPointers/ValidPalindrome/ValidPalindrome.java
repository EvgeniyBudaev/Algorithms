package twoPointers.ValidPalindrome;

/* 125. Valid Palindrome
https://leetcode.com/problems/valid-palindrome/description/

Фраза является палиндромом, если после преобразования всех прописных букв в строчные и удаления всех небуквенно-цифровых
символов она читается одинаково и вперед, и назад. Буквенно-цифровые символы включают буквы и цифры.
Учитывая строку s, верните true, если это палиндром, или false в противном случае.

Input: s = "A man, a plan, a canal: Panama"
Output: true
Объяснение: "amanaplanacanalpanama" палиндром.

Input: s = "race a car"
Output: false
Объяснение: "raceacar" не палиндром.

Input: s = " "
Output: true
Объяснение: s — это пустая строка "" после удаления небуквенно-цифровых символов.
Поскольку пустая строка читается одинаково и вперед, и назад, она является палиндромом.
*/

public class ValidPalindrome {
    public static void main(String[] args) {
        System.out.println(isPalindrome("A man, a plan, a canal: Panama")); // true
    }

    // isAlphaNumeric проверяет, является ли переданный символ буквенно-цифровым
    // time: O(1), space: O(1)
    private static boolean isAlphaNumeric(char c) {
        return Character.isLetter(c) || Character.isDigit(c);
    }

    // isPalindrome проверяет, является ли строка s палиндромом после нормализации.
    // time: O(n), где n - количество символов в строке, space: O(1)
    private static boolean isPalindrome(String s) {
        int left = 0;
        int right = s.length() - 1;

        while (left < right) {
            // пропускаем символы, которые не являются буквенно-цифровыми
            boolean skipLeft = !isAlphaNumeric(s.charAt(left));
            if (skipLeft) {
                left++;
                continue;
            }

            // пропускаем символы, которые не являются буквенно-цифровыми
            boolean skipRight = !isAlphaNumeric(s.charAt(right));
            if (skipRight) {
                right--;
                continue;
            }

            // сравниваем символы, которые являются буквенно-цифровыми
            boolean endsEqual = Character.toLowerCase(s.charAt(left)) == Character.toLowerCase(s.charAt(right));
            if (!endsEqual) {
                return false;
            }

            left++;
            right--;
        }

        return true; // строка является палиндромом
    }
}
