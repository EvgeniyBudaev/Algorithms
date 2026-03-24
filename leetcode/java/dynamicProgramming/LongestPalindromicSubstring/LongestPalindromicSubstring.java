package leetcode.java.dynamicProgramming.LongestPalindromicSubstring;

/* 5. Longest Palindromic Substring
https://leetcode.com/problems/longest-palindromic-substring/description/

Учитывая строку s, вернуть самую длинную палиндромный подстрока в s.

Input: s = "babad"
Output: "bab"
Пояснение: «aba» также является допустимым ответом.

Input: s = "cbbd"
Output: "bb"
*/

public class LongestPalindromicSubstring {
    public static void main(String[] args) {
        System.out.println(longestPalindrome("babad")); // "bab" или "aba"
    }

    private static String longestPalindrome(String s) {
        if (s == null || s.length() < 1) {
            return "";
        }

        int start = 0, end = 0;

        for (int i = 0; i < s.length(); i++) {
            int len1 = expandAroundCenter(s, i, i);
            int len2 = expandAroundCenter(s, i, i + 1);
            int maxLen = Math.max(len1, len2);

            if (maxLen > end - start) {
                start = i - (maxLen - 1) / 2;
                end = i + maxLen / 2;
            }
        }

        return s.substring(start, end + 1);
    }

    private static int expandAroundCenter(String s, int left, int right) {
        while (left >= 0 && right < s.length() && s.charAt(left) == s.charAt(right)) {
            left--;
            right++;
        }
        return right - left - 1;
    }
}
