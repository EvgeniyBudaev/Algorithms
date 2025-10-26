package array.LongestPalindromicSubstring;

/* 5. Longest Palindromic Substring
https://leetcode.com/problems/longest-palindromic-substring/description/

Для заданной строки s вернуть самую длинную палиндромную подстроку в s.

Input: s = "babad"
Output: "bab"
Пояснение: «aba» также является допустимым ответом.
*/

public class LongestPalindromicSubstring {
    public static void main(String[] args) {
        System.out.println(longestPalindrome("babad")); // "aba"
    }

    // longestPalindrome - возвращает самую длинную палиндромную подстроку в строке s.
    // time: O(n^2), space: O(1)
    private static String longestPalindrome(String s) {
        if (s == null || s.length() < 1) {
            return "";
        }

        int start = 0, end = 0;

        for (int i = 0; i < s.length(); i++) {
            // Проверяем палиндромы нечетной длины
            int len1 = expandAroundCenter(s, i, i);
            // Проверяем палиндромы четной длины
            int len2 = expandAroundCenter(s, i, i + 1);
            // Выбираем максимальную длину
            int maxLen = Math.max(len1, len2);

            // Если нашли палиндром длиннее текущего
            if (maxLen > end - start) {
                start = i - (maxLen - 1) / 2; // Вычитаем половину от длины палиндрома, чтобы получить начало палиндрома
                end = i + maxLen / 2;         // Прибавляем половину от длины палиндрома, чтобы получить конец палиндрома
            }
        }

        // Возвращаем палиндром
        return s.substring(start, end + 1);
    }

    // expandAroundCenter - возвращает длину палиндрома, который находится вокруг центра left и right.
    // time: O(n), space: O(1)
    private static int expandAroundCenter(String s, int left, int right) {
        // Проверяем, что символы с обоих сторон совпадают
        while (left >= 0 && right < s.length() && s.charAt(left) == s.charAt(right)) {
            left--;
            right++;
        }
        return right - left - 1; // Возвращаем длину палиндрома
    }
}
