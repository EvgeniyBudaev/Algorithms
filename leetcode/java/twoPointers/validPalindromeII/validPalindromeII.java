package twoPointers.validPalindromeII;

/* 680. Valid Palindrome II
https://leetcode.com/problems/valid-palindrome-ii/description/

Учитывая строку s, верните true, если s может быть палиндромом после удаления из нее не более одного символа.

Input: s = "aba"
Output: true

Input: s = "abca"
Output: true
Объяснение: Вы можете удалить символ 'c'.

Input: s = "abc"
Output: false
*/

public class validPalindromeII {
    public static void main(String[] args) {
        System.out.println(validPalindrome("aba"));   // true
        System.out.println(validPalindrome("abca"));  // true
        System.out.println(validPalindrome("abc"));   // false
    }

    // validPalindrome проверяет, можно ли сделать строку палиндромом, удалив один символ.
    // time: O(n), space: O(1)
    public static boolean validPalindrome(String s) {
        int left = 0;
        int right = s.length() - 1;
        
        while (left < right) {
            // Если символы не равны, проверяем, можно ли сделать строку палиндромом, удалив один символ
            if (s.charAt(left) != s.charAt(right)) {
                return isPalindrome(s, left + 1, right) || isPalindrome(s, left, right - 1);
            }
            left++;
            right--;
        }
        return true; // Если цикл завершился успешно, строка является палиндромом
    }

    // isPalindrome проверяет, является ли строка палиндромом.
    // time: O(n), space: O(1)
    private static boolean isPalindrome(String s, int left, int right) {
        while (left < right) {
            // Если символы не равны, строка не является палиндромом
            if (s.charAt(left) != s.charAt(right)) {
                return false;
            }
            left++;
            right--;
        }
        return true; // Если цикл завершился успешно, строка является палиндромом
    }
}
