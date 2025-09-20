package twoPointers.PalindromeNumber;

/* 9. Palindrome Number
https://leetcode.com/problems/palindrome-number/description/

Учитывая целое число x, верните true, если x является палиндром и false в противном случае.

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
*/

public class PalindromeNumber {
    public static void main(String[] args) {
        System.out.println(isPalindrome(121)); // true
        System.out.println(isPalindrome(-121)); // false
    }

    // isPalindrome проверяет, является ли целое число x палиндромом.
    // time: O(n), space: O(1)
    private static boolean isPalindrome(int x) {
        // Отрицательные числа не могут быть палиндромами
        if (x < 0) {
            return false;
        }

        // Числа, оканчивающиеся на 0 (кроме 0) не могут быть палиндромами
        if (x % 10 == 0 && x != 0) {
            return false;
        }

        int reversed = 0;

        // Переворачиваем половину числа
        while (x > reversed) {
            reversed = reversed * 10 + x % 10;
            x /= 10;
        }
        
        // Для четного количества цифр: x == reversed
        // Для нечетного количества цифр: x == reversed / 10
        return x == reversed || x == reversed / 10;
    }

    // isPalindrome проверяет, является ли целое число x палиндромом.
    // time: O(n), space: O(n) - из-за создания строки
    // private static boolean isPalindrome(int x) {
    //     String str = String.valueOf(x); // Преобразуем число в строку
    //     int left = 0;
    //     int right = str.length() - 1;
        
    //     while (left < right) {
    //         if (str.charAt(left) != str.charAt(right)) {
    //             return false;
    //         }
    //         left++;
    //         right--;
    //     }
        
    //     return true;
    // }
}
