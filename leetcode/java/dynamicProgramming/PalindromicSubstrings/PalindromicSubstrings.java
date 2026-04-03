package leetcode.java.dynamicProgramming.PalindromicSubstrings;

/* 647. Palindromic Substrings
https://leetcode.com/problems/palindromic-substrings/description/

Дана строка s, верните количество палиндромных подстрок в ней.
Строка является палиндромом, если она читается как в прямом, так и в обратном направлении.
Подстрока — это непрерывная последовательность символов внутри строки.

Input: s = "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".

Input: s = "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
*/

public class PalindromicSubstrings {
    public static void main(String[] args) {
        System.out.println(countSubstrings("abc")); // 3 ("a", "b", "c")
    }

    // countSubstrings подсчитывает количество палиндромных подстрок в строке
    // time: O(n^2), где n - длина строки, space: O(1)
    private static int countSubstrings(String s) {
        int totalCount = 0;
        for (int i = 0; i < s.length(); i++) {
            // Подсчет палиндромов нечетной длины
            totalCount += extendPalindrome(s, i, i);
            // Подсчет палиндромов четной длины
            totalCount += extendPalindrome(s, i, i + 1);
        }
        return totalCount;
    }

    // extendPalindrome расширяется от центра и подсчитывает палиндромы
    // time: O(n), где n - длина строки, space: O(1)
    public static int extendPalindrome(String s, int left, int right) {
        int count = 0;
        while (left >= 0 && right < s.length() && s.charAt(left) == s.charAt(right)) {
            count++;
            left--;
            right++;
        }
        return count;
    }
}
